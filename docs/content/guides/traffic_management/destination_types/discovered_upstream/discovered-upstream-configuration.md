---
title: Discovered Upstream Configuration via Annotations
weight: 101
---

Gloo Gateway looks for discovered upstream configuration in the annotations of any Kubernetes service that it identifies. For Gloo Gateway to discover the upstream configuration, include an annotation in the service in a key-value format. The key is `gloo.solo.io/upstream_config` and the value is the upstream configuration, formatted as JSON.

For example, we can set the initial stream window size on the discovered upstream using the a modified version of the pet store manifest provided in the parent document:

{{< highlight yaml "hl_lines=7" >}}
kubectl apply -f - <<EOF
# petstore service
apiVersion: v1
kind: Service
metadata:
  annotations:
    gloo.solo.io/upstream_config: '{"initialStreamWindowSize": 2048}'
  name: petstore
  namespace: default
  labels:
    service: petstore
spec:
  ports:
  - port: 8080
    protocol: TCP
  selector:
    app: petstore
---
#petstore deployment 
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: petstore
  name: petstore
  namespace: default
spec:
  selector:
    matchLabels:
      app: petstore
  replicas: 1
  template:
    metadata:
      labels:
        app: petstore
    spec:
      containers:
      - image: soloio/petstore-example:latest
        name: petstore
        ports:
        - containerPort: 8080
          name: http
EOF
{{< /highlight >}}

Now that you created the pet store app, check for the discovered upstream. In the output of the following command, note the upstream with the namespace, name, and port of the service , `default-petstore-8080`. 

    kubectl get upstreams -n gloo-system

Review the upstream to make sure that the configuration from the Kubernetes service was picked up in the upstream.


```shell
kubectl get upstream -n gloo-system default-petstore-8080 -oyaml
```

{{< highlight yaml "hl_lines=5, 25" >}}
apiVersion: gloo.solo.io/v1
kind: Upstream
metadata:
  annotations:
    gloo.solo.io/upstream_config: '{"initialStreamWindowSize": 2048}'
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{"gloo.solo.io/upstream_config":" {\"initialStreamWindowSize\": 2048}"},"labels":{"service":"petstore"},"name":"petstore","namespace":"default"},"spec":{"ports":[{"port":8080,"protocol":"TCP"}],"selector":{"app":"petstore"}}}
  creationTimestamp: "2021-10-14T13:22:12Z"
  generation: 2
  labels:
    discovered_by: kubernetesplugin
  name: default-petstore-8080
  namespace: default
  resourceVersion: "5679"
  uid: 0ab14ba5-6377-40c5-a781-ce33b7755cdc
spec:
  discoveryMetadata:
    labels:
      service: petstore
  kube:
    selector:
      app: petstore
    serviceName: petstore
    serviceNamespace: default
    servicePort: 8080
  initialStreamWindowSize: 2048
status:
  statuses:
    default:
      reportedBy: gloo
      state: 1
{{< /highlight >}}

As you can see, the configuration set `spec.initialStreamWindowSize` to `2048` on the discovered upstream! 

## Merge strategies

By default, discovered upstreams configured via the `gloo.solo.io/upstream_config` annotation will completely overwrite top-level upstream fields for which configuration has been specified.

By setting the `gloo.solo.io/upstream_config.deep_merge` annotation to `true` on the service for which an upstream is to be discovered, you can configure Gloo Gateway to merge the provided configuration with the default upstream config. This can be useful if you rely on certain default values present when a new upstream is discovered.

## Global annotations for upstreams

You can consistently configure upstreams with the global annotations setting (`Settings.UpstreamOptions.globalAnnotations`).

When using global annotations, keep in mind the following considerations:
* To overwrite global annotations for an upstream, define the same annotation key with a different value on the Kubernetes service. For possible upstream annotations to set, see the [API docs]({{< versioned_link_path fromRoot="/reference/api/github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto.sk/" >}}).
* You can use the merge strategy annotation `gloo.solo.io/upstream_config.deep_merge` for global annotations.
* You cannot set an [SSL global annotation]({{< versioned_link_path fromRoot="/guides/security/tls/client_tls_service_annotations/" >}}) with the prefix `gloo.solo.io/sslService`.

Review the following example configuration and description. For more information, see the [API docs]({{< versioned_link_path fromRoot="/reference/api/github.com/solo-io/gloo/projects/gloo/api/v1/settings.proto.sk/#upstreamoptions" >}}).

```yaml
# Gloo Gateway Settings configuration
              upstreamOptions:
                properties:
                  globalAnnotations:
                    gloo.solo.io/upstream_config: '{"initial_stream_window_size": 1024, "max_concurrent_streams": 64}'
---
# Kubernetes Service configuration
metadata:
  annotations:
    gloo.solo.io/upstream_config: '{"use_http2": false, "initial_stream_window_size": 2048}'
```

In the Gloo Gateway settings, two global annotations are defined:
* `"initial_stream_window_size": 1024`
* `"max_concurrent_streams": 64`

In the Kubernetes service, two annotations are defined:
* `"initial_stream_window_size": 2048`
* `"use_http2": false`

The resulting upstream for the service has the following annotations configured:
* `"initial_stream_window_size": 2048`, the value in the service overwrites the global value.
* `"use_http2": false`, the annotation from the service.
* `"max_concurrent_streams": 64`, the global annotation set across upstreams.
