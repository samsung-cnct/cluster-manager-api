# Protocol Documentation
<a name="top"/>

## Table of Contents

- [api.proto](#api.proto)
    - [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem)
    - [ClusterItem](#cluster_manager_api.ClusterItem)
    - [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec)
    - [CreateClusterMaaSSpec](#cluster_manager_api.CreateClusterMaaSSpec)
    - [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec)
    - [CreateClusterReply](#cluster_manager_api.CreateClusterReply)
    - [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg)
    - [DeleteClusterReply](#cluster_manager_api.DeleteClusterReply)
    - [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg)
    - [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartReply)
    - [Error](#cluster_manager_api.Error)
    - [GenericHelmChart](#cluster_manager_api.GenericHelmChart)
    - [GenericTillerSetting](#cluster_manager_api.GenericTillerSetting)
    - [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg)
    - [GetClusterListReply](#cluster_manager_api.GetClusterListReply)
    - [GetClusterMsg](#cluster_manager_api.GetClusterMsg)
    - [GetClusterReply](#cluster_manager_api.GetClusterReply)
    - [GetPodCountMsg](#cluster_manager_api.GetPodCountMsg)
    - [GetPodCountReply](#cluster_manager_api.GetPodCountReply)
    - [HelloWorldMsg](#cluster_manager_api.HelloWorldMsg)
    - [HelloWorldReply](#cluster_manager_api.HelloWorldReply)
    - [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg)
    - [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartReply)
    - [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg)
    - [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerReply)
  
  
  
    - [Cluster](#cluster_manager_api.Cluster)
  

- [api.proto](#api.proto)
    - [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem)
    - [ClusterItem](#cluster_manager_api.ClusterItem)
    - [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec)
    - [CreateClusterMaaSSpec](#cluster_manager_api.CreateClusterMaaSSpec)
    - [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec)
    - [CreateClusterReply](#cluster_manager_api.CreateClusterReply)
    - [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg)
    - [DeleteClusterReply](#cluster_manager_api.DeleteClusterReply)
    - [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg)
    - [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartReply)
    - [Error](#cluster_manager_api.Error)
    - [GenericHelmChart](#cluster_manager_api.GenericHelmChart)
    - [GenericTillerSetting](#cluster_manager_api.GenericTillerSetting)
    - [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg)
    - [GetClusterListReply](#cluster_manager_api.GetClusterListReply)
    - [GetClusterMsg](#cluster_manager_api.GetClusterMsg)
    - [GetClusterReply](#cluster_manager_api.GetClusterReply)
    - [GetPodCountMsg](#cluster_manager_api.GetPodCountMsg)
    - [GetPodCountReply](#cluster_manager_api.GetPodCountReply)
    - [HelloWorldMsg](#cluster_manager_api.HelloWorldMsg)
    - [HelloWorldReply](#cluster_manager_api.HelloWorldReply)
    - [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg)
    - [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartReply)
    - [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg)
    - [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerReply)
  
  
  
    - [Cluster](#cluster_manager_api.Cluster)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cluster_manager_api.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |






<a name="cluster_manager_api.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |






<a name="cluster_manager_api.CreateClusterAWSSpec"/>

### CreateClusterAWSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | The region for AWS |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |






<a name="cluster_manager_api.CreateClusterMaaSSpec"/>

### CreateClusterMaaSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  | The MaaS API endpoint |
| username | [string](#string) |  | The username in the MaaS API |
| oauth_key | [string](#string) |  | The OAuth key for the endpoint |






<a name="cluster_manager_api.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec) |  | The provider specification |






<a name="cluster_manager_api.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - currently this is aws or maas |
| maas | [CreateClusterMaaSSpec](#cluster_manager_api.CreateClusterMaaSSpec) |  | The MaaS specification |
| aws | [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec) |  | The AWS specification |






<a name="cluster_manager_api.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cluster_manager_api.ClusterItem) |  |  |
| error | [Error](#cluster_manager_api.Error) |  |  |






<a name="cluster_manager_api.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |






<a name="cluster_manager_api.DeleteClusterReply"/>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cluster_manager_api.DeleteHelmChartMsg"/>

### DeleteHelmChartMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| tiller | [GenericTillerSetting](#cluster_manager_api.GenericTillerSetting) |  | Tiller settings |
| chart | [string](#string) |  | Chart Name |






<a name="cluster_manager_api.DeleteHelmChartReply"/>

### DeleteHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.Error"/>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | The error code |
| message | [string](#string) |  | The error message |






<a name="cluster_manager_api.GenericHelmChart"/>

### GenericHelmChart



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the name of the deployment |
| namespace | [string](#string) |  | What is the namespace to deploy the application to |
| repo | [string](#string) |  | What is the chart repository |
| chart | [string](#string) |  | What is the chart name |
| values | [string](#string) |  | What are the options (nested yaml - the Values.yaml contents) |






<a name="cluster_manager_api.GenericTillerSetting"/>

### GenericTillerSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | What is the tiller namespace |
| version | [string](#string) |  | What is the version of tiller/helm |






<a name="cluster_manager_api.GetClusterListMsg"/>

### GetClusterListMsg







<a name="cluster_manager_api.GetClusterListReply"/>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cluster_manager_api.ClusterItem) | repeated | List of clusters |






<a name="cluster_manager_api.GetClusterMsg"/>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |






<a name="cluster_manager_api.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem) |  |  |
| error | [Error](#cluster_manager_api.Error) |  |  |






<a name="cluster_manager_api.GetPodCountMsg"/>

### GetPodCountMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | Namespace to fetch the pod count Leave empty to query all namespaces |






<a name="cluster_manager_api.GetPodCountReply"/>

### GetPodCountReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pods | [int32](#int32) |  | Number of pods in the namespace (or all namespaces) |






<a name="cluster_manager_api.HelloWorldMsg"/>

### HelloWorldMsg
The Hello World request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="cluster_manager_api.HelloWorldReply"/>

### HelloWorldReply
The response to Hello World


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="cluster_manager_api.InstallHelmChartMsg"/>

### InstallHelmChartMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| tiller | [GenericTillerSetting](#cluster_manager_api.GenericTillerSetting) |  | Tiller settings |
| chart | [GenericHelmChart](#cluster_manager_api.GenericHelmChart) |  | Chart Settings |






<a name="cluster_manager_api.InstallHelmChartReply"/>

### InstallHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.ProvisionTillerMsg"/>

### ProvisionTillerMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| namespace | [string](#string) |  | Namespace tiller should be installed in |
| version | [string](#string) |  | Versino of tiller/helm to install / upgrade to |
| cluster_wide | [bool](#bool) |  | Is the tiller a cluster-wide tiller? Should it have cluster-wide admin privileges? |
| admin_namespaces | [string](#string) | repeated | Namespaces that it should be able to admin on |






<a name="cluster_manager_api.ProvisionTillerReply"/>

### ProvisionTillerReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |





 

 

 


<a name="cluster_manager_api.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| HelloWorld | [HelloWorldMsg](#cluster_manager_api.HelloWorldMsg) | [HelloWorldReply](#cluster_manager_api.HelloWorldMsg) | Simple Hello World Service - will repeat a greeting to the name provided |
| GetPodCount | [GetPodCountMsg](#cluster_manager_api.GetPodCountMsg) | [GetPodCountReply](#cluster_manager_api.GetPodCountMsg) | Simple pod count response. Will respond with the number of pods in the namespace provided |
| CreateCluster | [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg) | [CreateClusterReply](#cluster_manager_api.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cluster_manager_api.GetClusterMsg) | [GetClusterReply](#cluster_manager_api.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg) | [DeleteClusterReply](#cluster_manager_api.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg) | [GetClusterListReply](#cluster_manager_api.GetClusterListMsg) | Will retrieve a list of clusters |
| ProvisionTiller | [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg) | [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerMsg) | Will install (or reinstall) tiller |
| InstallHelmChart | [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg) | [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartMsg) | Will install (or reinstall) helm chart This will be destructive if a chart has already been deployed with the same name |
| DeleteHelmChart | [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg) | [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartMsg) | Will delete deployed helm chart |

 



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cluster_manager_api.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |






<a name="cluster_manager_api.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |






<a name="cluster_manager_api.CreateClusterAWSSpec"/>

### CreateClusterAWSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | The region for AWS |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |






<a name="cluster_manager_api.CreateClusterMaaSSpec"/>

### CreateClusterMaaSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  | The MaaS API endpoint |
| username | [string](#string) |  | The username in the MaaS API |
| oauth_key | [string](#string) |  | The OAuth key for the endpoint |






<a name="cluster_manager_api.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec) |  | The provider specification |






<a name="cluster_manager_api.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - currently this is aws or maas |
| maas | [CreateClusterMaaSSpec](#cluster_manager_api.CreateClusterMaaSSpec) |  | The MaaS specification |
| aws | [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec) |  | The AWS specification |






<a name="cluster_manager_api.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cluster_manager_api.ClusterItem) |  |  |
| error | [Error](#cluster_manager_api.Error) |  |  |






<a name="cluster_manager_api.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |






<a name="cluster_manager_api.DeleteClusterReply"/>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cluster_manager_api.DeleteHelmChartMsg"/>

### DeleteHelmChartMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| tiller | [GenericTillerSetting](#cluster_manager_api.GenericTillerSetting) |  | Tiller settings |
| chart | [string](#string) |  | Chart Name |






<a name="cluster_manager_api.DeleteHelmChartReply"/>

### DeleteHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.Error"/>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | The error code |
| message | [string](#string) |  | The error message |






<a name="cluster_manager_api.GenericHelmChart"/>

### GenericHelmChart



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the name of the deployment |
| namespace | [string](#string) |  | What is the namespace to deploy the application to |
| repo | [string](#string) |  | What is the chart repository |
| chart | [string](#string) |  | What is the chart name |
| values | [string](#string) |  | What are the options (nested yaml - the Values.yaml contents) |






<a name="cluster_manager_api.GenericTillerSetting"/>

### GenericTillerSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | What is the tiller namespace |
| version | [string](#string) |  | What is the version of tiller/helm |






<a name="cluster_manager_api.GetClusterListMsg"/>

### GetClusterListMsg







<a name="cluster_manager_api.GetClusterListReply"/>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cluster_manager_api.ClusterItem) | repeated | List of clusters |






<a name="cluster_manager_api.GetClusterMsg"/>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |






<a name="cluster_manager_api.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem) |  |  |
| error | [Error](#cluster_manager_api.Error) |  |  |






<a name="cluster_manager_api.GetPodCountMsg"/>

### GetPodCountMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | Namespace to fetch the pod count Leave empty to query all namespaces |






<a name="cluster_manager_api.GetPodCountReply"/>

### GetPodCountReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pods | [int32](#int32) |  | Number of pods in the namespace (or all namespaces) |






<a name="cluster_manager_api.HelloWorldMsg"/>

### HelloWorldMsg
The Hello World request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="cluster_manager_api.HelloWorldReply"/>

### HelloWorldReply
The response to Hello World


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="cluster_manager_api.InstallHelmChartMsg"/>

### InstallHelmChartMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| tiller | [GenericTillerSetting](#cluster_manager_api.GenericTillerSetting) |  | Tiller settings |
| chart | [GenericHelmChart](#cluster_manager_api.GenericHelmChart) |  | Chart Settings |






<a name="cluster_manager_api.InstallHelmChartReply"/>

### InstallHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.ProvisionTillerMsg"/>

### ProvisionTillerMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| namespace | [string](#string) |  | Namespace tiller should be installed in |
| version | [string](#string) |  | Versino of tiller/helm to install / upgrade to |
| cluster_wide | [bool](#bool) |  | Is the tiller a cluster-wide tiller? Should it have cluster-wide admin privileges? |
| admin_namespaces | [string](#string) | repeated | Namespaces that it should be able to admin on |






<a name="cluster_manager_api.ProvisionTillerReply"/>

### ProvisionTillerReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |





 

 

 


<a name="cluster_manager_api.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| HelloWorld | [HelloWorldMsg](#cluster_manager_api.HelloWorldMsg) | [HelloWorldReply](#cluster_manager_api.HelloWorldMsg) | Simple Hello World Service - will repeat a greeting to the name provided |
| GetPodCount | [GetPodCountMsg](#cluster_manager_api.GetPodCountMsg) | [GetPodCountReply](#cluster_manager_api.GetPodCountMsg) | Simple pod count response. Will respond with the number of pods in the namespace provided |
| CreateCluster | [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg) | [CreateClusterReply](#cluster_manager_api.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cluster_manager_api.GetClusterMsg) | [GetClusterReply](#cluster_manager_api.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg) | [DeleteClusterReply](#cluster_manager_api.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg) | [GetClusterListReply](#cluster_manager_api.GetClusterListMsg) | Will retrieve a list of clusters |
| ProvisionTiller | [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg) | [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerMsg) | Will install (or reinstall) tiller |
| InstallHelmChart | [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg) | [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartMsg) | Will install (or reinstall) helm chart This will be destructive if a chart has already been deployed with the same name |
| DeleteHelmChart | [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg) | [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartMsg) | Will delete deployed helm chart |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

