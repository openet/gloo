package printers

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc_json"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/xdsinspection"
	plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws/ec2"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	"github.com/olekukonko/tablewriter"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

// key for the functionNames within the Details field of statuses, to be populated with the names of functions
// retrieved from gRPC descriptors if applicable
const functionNamesKey = "functionNames"

// PrintUpstreams prints an UpstreamList, leveraging cliutils.PrintList()
func PrintUpstreams(upstreams v1.UpstreamList, outputType OutputType, xdsDump *xdsinspection.XdsDump, namespace string) error {
	if outputType == KUBE_YAML {
		return PrintKubeCrdList(upstreams.AsInputResources(), v1.UpstreamCrd)
	}

	upstreams.Each(addFunctionsFromGrpcTranscoder(namespace))

	return cliutils.PrintList(outputType.String(), "", upstreams,
		func(data interface{}, w io.Writer) error {
			UpstreamTable(xdsDump, data.(v1.UpstreamList), w)
			return nil
		}, os.Stdout)
}

// UpstreamTable prints upstreams in table format to io.Writer
func UpstreamTable(xdsDump *xdsinspection.XdsDump, upstreams []*v1.Upstream, w io.Writer) {
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Upstream", "type", "status", "details"})

	for _, us := range upstreams {
		name := us.GetMetadata().GetName()
		s := upstreamStatus(us)

		u := upstreamType(us)
		details := upstreamDetails(us, xdsDump)

		if len(details) == 0 {
			details = []string{""}
		}
		for i, line := range details {
			if i == 0 {
				table.Append([]string{name, u, s, line})
			} else {
				table.Append([]string{"", "", "", line})
			}
		}
	}

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

func upstreamStatus(us *v1.Upstream) string {
	return AggregateNamespacedStatuses(us.GetNamespacedStatuses(), func(status *core.Status) string {
		return status.GetState().String()
	})
}

func upstreamType(up *v1.Upstream) string {
	if up == nil {
		return "Invalid"
	}

	switch up.GetUpstreamType().(type) {
	case *v1.Upstream_Aws:
		return "AWS Lambda"
	case *v1.Upstream_Azure:
		return "Azure"
	case *v1.Upstream_Consul:
		return "Consul"
	case *v1.Upstream_AwsEc2:
		return "AWS EC2"
	case *v1.Upstream_Kube:
		return "Kubernetes"
	case *v1.Upstream_Static:
		return "Static"
	default:
		return "Unknown"
	}
}

func upstreamDetails(up *v1.Upstream, xdsDump *xdsinspection.XdsDump) []string {
	if up == nil {
		return []string{"invalid: upstream was nil"}
	}

	var details []string
	add := func(s ...string) {
		details = append(details, s...)
	}
	switch usType := up.GetUpstreamType().(type) {
	case *v1.Upstream_Aws:
		var functions []string
		for _, fn := range usType.Aws.GetLambdaFunctions() {
			functions = append(functions, fn.GetLambdaFunctionName())
		}

		add(
			fmt.Sprintf("region: %v", usType.Aws.GetRegion()),
			fmt.Sprintf("secret: %s", stringifyKey(usType.Aws.GetSecretRef())),
		)
		for i := range functions {
			if i == 0 {
				add("functions:")
			}
			add(fmt.Sprintf("- %v", functions[i]))
		}
	case *v1.Upstream_AwsEc2:
		add(
			fmt.Sprintf("role:           %v", usType.AwsEc2.GetRoleArn()),
			fmt.Sprintf("uses public ip: %v", usType.AwsEc2.GetPublicIp()),
			fmt.Sprintf("port:           %v", usType.AwsEc2.GetPort()),
		)
		add(getEc2TagFiltersString(usType.AwsEc2.GetFilters())...)
		instances := xdsDump.GetEc2InstancesForUpstream(up.GetMetadata().Ref())
		add(
			"EC2 Instance Ids:",
		)
		add(
			instances...,
		)
	case *v1.Upstream_Azure:
		var functions []string
		for _, fn := range usType.Azure.GetFunctions() {
			functions = append(functions, fn.GetFunctionName())
		}
		add(
			fmt.Sprintf("function app name: %v", usType.Azure.GetFunctionAppName()),
			fmt.Sprintf("secret: %s", stringifyKey(usType.Azure.GetSecretRef())),
		)

		for i := range functions {
			if i == 0 {
				add("functions:")
			}
			add(fmt.Sprintf("- %v", functions[i]))
		}
	case *v1.Upstream_Consul:
		add(
			fmt.Sprintf("svc name: %v", usType.Consul.GetServiceName()),
			fmt.Sprintf("svc tags: %v", usType.Consul.GetServiceTags()),
		)
		if usType.Consul.GetServiceSpec() != nil {
			add(linesForServiceSpec(usType.Consul.GetServiceSpec())...)
		}
	case *v1.Upstream_Kube:
		add(
			fmt.Sprintf("svc name:      %v", usType.Kube.GetServiceName()),
			fmt.Sprintf("svc namespace: %v", usType.Kube.GetServiceNamespace()),
			fmt.Sprintf("port:          %v", usType.Kube.GetServicePort()),
		)
		if usType.Kube.GetServiceSpec() != nil {
			add(linesForServiceSpec(usType.Kube.GetServiceSpec())...)
		}
	case *v1.Upstream_Static:
		for i := range usType.Static.GetHosts() {
			if i == 0 {
				add("hosts:")
			}
			add(fmt.Sprintf("- %v:%v", usType.Static.GetHosts()[i].GetAddr(), usType.Static.GetHosts()[i].GetPort()))
		}
		if usType.Static.GetServiceSpec() != nil {
			add(linesForServiceSpec(usType.Static.GetServiceSpec())...)
		}
	}
	add("")
	return details
}

func linesForServiceSpec(serviceSpec *plugins.ServiceSpec) []string {
	var spec []string
	add := func(s ...string) {
		spec = append(spec, s...)
	}
	switch plug := serviceSpec.GetPluginType().(type) {
	case *plugins.ServiceSpec_Rest:
		add("REST service:")
		var functions []string
		for restFunc, transform := range plug.Rest.GetTransformations() {
			var method, path string
			methodP, ok := transform.GetHeaders()[":method"]
			if ok {
				method = methodP.GetText()
			}
			pathP, ok := transform.GetHeaders()[":path"]
			if ok {
				path = pathP.GetText()
			}
			if false {
				//TODO ilackarms: save it for -o wide
				functions = append(functions, fmt.Sprintf("- %v (%v %v)", restFunc, method, path))
			}
			functions = append(functions, fmt.Sprintf("- %v", restFunc))
		}
		// needed because map
		sort.Strings(functions)

		for i := range functions {
			if i == 0 {
				add("functions:")
			}
			add(functions[i])
		}
	case *plugins.ServiceSpec_Grpc:
		add("gRPC service:")
		for _, grpcService := range plug.Grpc.GetGrpcServices() {
			add(fmt.Sprintf("  %v", grpcService.GetServiceName()))
			for _, fn := range grpcService.GetFunctionNames() {
				add(fmt.Sprintf("  - %v", fn))
			}
		}
	case *plugins.ServiceSpec_GrpcJsonTranscoder:
		add("gRPC service:")
		descriptorBin := plug.GrpcJsonTranscoder.GetProtoDescriptorBin()
		for _, grpcService := range plug.GrpcJsonTranscoder.GetServices() {
			add(fmt.Sprintf("  %v", grpcService))
			methodDescriptors := getMethodDescriptors(grpcService, descriptorBin)
			for i := 0; i < methodDescriptors.Len(); i++ {
				add(fmt.Sprintf("  - %v", methodDescriptors.Get(i).Name()))
			}
		}
	}
	return spec
}

func getMethodDescriptors(service string, descriptorSet []byte) protoreflect.MethodDescriptors {
	fds := &descriptor.FileDescriptorSet{}
	err := proto.Unmarshal(descriptorSet, fds)
	if err != nil {
		fmt.Println("unable to unmarshal descriptor")
		return nil
	}
	files, err := protodesc.NewFiles(fds)
	if err != nil {
		fmt.Println("unable to create proto registry files")
		return nil
	}
	descriptor, err := files.FindDescriptorByName(protoreflect.FullName(service))
	if err != nil {
		fmt.Println("unable to fin descriptor")
		return nil
	}
	serviceDescriptor, ok := descriptor.(protoreflect.ServiceDescriptor)
	if !ok {
		fmt.Println("unable to decode service descriptor")
		return nil
	}
	return serviceDescriptor.Methods()
}

// stringifyKey for a resource likely could be done more nicely with spew
// or a better accessor but minimal this avoids panicing on nested references to nils
func stringifyKey(plausiblyNilRef *core.ResourceRef) string {

	if plausiblyNilRef == nil {
		return "<None>"
	}
	return plausiblyNilRef.Key()

}

func getEc2TagFiltersString(filters []*ec2.TagFilter) []string {
	var out []string
	add := func(s ...string) {
		out = append(out, s...)
	}

	var kFilters []*ec2.TagFilter_Key
	var kvFilters []*ec2.TagFilter_KvPair
	for _, f := range filters {
		switch x := f.GetSpec().(type) {
		case *ec2.TagFilter_Key:
			kFilters = append(kFilters, x)
		case *ec2.TagFilter_KvPair_:
			kvFilters = append(kvFilters, x.KvPair)
		}
	}
	if len(kFilters) == 0 {
		add(fmt.Sprintf("key filters: (none)"))
	} else {
		add(fmt.Sprintf("key filters:"))
		for _, f := range kFilters {
			add(fmt.Sprintf("- %v", f.Key))
		}
	}
	if len(kvFilters) == 0 {
		add(fmt.Sprintf("key-value filters: (none)"))
	} else {
		add(fmt.Sprintf("key-value filters:"))
		for _, f := range kvFilters {
			add(fmt.Sprintf("- %v: %v", f.GetKey(), f.GetValue()))
		}
	}
	return out
}

// returns a function that, if applicable, augments a given Upstream's status to contain the names of gRPC functions
// extracted from the GrpcJsonTranscoder
// this is a no-op if the Upstream doesn't have a GrpcJsonTranscoder
// TODO: retrieve function names from GrpcJsonTranscoder descriptor sources other than in-line bin (see https://github.com/solo-io/gloo/issues/9798)
func addFunctionsFromGrpcTranscoder(namespace string) func(*v1.Upstream) {
	return func(up *v1.Upstream) {
		var functionNames map[string]any

		// retrieve function names if applicable
		switch usType := up.GetUpstreamType().(type) {
		case *v1.Upstream_Kube:
			if gjt := usType.GetServiceSpec().GetGrpcJsonTranscoder(); gjt != nil {
				if gjt.GetProtoDescriptorBin() != nil {
					functionNames = getFunctionsFromDescriptorBin(gjt)
				}
			}
		case *v1.Upstream_Consul:
			if gjt := usType.GetServiceSpec().GetGrpcJsonTranscoder(); gjt != nil {
				if gjt.GetProtoDescriptorBin() != nil {
					functionNames = getFunctionsFromDescriptorBin(gjt)
				}
			}
		case *v1.Upstream_Static:
			if gjt := usType.GetServiceSpec().GetGrpcJsonTranscoder(); gjt != nil {
				if gjt.GetProtoDescriptorBin() != nil {
					functionNames = getFunctionsFromDescriptorBin(gjt)
				}
			}
		}

		// if we got any function names, add them to the status for the given namespace
		// the function names are the same regardless of namespace, so adding to multiple statuses would be redundant
		// we therefore only add to the given namespace's status
		if functionNames != nil {
			for ns, status := range up.GetNamespacedStatuses().GetStatuses() {
				if ns == namespace {
					addFunctionNamesToStatus(status, functionNames)
				}
			}
		}
	}
}

// returns a map of service names to lists of function names
// the values in the map are of `any` type, which is supported by structpb.NewStruct(), while []string is not
func getFunctionsFromDescriptorBin(gjt *grpc_json.GrpcJsonTranscoder) map[string]any {
	grpcFunctions := make(map[string]any)

	descriptorBin := gjt.GetProtoDescriptorBin()

	for _, grpcService := range gjt.GetServices() {
		methodDescriptors := getMethodDescriptors(grpcService, descriptorBin)

		funcList := make([]any, methodDescriptors.Len())
		for i := 0; i < methodDescriptors.Len(); i++ {
			funcList[i] = fmt.Sprintf("%s", methodDescriptors.Get(i).Name())
		}
		grpcFunctions[grpcService] = funcList
	}

	return grpcFunctions
}

// add a struct created from the functionNames map to the given status under the Details field for function names
// note that this will result in an object in the following format:
//
//	{
//	  "fields": {
//	    "functionNames": {
//	      "structValue": {
//	        "fields": {
//	          "solo.examples.v1.StoreService": {
//	            "listValue": {
//	              "values": [
//	                {
//	                  "stringValue": "CreateItem"
//	                },
//	                {
//	                  "stringValue": "ListItems"
//	                },
//	                {
//	                  "stringValue": "DeleteItem"
//	                },
//	                {
//	                  "stringValue": "GetItem"
//	                }
//	              ]
//	            }
//	          }
//	        }
//	      }
//	    }
//	  }
//	}
func addFunctionNamesToStatus(status *core.Status, functionNames map[string]any) {
	if status.GetDetails() == nil {
		status.Details = &structpb.Struct{
			Fields: make(map[string]*structpb.Value),
		}
	}

	functionNamesStruct, _ := structpb.NewStruct(functionNames)

	status.GetDetails().GetFields()[functionNamesKey] = &structpb.Value{
		Kind: &structpb.Value_StructValue{
			StructValue: functionNamesStruct,
		},
	}
}
