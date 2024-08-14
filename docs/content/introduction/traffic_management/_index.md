---
title: Traffic management
weight: 20
---

Gloo Gateway acts as the control plane to manage traffic flowing between downstream clients and upstream services. Traffic management can take many forms as a request flows through the Envoy proxies managed by Gloo Gateway. Requests from clients can be transformed, redirected, routed, and shadowed, to cite just a few examples.

---

## Fundamentals

The primary components that deal with traffic management in Gloo Gateway are as follows:

* **Gateways** - Gloo Gateway listens for incoming traffic on *Gateways*. The Gateway definition includes the protocols and ports on which Gloo Gateway listens for traffic.
* **Virtual Services** - *Virtual Services* are bound to a Gateway and configured to respond for specific domains. Each contains a set of route rules, security configuration, rate limiting, transformations, and other core routing capabilities supported by Gloo Gateway.
* **Routes** - Routes are associated with Virtual Services and direct traffic based on characteristics of the request and the upstream destination.
* **Upstreams** - Routes send traffic to destinations, called *Upstreams*. Upstreams take many forms, including Kubernetes services, AWS Lambda functions, or Consul services.

Additional information can be found in the [Gloo Gateway Core Concepts document]({{% versioned_link_path fromRoot="/introduction/architecture/concepts/" %}}).

---

## Gloo Gateway Configuration

Let's see what underpins Gloo Gateway routing with a high-level look at the layout of the Gloo Gateway configuration. This can be seen as 3 layers: the *Gateway listeners*, *Virtual Services*, and *Upstreams*. Mostly, you'll be interacting with [Virtual Services]({{% versioned_link_path fromRoot="/introduction/architecture/concepts#virtual-services" %}}), which allow you to configure the details of the API you wish to expose on the Gateway and how routing happens to the backends. [Upstreams]({{% versioned_link_path fromRoot="/introduction/architecture/concepts#upstreams" %}}) represent those backends. [Gateway]({{% versioned_link_path fromRoot="/introduction/architecture/concepts#gateways" %}}) objects help you control the listeners for incoming traffic.

<figure><img src="{{% versioned_link_path fromRoot="/img/traffic-config-ov.svg" %}}">
<figcaption style="text-align:center;font-style:italic">Figure: Example configuration of gateway, virtual service, and upstream resources.</figcaption></figure>

---
## Route Rules

Routes are the primary building block of the *Virtual Service*. A route contains matchers and an upstream which could be a single destination, a list of weighted destinations, or an upstream group. 

There are many types of **matchers**, including **Path Matching**, **Header Matching**, **Query Parameter Matching**, and **HTTP Method Matching**. Matchers can be combined in a single rule to further refine which requests will be matched against that rule.

Gloo Gateway is capable of sending matching requests to many different types of *Upstreams*, including **Single Upstream**, **Multiple Upstream**, **Upstream Groups**, Kubernetes services, and Consul services. The ability to route a request to multiple *Upstreams* or *Upstream Groups* allows Gloo Gateway to load balance requests and perform Canary Releases.

<figure><img src="{{% versioned_link_path fromRoot="/img/traffic-route-rule.svg" %}}">
<figcaption style="text-align:center;font-style:italic">Figure: Sample of routing rules that are configured in a virtual service.</figcaption></figure>

Configuring the routing engine is done with defined predicates that match on incoming requests. The contents of a request, such as headers, path, method, etc., are examined to see if they match the predicates of a route rule. If they do, the request is processed based on enabled routing features and routed to an Upstream destinations such as REST or gRPC services running in Kubernetes, EC2, etc. or Cloud Functions like Lambda. In the [Traffic Management section]({{% versioned_link_path fromRoot="/guides/traffic_management/" %}}) we'll dig into this process further.

---

## Listener configuration

The Gateway component of Gloo Gateway is what listens for incoming requests. An example configuration is shown below for an SSL Gateway. The `spec` portion defines the options for the Gateway.

{{< highlight proto "hl_lines=8-15" >}}
apiVersion: gateway.solo.io/v1
kind: Gateway
metadata:
  labels:
    app: gloo
  name: gateway-proxy-ssl
  namespace: gloo-system
