---
title: External Auth Custom Cert Authority
weight: 80
description: Configuring a custom certificate authority for extauth to use.
---

Gloo Gateway Enterprise includes external authentication, which allows you to offload authentication responsibilities from Envoy to an external authentication server. There may be cases where you need the external authentication server to trust certificates issued from a custom certificate authority. In this guide, we will show you how to add the certificate authority during Gloo Gateway Enterprise installation or after installation is complete.

The external authentication server runs as its own Kubernetes pod or as a sidecar to the `gateway-proxy` pods. The certificate authority public certificate will be saved as Kubernetes secret, and then an initialization container will be used to inject the CA certificate into the list of trusted certificate authorities for the external authentication pods. 

For this guide, we will create a temporary certificate authority using OpenSSL. In a production scenario, you would retrieve the public certificate from an existing certificate authority you wish to be trusted.

This guide assumes that you already have a Kubernetes cluster available for installation of Gloo Gateway Enterprise, or that you have a running instance of Gloo Gateway Enterprise.

## Create a certificate authority

We are going to use OpenSSL to create a simple certificate authority and upload the public certificate as a Kubernetes secret. First let's create the certificate authority:

```bash
# Enter whatever passphrase you'd like
openssl genrsa -des3 -out ca.key 4096

# Enter glooe.example.com for the Common Name, leave all other defaults
openssl req -new -x509 -days 365 -key ca.key -out ca.cert.pem
```

Now we are a certificate authority! Let's go ahead and get the `ca.cert.pem` file added as a Kubernetes secret in our cluster.

```bash
# Create the gloo-system namepace if it doesn't exist
kubectl create namespace gloo-system

# Add the CA pem file as a generic secret
kubectl create secret generic trusted-ca --from-file=tls.crt=ca.cert.pem -n gloo-system
```

Now we are ready to either [install Gloo Gateway Enterprise](#install-gloo-edge-enterprise) or [update an existing Gloo Gateway Enterprise installation](#update-gloo-edge-enterprise).

## Install Gloo Gateway Enterprise

To add the customization of a trusted certificate authority to the Gloo Gateway Enterprise installation, we are going to need to use Helm for the installation and customization.

```bash
# Add the Gloo Gateway Enterprise repo to Helm if you haven't already
helm repo add glooe https://storage.googleapis.com/gloo-ee-helm
helm repo update

# Create the helm values file (or merge into existing)
cat > gloo-edge-bring-cert-values.yaml <<EOF
global:
  extensions:
    extAuth:
      deployment:
        extraVolume:
        - name: ca-certs-custom
          secret:
            secretName: trusted-ca
        extraVolumeMount:
        - name: ca-certs-custom
          mountPath: /etc/ssl/certs/ca-certs-custom.crt
          subPath: tls.crt
          readOnly: true
EOF
```

Finally, we'll install Gloo Gateway Enterprise with Helm. Be sure to update the value for the license key.
Include the `--install` flag to upgrade the existing installation or install a new release if one does not already exist.

```bash
helm upgrade --install gloo glooe/gloo-ee --namespace gloo-system \
  --set-string license_key=LICENSE_KEY \
  -f gloo-edge-bring-cert-values.yaml
```

Once the installation is complete, we can validate our change with the following command:

```bash
kubectl exec -n gloo-system deploy/extauth -- cat /etc/ssl/certs/ca-certs-custom.crt
```

You've successfully added a custom certificate authority for external authentication!

## Summary

In this guide you learned how to add a custom certificate authority as trusted to the external authentication server. If you want to know more about the capabilities of external authentication, be sure to check out our [guides]({{< versioned_link_path fromRoot="/guides/security/auth/extauth/" >}}).
