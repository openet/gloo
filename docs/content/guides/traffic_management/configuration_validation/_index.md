---
menuTitle: Configuration Validation
title: Configuration Validation
weight: 60
description: (Kubernetes Only) Gloo Gateway can be configured to validate configuration before it is applied to the cluster. With validation enabled, any attempt to apply invalid configuration to the cluster will be rejected.
---

Learn how to prevent the gateway proxies from getting invalid Gloo resource configuration. This way, you reduce the likelihood of bugs, service outages, or security vulnerabilities.

## Resource validation in Gloo Gateway

Kubernetes provides dynamic admission control capabilities that intercept requests to the Kubernetes API server and validate or mutate objects before they are persisted in etcd. Gloo Gateway leverages this capability and provides a [validation admission webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/) that ensures that only Gloo resources with valid Envoy configuration are written to etcd.

The following image shows how Gloo Gateway validates resource configuration with the validating admission webhook before applying the configuration in the cluster.

<figure><img src="{{% versioned_link_path fromRoot="/img/admission-control.svg" %}}"/>
<figcaption style="text-align:center;font-style:italic">Resource validation in Gloo Gateway</figcaption></figure>

1. The user performs an action on a Gloo custom resource that invokes the validating addmission webhook, such as to create a Gloo virtual service. To find the rules for when the webhook is being invoked, see the [validating webhook configuration](https://github.com/solo-io/gloo/blob/main/install/helm/gloo/templates/5-gateway-validation-webhook-configuration.yaml). 
2. An API request is sent to the Kubernetes API server that contains the new Gloo resource configuration. After being successfully authenticated, authorized, and processed by any mutating admission webhooks that are configured in the cluster, the configuration is sent to the object schema validation component. This component performs an OpenAPI schema validation to verify that the provided YAML or JSON configuration is valid. If the configuration is not written in valid YAML or JSON format, it is rejected by the Kubernetes API server. For example, if you upgraded your Gloo version, but did not apply the matching Gloo custom resources, you might see errors similar to the following when fields in the custom resource configuration are unknown. 
   ```
   Error: UPGRADE FAILED: error validating "": error validating data: ValidationError(Settings.spec.gateway.validation): unknown field "validationServerGrpcMaxSizeBytes" in io.solo.gloo.v1.Settings.spec.gateway.validation
   ```

3. If the resource configuration passes the schema validation, the configuration is sent to the validation webhook server that is configured with the Gloo Gateway validating admission webhook for semantic validation. 
4. The webhook kicks off the processes within the translation engine component to simulate the translation and xDS snapshot creation in Gloo Gateway. 
5. The translation engine retrieves the current Gloo Gateway snapshot and compares it with the changes in the new Gloo resource configuration. 
6. Gloo Gateway tries to translate the resource into valid Envoy configuration. 
7. Gloo Gateway tries to process the Envoy configuration along with service discovery data to create the final xDS snapshot. 
8. If the configuration was succesfully processed, Gloo Gateway creates the xDS snapshot. 
9. The validation result is returned to the validation admission server. 
10. The validating admission webhook accepts or rejects the configuration, depending on the following modes. For more information about how to configure and test the webhook, see the [admission control]({{% versioned_link_path fromRoot="/guides/traffic_management/configuration_validation/admission_control/" %}}) guide.
    * Permissive mode (default): By default, the validating admission webhook is set up in permissive mode and only logs invalid resource configurations without rejecting them. 
    * Strict mode: You can [enable strict validation]({{% versioned_link_path fromRoot="/guides/traffic_management/configuration_validation/admission_control/#enable-strict-resource-validation" %}}) to reject invalid resource configuration before it is stored in etcd.
    * Strict mode with warnings: You can also configure the webhook to reject resource configuration that resulted in a `warning` status. 
11. If the resource configuration is found to be schematically and semantically correct, it is admitted by the validation webhook server and persisted in the etcd data store.
12. Gloo Gateway monitors the status of custom resources and picks up the latest Gloo custom resource changes from the etcd data store. 
13. The resource configuration runs through the translation engine processes as described in steps 5-8. The results depend on the translation and settings as follows.
    * Successful translation: The resource's status is changed to `Accepted` and the configuration is admitted to be part of the xDS snapshot. 
    * Unsuccessful translation: If errors are detected during the translation process, Gloo Gateway sets the resource's {{< protobuf name="core.solo.io.Status" display="Status">}} to `Rejected` and omits the resource from the final xDS snapshot. 
    * Last valid config: For previously admitted resources that become invalid due to an update, the last valid configuration of that resource is persisted in etcd. 
    * Warnings: If the resource configuration is considered valid, but references missing or misconfigured resources, the validation status is set to `Warning` and the resource is prepared to be added to the xDS snapshot. 
14. Gloo Gateway processes the resource configuration along with service discovery data to create the final [Envoy xDS Snapshot](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol). The xDS snapshot is sent to the Gloo xDS server component. 
15. Proxies in the cluster can pull the latest Envoy configuration from the Gloo xDS server. 

For more information about the validating admission webhook configuration, see the [admission control]({{% versioned_link_path fromRoot="/guides/traffic_management/configuration_validation/admission_control/" %}}) guide. 


## Sanitize resource configuration

Gloo Gateway can be configured to pass partially-valid configuration to Envoy by admitting it through an internal process referred to as *sanitizing*. Rather than refusing to update Envoy with invalid configuration, Gloo Gateway can replace invalid configuration with preconfigured defaults.

For more information about how to configure and use the Gloo Gateway sanitization feature, see the [Route Replacement]({{% versioned_link_path fromRoot="/guides/traffic_management/configuration_validation/invalid_route_replacement/" %}}) guide. 