spec:
  bindAddress: '::'
  bindPort: 8443
  httpGateway: {}
  proxyNames:
  - gateway-proxy
  ssl: true
  useProxyProto: false
{{< /highlight >}}

A full listing of configuration options is available in the {{< protobuf name="gateway.solo.io.Gateway" display="API reference for Gateways.">}}

The listeners on a gateway typically listen for HTTP requests coming in on a specific address and port as defined by `bindAddress` and `bindPort`. Additional options can be configured by including an `options` section in the spec. SSL for a gateway is enabled by setting the `ssl` property to `true`.

Gloo Gateway can be configured to act as a gateway on layer 7 (HTTP/S) or layer 4 (TCP). The majority of services will likely be using HTTP, but there may be some cases where applications either do not use HTTP or should be presented as a TCP endpoint. When Gloo Gateway operates as a TCP Proxy, the options for traffic management are greatly reduced. Gloo Gateway currently supports standard routing, SSL, and Server Name Indication (SNI) domain matching. Applications not using HTTP can be configured using the [TCP Proxy guide]({{% versioned_link_path fromRoot="/guides/traffic_management/listener_configuration/tcp_proxy/" %}}).

Gloo Gateway is meant to serve as an abstraction layer, simplifying the configuration of the underlying Envoy proxy and adding new functionality. The advanced options on Envoy are not exposed by default, but they can be accessed by adding an `httpGateway` section to your listener configuration. 

```yaml
apiVersion: gateway.solo.io/v1
kind: Gateway

spec:
  httpGateway:
    options:
      httpConnectionManagerSettings:
        tracing:
          verbose: true
          requestHeadersForTags:
            - path
            - origin
```

Some of the advanced options include [enabling tracing]({{% versioned_link_path fromRoot="/guides/observability/tracing/" %}}), [access log configuration]({{% versioned_link_path fromRoot="/guides/security/access_logging//" %}}), disabling [gRPC Web transcoding]({{% versioned_link_path fromRoot="/guides/traffic_management/listener_configuration/grpc_web/" %}}), and fine-grained control over [Websockets]({{% versioned_link_path fromRoot="/guides/traffic_management/listener_configuration/websockets/" %}}). More detail on how to perform advanced listener configuration can be found in the [HTTP Connection Manager guide]({{% versioned_link_path fromRoot="/guides/traffic_management/listener_configuration/http_connection_manager/" %}}).

---

## Traffic processing

Traffic that arrives at a listener is processed using one of the Virtual Services bound to the Gateway. The selection of a Virtual Service is based on the domain specified in the request. A Virtual Service contains rules regarding how a destination is selected and if the request should be altered in any way before sending it along.

### Destination selection

Routes are the primary building block of a Virtual Service. Routes contain matchers and an Upstream which could be a single destination, a list of weighted destinations, or an Upstream Group.

{{< highlight proto "hl_lines=8-15" >}}
apiVersion: gateway.solo.io/v1
kind: VirtualService

spec:
  virtualHost:
    domains:
      - 'example.com'
    routes:
      - matchers:
         - prefix: /app/cart
        routeAction:
          single:
            upstream:
              name: shopping-cart
              namespace: gloo-system
{{< /highlight >}}

Matchers inspect information about a request and determine if the data in the request matches a value defined in the rule. The content inspected can include the request path, header, query, and method. Matchers can be combined in a single rule to further refine which requests will be matched against that rule. For instance, a request could be using the POST method and reference the path `/app/cart`. The combination of an HTTP Method matcher and a Path matcher could identify the request, and send it to a shopping cart Upstream.

More information on each type of matcher is available in the following guides.

* [Path matching]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_selection/path_matching/" %}})
* [Header matching]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_selection/header_matching/" %}})
* [Query Parameter Matching]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_selection/query_parameter_matching/" %}})
* [HTTP Method Matching]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_selection/http_method_matching/" %}})

---

### Destination types

Once an incoming request has been matched by a route rule, the traffic can either be sent to a destination or processed locally. The most common destination for a route is a single Gloo Gateway Upstream. It’s also possible to route to multiple Upstreams, by either specifying multiple destinations, or by configuring an Upstream Group. Finally, it’s possible to route directly to Kubernetes or Consul services, without needing to use Gloo Gateway Upstreams or discovery.

