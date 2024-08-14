package check

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/solo-io/gloo/pkg/utils/kubeutils"

	"github.com/golang/protobuf/jsonpb"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/printers"
	glooinstancev1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1"
	"github.com/solo-io/solo-apis/pkg/api/fed.solo.io/v1/types"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CheckMulticlusterResources(ctx context.Context, printer printers.P, opts *options.Options) {
	// check if the gloo fed deployment exists
	client := helpers.MustKubeClientWithKubecontext(opts.Top.KubeContext)
	_, err := client.AppsV1().Deployments(opts.Metadata.GetNamespace()).Get(ctx, constants.GlooFedDeploymentName, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			printer.AppendMessage("\nSkipping Gloo Instance check -- Gloo Federation not detected.")
		} else {
			fmt.Printf("Warning: could not get Gloo Fed deployment: %v. Skipping Gloo Instance check.\n", err)
		}
		return
	}

	cfg, err := kubeutils.GetRestConfigWithKubeContext(opts.Top.KubeContext)
	if err != nil {
		fmt.Printf("Warning: could not get kubernetes config to check multicluster resources: %v. "+
			"Skipping Gloo Instance check.\n", err)
		return
	}
	instanceReader, err := getUnstructuredGlooInstanceReader(cfg)
	if err != nil {
		fmt.Printf("Warning: could not get Gloo Instance client: %v. Skipping Gloo Instance check.\n", err)
		return
	}
	glooInstanceList, err := instanceReader.listGlooInstances(ctx)
	if err != nil {
		if meta.IsNoMatchError(err) {
			printer.AppendMessage("Skipping Gloo Instance check -- Gloo Federation not detected.")
			return
		}
		fmt.Printf("Warning: could not list Gloo Instances: %v\n", err)
		return
	}
	printer.AppendMessage("\nDetected Gloo Federation!")
	for _, glooInstance := range glooInstanceList.Items {
		fmt.Printf("\nChecking Gloo Instance %s... ", glooInstance.GetName())
		printGlooInstanceCheckSummary("Deployments", glooInstance.Spec.GetCheck().GetDeployments())
		printGlooInstanceCheckSummary("Pods", glooInstance.Spec.GetCheck().GetPods())
		printGlooInstanceCheckSummary("Settings", glooInstance.Spec.GetCheck().GetSettings())
		printGlooInstanceCheckSummary("Upstreams", glooInstance.Spec.GetCheck().GetUpstreams())
		printGlooInstanceCheckSummary("UpstreamGroups", glooInstance.Spec.GetCheck().GetUpstreamGroups())
		printGlooInstanceCheckSummary("AuthConfigs", glooInstance.Spec.GetCheck().GetAuthConfigs())
		printGlooInstanceCheckSummary("RateLimitConfigs", glooInstance.Spec.GetCheck().GetRateLimitConfigs())
		printGlooInstanceCheckSummary("VirtualServices", glooInstance.Spec.GetCheck().GetVirtualServices())
		printGlooInstanceCheckSummary("RouteTables", glooInstance.Spec.GetCheck().GetRouteTables())
		printGlooInstanceCheckSummary("Gateways", glooInstance.Spec.GetCheck().GetGateways())
		printGlooInstanceCheckSummary("MatchableHttpGateways", glooInstance.Spec.GetCheck().GetMatchableHttpGateways())
		printGlooInstanceCheckSummary("MatchableTcpGateways", glooInstance.Spec.GetCheck().GetMatchableTcpGateways())
		printGlooInstanceCheckSummary("Proxies", glooInstance.Spec.GetCheck().GetProxies())
		fmt.Printf("\n\n")
	}
}

func printGlooInstanceCheckSummary(resourceType string, resource *types.GlooInstanceSpec_Check_Summary) {
	fmt.Printf("\nChecking %s... ", resourceType)

	ok := true
	for _, errReport := range resource.GetErrors() {
		fmt.Printf("\nFound error in %s %s\n", errReport.GetRef().GetNamespace(), errReport.GetRef().GetName())
		fmt.Printf("Reason: %s\n", errReport.GetMessage())
		ok = false
	}
	for _, warningReport := range resource.GetWarnings() {
		fmt.Printf("Found warning in %s %s\n", warningReport.GetRef().GetNamespace(), warningReport.GetRef().GetName())
		fmt.Printf("Reason: %s\n", warningReport.GetMessage())
		ok = false
	}
	if ok {
		fmt.Printf("OK")
	}
}

// unstructuredGlooInstanceReader provides a forwards-compatible means of listing GlooInstances.
// If new fields are added to the GlooInstance API in future Gloo Fed releases, we do not want to
// fail to list GlooInstances due to an error unmarshaling unrecognized fields.
type unstructuredGlooInstanceReader struct {
	unstructuredReader client.Reader
}

func getUnstructuredGlooInstanceReader(cfg *rest.Config) (*unstructuredGlooInstanceReader, error) {
	scheme := runtime.NewScheme()
	if err := glooinstancev1.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return &unstructuredGlooInstanceReader{unstructuredReader: client}, nil
}

func (c *unstructuredGlooInstanceReader) listGlooInstances(ctx context.Context) (glooinstancev1.GlooInstanceList, error) {
	glooInstanceGVK := schema.GroupVersionKind{
		Group:   "fed.solo.io",
		Version: "v1",
		Kind:    "GlooInstance",
	}
	unstructuredList := &unstructured.UnstructuredList{}
	unstructuredList.SetGroupVersionKind(glooInstanceGVK)

	if err := c.unstructuredReader.List(ctx, unstructuredList); err != nil {
		return glooinstancev1.GlooInstanceList{}, err
	}

	glooInstanceList := glooinstancev1.GlooInstanceList{}
	for _, item := range unstructuredList.Items {
		glooInstance, err := toGlooInstance(item)
		if err != nil {
			return glooinstancev1.GlooInstanceList{}, err
		}
		glooInstanceList.Items = append(glooInstanceList.Items, glooInstance)
	}

	return glooInstanceList, nil
}

func toGlooInstance(unstr unstructured.Unstructured) (glooinstancev1.GlooInstance, error) {
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}

	spec := unstr.Object["spec"]
	specBytes, err := json.Marshal(spec)
	if err != nil {
		return glooinstancev1.GlooInstance{}, err
	}
	specPb := &types.GlooInstanceSpec{}
	if err := unmarshaler.Unmarshal(bytes.NewBuffer(specBytes), specPb); err != nil {
		return glooinstancev1.GlooInstance{}, err
	}

	metadata := unstr.Object["metadata"]
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return glooinstancev1.GlooInstance{}, err
	}
	objectMeta := &metav1.ObjectMeta{}
	if err := json.Unmarshal(metadataBytes, objectMeta); err != nil {
		return glooinstancev1.GlooInstance{}, err
	}

	return glooinstancev1.GlooInstance{
		ObjectMeta: *objectMeta,
		Spec:       *specPb,
	}, nil
}
