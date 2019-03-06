# cluster-manager-api
Cluster Manager API (CMA) is used to create clusters in different cloud and on-premises environments from a single API entrypoint.

Currently supported providers:
1. [aws](https://github.com/samsung-cnct/cma-aws)
2. [azure (aks)](https://github.com/samsung-cnct/cma-aks)
3. [vmware (via ssh)](https://github.com/samsung-cnct/cma-ssh)


## Deployment
The default way to deploy CMA is by the provided helm charts located in the `deployments/helm/cluster-manager-api` directory

#### Prerequisites
1. [ingress controller](https://github.com/helm/charts/tree/master/stable/nginx-ingress)
1. [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager)

#### install via [helm](https://helm.sh/docs/using_helm/#quickstart)
1. Determine which provider(s) you would like enable (aws, aks, vmware).
1. Install helm chart:
    ```bash
    # example of enabling provider aws with default values.
    helm install deployments/helm/cluster-manager-api --name cma --set helpers.aws.enabled=true
    ```

Detailed list of helm chart values:

|Value|Description|
|:----|:----------|
|name|name of the helm release|
|image.repo|image installed including its tag|
|port|port CMA pod is listening on|
|service.port|port the CMA kubernetes service is listing on|
|service.type|service type defaults to ClusterIP|
|ingress.rest.name|name of the rest endpoint exposed as an ingress resource|
|ingress.rest.host|url of the rest endpoint exposed as an ingress resource|
|ingress.grpc.name|name of the grpc endpoint exposed as an ingress resource|
|ingress.grpc.host|url of the grpc endpoint exposed as an ingress resource|
|issuer.email|email used by cert-manager when issuing certificates|
|issuer.region|aws region used by cert-manager to issue certificates|
|issuer.accesskey.id|aws access key id used by cert-manager|
|issuer.accesskey.secret|aws access secret used by cert-manager|
|issuer.hostedzoneid|aws hostedzoneid of the domain used by cert-manager to issue certificates|
|helpers.aks.enabled|allows you to enable the aks provider API|
|helpers.aks.endpoint|service or FQDN of the aks provider including the port|
|helpers.aks.insecure|set as true when using port 80 and kubernetes service or false if using FQDN and certificates|
|helpers.aws.enabled|allows you to enable the aws provider API|
|helpers.aws.endpoint|service or FQDN of the aws provider including the port|
|helpers.aws.insecure|set as true when using port 80 and kubernetes service or false if using FQDN and certificates|
|helpers.vmware.enabled|allows you to enable the vmware provider API|
|helpers.vmware.endpoint|service or FQDN of the vmware provider including the port|
|helpers.vmware.insecure|set as true when using port 80 and kubernetes service or false if using FQDN and certificates|