#### Single Upstreams

Upstreams can be added manually, creating what are called [Static Upstreams]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/static_upstream//" %}}). Gloo Gateway also has a discovery service that can monitor Kubernetes or Consul and [automatically add new services]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/discovered_upstream/" %}}) as they are discovered. When routing to an Upstream, you can take advantage of Gloo Gateway’s endpoint discovery system, and configure routes to specific functions, either on a REST or gRPC service, or on a cloud function.

#### Multiple Upstreams

There may be times you want to specify multiple Upstreams for a given route. Perhaps you are performing Blue/Green testing, and want to send a certain percentage of traffic to an alternate version of a service. You can specify [multiple Upstream destinations]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/multi_destination/" %}}) in your route, [create an Upstream Group]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/upstream_groups//" %}}) for your route, or send traffic to a [subset of pods in Kubernetes]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/upstream_groups//" %}}).

Gloo Gateway can also use Upstream Groups to perform a [canary release]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/canary/" %}}), by slowly and iteratively introducing a new destination for a percentage of the traffic on a Virtual Service. Gloo Gateway can be used with [Flagger](https://docs.flagger.app/tutorials/gloo-progressive-delivery) to automatically change the percentages in an Upstream Group as part of a canary release.

In addition to static and discovered Upstreams, the following Upstreams can be created to map directly a specialty construct:

* [Kubernetes services]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/kubernetes_services/" %}})
* [Consul services]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/consul_services/" %}})
* [AWS Lambda]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/aws_lambda/" %}})
* [REST endpoint]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/rest_endpoint/" %}})
* [gRPC]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/grpc_to_rest/" %}})

#### Route Delegation

While it is possible to have a single Virtual Service directly route all traffic for a domain in its main route configuration, that may not always be the best approach. Let's say you have a domain called example.com which runs a shopping cart API and a community forum. The shopping cart API is handled by one team and the community forum is handled by another. You want to enable each team to make updates on their routing rules, without stepping on each other toes or messing with the main Virtual Service for the domain. Sounds like a job for route delegation!

In a route delegation, a prefix of the main Virtual Service can be delegated to a *Route Table*.  The Route Table is a collection of routes, just like in the main Virtual Service, but the permissions of the Route Table could be scoped to your shopping cart API team. In our example, you can create route delegations for the `/api/cart` prefix and the `/community` prefix to route tables managed by the respective teams. Now each team is free to manage their own set of routing rule, and you have the freedom to expand this model as new services are added to the domain. You can find out more in the [route delegation guide]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/delegation//" %}}).

---

## Traffic processing

Gloo Gateway can also alter requests before sending them to a destination, including transformations, fault injections, response header editing, and prefix rewrites. The ability to edit requests on the fly gives Gloo Gateway the power to specify the proper parameters for a function or transform and error check incoming requests before passing them along.

For more information, see [Traffic processing]({{% versioned_link_path fromRoot="/introduction/traffic_filter/" %}}).

---

## Configuration validation

When configuring an API gateway or edge proxy, invalid configurations can quickly lead to bugs, service outages, and security vulnerabilities. Whenever Gloo Gateway configuration objects are updated, Gloo Gateway validates and processes the new configuration. This is achieved through a four-step process:

1. Admit or reject change with a Kubernetes Validating Webhook
1. Process a batch of changes and report any errors
1. Report the status on change
1. Process the changes and apply to Envoy

More detail on the validation process and its settings can be found in the [Configuration Validation guide]({{% versioned_link_path fromRoot="/guides/traffic_management/configuration_validation/" %}}).

---

## Next Steps

Now that you have an understanding of how Gloo Gateway handles traffic management we have a few suggested paths:

* **[Security]({{% versioned_link_path fromRoot="/introduction/security/" %}})** - learn more about Gloo Gateway and its security features
* **[Setup]({{% versioned_link_path fromRoot="/installation/" %}})** - Deploy your own instance of Gloo Gateway
* **[Traffic management guides]({{% versioned_link_path fromRoot="/guides/traffic_management/" %}})** - Try out the traffic management guides to learn more

