# Protocol Documentation
<a name="top"/>

## Table of Contents

- [api.proto](#api.proto)
    - [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem)
    - [ClusterItem](#cluster_manager_api.ClusterItem)
    - [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec)
    - [CreateClusterAWSSpec.AWSCredentials](#cluster_manager_api.CreateClusterAWSSpec.AWSCredentials)
    - [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter)
    - [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup)
    - [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems)
    - [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec)
    - [CreateClusterReply](#cluster_manager_api.CreateClusterReply)
    - [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg)
    - [DeleteClusterReply](#cluster_manager_api.DeleteClusterReply)
    - [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg)
    - [GetClusterListReply](#cluster_manager_api.GetClusterListReply)
    - [GetClusterMsg](#cluster_manager_api.GetClusterMsg)
    - [GetClusterReply](#cluster_manager_api.GetClusterReply)
    - [GetVersionMsg](#cluster_manager_api.GetVersionMsg)
    - [GetVersionReply](#cluster_manager_api.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cluster_manager_api.GetVersionReply.VersionInformation)
  
  
  
    - [Cluster](#cluster_manager_api.Cluster)
  

- [api.proto](#api.proto)
    - [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem)
    - [ClusterItem](#cluster_manager_api.ClusterItem)
    - [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec)
    - [CreateClusterAWSSpec.AWSCredentials](#cluster_manager_api.CreateClusterAWSSpec.AWSCredentials)
    - [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter)
    - [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup)
    - [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems)
    - [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec)
    - [CreateClusterReply](#cluster_manager_api.CreateClusterReply)
    - [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg)
    - [DeleteClusterReply](#cluster_manager_api.DeleteClusterReply)
    - [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg)
    - [GetClusterListReply](#cluster_manager_api.GetClusterListReply)
    - [GetClusterMsg](#cluster_manager_api.GetClusterMsg)
    - [GetClusterReply](#cluster_manager_api.GetClusterReply)
    - [GetVersionMsg](#cluster_manager_api.GetVersionMsg)
    - [GetVersionReply](#cluster_manager_api.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cluster_manager_api.GetVersionReply.VersionInformation)
  
  
  
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
| data_center | [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter) |  | The AWS Data Center |
| credentials | [CreateClusterAWSSpec.AWSCredentials](#cluster_manager_api.CreateClusterAWSSpec.AWSCredentials) |  | Credentials to build the cluster |
| resources | [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems) |  | BYO items |
| instance_groups | [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup) | repeated | Instance groups |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSCredentials"/>

### CreateClusterAWSSpec.AWSCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter"/>

### CreateClusterAWSSpec.AWSDataCenter
Which Data Center


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | Which region (us-east-1, etc.) |
| availability_zones | [string](#string) | repeated | Which availability zones (us-east-1b, us-east-2c, us-west-2d, etc.) |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup"/>

### CreateClusterAWSSpec.AWSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Instance type (m5.large, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems"/>

### CreateClusterAWSSpec.AWSPreconfiguredItems
For when some things are already created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpc_id | [string](#string) |  | The VPC id, blank for for &#34;create one for you&#34;, filled if you are BYO VPC |
| security_group_id | [string](#string) |  | Security group |
| iam_role_arn | [string](#string) |  | The IAM role for the cluster (arn) |






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
| k8s_version | [string](#string) |  | The version of Kubernetes |
| aws | [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec) |  | The AWS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cluster_manager_api.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cluster_manager_api.ClusterItem) |  | The details of the cluster request response |






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






<a name="cluster_manager_api.GetVersionMsg"/>

### GetVersionMsg
Get version of API Server






<a name="cluster_manager_api.GetVersionReply"/>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cluster_manager_api.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cluster_manager_api.GetVersionReply.VersionInformation"/>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 

 

 


<a name="cluster_manager_api.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg) | [CreateClusterReply](#cluster_manager_api.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cluster_manager_api.GetClusterMsg) | [GetClusterReply](#cluster_manager_api.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg) | [DeleteClusterReply](#cluster_manager_api.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg) | [GetClusterListReply](#cluster_manager_api.GetClusterListMsg) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cluster_manager_api.GetVersionMsg) | [GetVersionReply](#cluster_manager_api.GetVersionMsg) | Will return version information about api server |

 



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
| data_center | [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter) |  | The AWS Data Center |
| credentials | [CreateClusterAWSSpec.AWSCredentials](#cluster_manager_api.CreateClusterAWSSpec.AWSCredentials) |  | Credentials to build the cluster |
| resources | [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems) |  | BYO items |
| instance_groups | [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup) | repeated | Instance groups |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSCredentials"/>

### CreateClusterAWSSpec.AWSCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter"/>

### CreateClusterAWSSpec.AWSDataCenter
Which Data Center


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | Which region (us-east-1, etc.) |
| availability_zones | [string](#string) | repeated | Which availability zones (us-east-1b, us-east-2c, us-west-2d, etc.) |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup"/>

### CreateClusterAWSSpec.AWSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Instance type (m5.large, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |






<a name="cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems"/>

### CreateClusterAWSSpec.AWSPreconfiguredItems
For when some things are already created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpc_id | [string](#string) |  | The VPC id, blank for for &#34;create one for you&#34;, filled if you are BYO VPC |
| security_group_id | [string](#string) |  | Security group |
| iam_role_arn | [string](#string) |  | The IAM role for the cluster (arn) |






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
| k8s_version | [string](#string) |  | The version of Kubernetes |
| aws | [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec) |  | The AWS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cluster_manager_api.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cluster_manager_api.ClusterItem) |  | The details of the cluster request response |






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






<a name="cluster_manager_api.GetVersionMsg"/>

### GetVersionMsg
Get version of API Server






<a name="cluster_manager_api.GetVersionReply"/>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cluster_manager_api.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cluster_manager_api.GetVersionReply.VersionInformation"/>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 

 

 


<a name="cluster_manager_api.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg) | [CreateClusterReply](#cluster_manager_api.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cluster_manager_api.GetClusterMsg) | [GetClusterReply](#cluster_manager_api.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg) | [DeleteClusterReply](#cluster_manager_api.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg) | [GetClusterListReply](#cluster_manager_api.GetClusterListMsg) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cluster_manager_api.GetVersionMsg) | [GetVersionReply](#cluster_manager_api.GetVersionMsg) | Will return version information about api server |

 



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

