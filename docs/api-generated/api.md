# Protocol Documentation
<a name="top"/>

## Table of Contents

- [api.proto](#api.proto)
    - [AWSCredentials](#cluster_manager_api.AWSCredentials)
    - [AdjustClusterMsg](#cluster_manager_api.AdjustClusterMsg)
    - [AdjustClusterMsg.AdjustClusterAKSSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterAKSSpec)
    - [AdjustClusterMsg.AdjustClusterSshSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterSshSpec)
    - [AdjustClusterMsg.AdjustClusterVMWareSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterVMWareSpec)
    - [AdjustClusterMsg.SshRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.SshRemoveMachineSpec)
    - [AdjustClusterMsg.VMWareRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.VMWareRemoveMachineSpec)
    - [AdjustClusterReply](#cluster_manager_api.AdjustClusterReply)
    - [AzureClusterServiceAccount](#cluster_manager_api.AzureClusterServiceAccount)
    - [AzureCredentials](#cluster_manager_api.AzureCredentials)
    - [Callback](#cluster_manager_api.Callback)
    - [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem)
    - [ClusterItem](#cluster_manager_api.ClusterItem)
    - [CreateClusterAKSSpec](#cluster_manager_api.CreateClusterAKSSpec)
    - [CreateClusterAKSSpec.AKSInstanceGroup](#cluster_manager_api.CreateClusterAKSSpec.AKSInstanceGroup)
    - [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec)
    - [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter)
    - [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup)
    - [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems)
    - [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec)
    - [CreateClusterReply](#cluster_manager_api.CreateClusterReply)
    - [CreateClusterSshSpec](#cluster_manager_api.CreateClusterSshSpec)
    - [CreateClusterVMWareSpec](#cluster_manager_api.CreateClusterVMWareSpec)
    - [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg)
    - [DeleteClusterReply](#cluster_manager_api.DeleteClusterReply)
    - [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg)
    - [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartReply)
    - [GenericHelmChart](#cluster_manager_api.GenericHelmChart)
    - [GenericHelmChart.ChartRepository](#cluster_manager_api.GenericHelmChart.ChartRepository)
    - [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg)
    - [GetClusterListReply](#cluster_manager_api.GetClusterListReply)
    - [GetClusterMsg](#cluster_manager_api.GetClusterMsg)
    - [GetClusterReply](#cluster_manager_api.GetClusterReply)
    - [GetUpgradeClusterInformationMsg](#cluster_manager_api.GetUpgradeClusterInformationMsg)
    - [GetUpgradeClusterInformationReply](#cluster_manager_api.GetUpgradeClusterInformationReply)
    - [GetVersionMsg](#cluster_manager_api.GetVersionMsg)
    - [GetVersionReply](#cluster_manager_api.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cluster_manager_api.GetVersionReply.VersionInformation)
    - [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg)
    - [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartReply)
    - [KubernetesLabel](#cluster_manager_api.KubernetesLabel)
    - [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg)
    - [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerReply)
    - [SshMachineSpec](#cluster_manager_api.SshMachineSpec)
    - [UpdateAWSCredentialsMsg](#cluster_manager_api.UpdateAWSCredentialsMsg)
    - [UpdateAWSCredentialsReply](#cluster_manager_api.UpdateAWSCredentialsReply)
    - [UpdateAzureCredentialsMsg](#cluster_manager_api.UpdateAzureCredentialsMsg)
    - [UpdateAzureCredentialsReply](#cluster_manager_api.UpdateAzureCredentialsReply)
    - [UpgradeClusterMsg](#cluster_manager_api.UpgradeClusterMsg)
    - [UpgradeClusterReply](#cluster_manager_api.UpgradeClusterReply)
    - [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec)
  
    - [ClusterStatus](#cluster_manager_api.ClusterStatus)
    - [Provider](#cluster_manager_api.Provider)
  
  
    - [Cluster](#cluster_manager_api.Cluster)
  

- [api.proto](#api.proto)
    - [AWSCredentials](#cluster_manager_api.AWSCredentials)
    - [AdjustClusterMsg](#cluster_manager_api.AdjustClusterMsg)
    - [AdjustClusterMsg.AdjustClusterAKSSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterAKSSpec)
    - [AdjustClusterMsg.AdjustClusterSshSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterSshSpec)
    - [AdjustClusterMsg.AdjustClusterVMWareSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterVMWareSpec)
    - [AdjustClusterMsg.SshRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.SshRemoveMachineSpec)
    - [AdjustClusterMsg.VMWareRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.VMWareRemoveMachineSpec)
    - [AdjustClusterReply](#cluster_manager_api.AdjustClusterReply)
    - [AzureClusterServiceAccount](#cluster_manager_api.AzureClusterServiceAccount)
    - [AzureCredentials](#cluster_manager_api.AzureCredentials)
    - [Callback](#cluster_manager_api.Callback)
    - [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem)
    - [ClusterItem](#cluster_manager_api.ClusterItem)
    - [CreateClusterAKSSpec](#cluster_manager_api.CreateClusterAKSSpec)
    - [CreateClusterAKSSpec.AKSInstanceGroup](#cluster_manager_api.CreateClusterAKSSpec.AKSInstanceGroup)
    - [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec)
    - [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter)
    - [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup)
    - [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems)
    - [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec)
    - [CreateClusterReply](#cluster_manager_api.CreateClusterReply)
    - [CreateClusterSshSpec](#cluster_manager_api.CreateClusterSshSpec)
    - [CreateClusterVMWareSpec](#cluster_manager_api.CreateClusterVMWareSpec)
    - [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg)
    - [DeleteClusterReply](#cluster_manager_api.DeleteClusterReply)
    - [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg)
    - [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartReply)
    - [GenericHelmChart](#cluster_manager_api.GenericHelmChart)
    - [GenericHelmChart.ChartRepository](#cluster_manager_api.GenericHelmChart.ChartRepository)
    - [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg)
    - [GetClusterListReply](#cluster_manager_api.GetClusterListReply)
    - [GetClusterMsg](#cluster_manager_api.GetClusterMsg)
    - [GetClusterReply](#cluster_manager_api.GetClusterReply)
    - [GetUpgradeClusterInformationMsg](#cluster_manager_api.GetUpgradeClusterInformationMsg)
    - [GetUpgradeClusterInformationReply](#cluster_manager_api.GetUpgradeClusterInformationReply)
    - [GetVersionMsg](#cluster_manager_api.GetVersionMsg)
    - [GetVersionReply](#cluster_manager_api.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cluster_manager_api.GetVersionReply.VersionInformation)
    - [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg)
    - [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartReply)
    - [KubernetesLabel](#cluster_manager_api.KubernetesLabel)
    - [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg)
    - [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerReply)
    - [SshMachineSpec](#cluster_manager_api.SshMachineSpec)
    - [UpdateAWSCredentialsMsg](#cluster_manager_api.UpdateAWSCredentialsMsg)
    - [UpdateAWSCredentialsReply](#cluster_manager_api.UpdateAWSCredentialsReply)
    - [UpdateAzureCredentialsMsg](#cluster_manager_api.UpdateAzureCredentialsMsg)
    - [UpdateAzureCredentialsReply](#cluster_manager_api.UpdateAzureCredentialsReply)
    - [UpgradeClusterMsg](#cluster_manager_api.UpgradeClusterMsg)
    - [UpgradeClusterReply](#cluster_manager_api.UpgradeClusterReply)
    - [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec)
  
    - [ClusterStatus](#cluster_manager_api.ClusterStatus)
    - [Provider](#cluster_manager_api.Provider)
  
  
    - [Cluster](#cluster_manager_api.Cluster)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cluster_manager_api.AWSCredentials"/>

### AWSCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |
| region | [string](#string) |  | The Region for API access |






<a name="cluster_manager_api.AdjustClusterMsg"/>

### AdjustClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AdjustClusterMsg.AdjustClusterAKSSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterAKSSpec) |  | The AWS specification AdjustClusterAWSSpec aws = 5; |
| vmware | [AdjustClusterMsg.AdjustClusterVMWareSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterVMWareSpec) |  |  |
| ssh | [AdjustClusterMsg.AdjustClusterSshSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterSshSpec) |  |  |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.AdjustClusterMsg.AdjustClusterAKSSpec"/>

### AdjustClusterMsg.AdjustClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Credentials to build the cluster |
| node_pool | [string](#string) |  | Node Pool Name |
| add_count | [int32](#int32) |  | umber to increase by |
| remove_count | [int32](#int32) |  | number to decrease by |






<a name="cluster_manager_api.AdjustClusterMsg.AdjustClusterSshSpec"/>

### AdjustClusterMsg.AdjustClusterSshSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| add_nodes | [SshMachineSpec](#cluster_manager_api.SshMachineSpec) | repeated | Machines which we want to add to the cluster |
| remove_nodes | [AdjustClusterMsg.SshRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.SshRemoveMachineSpec) | repeated | Machines which we want to remove from the cluster |






<a name="cluster_manager_api.AdjustClusterMsg.AdjustClusterVMWareSpec"/>

### AdjustClusterMsg.AdjustClusterVMWareSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| add_nodes | [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec) | repeated | Machines which we want to add to the cluster |
| remove_nodes | [AdjustClusterMsg.VMWareRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.VMWareRemoveMachineSpec) | repeated | Machines which we want to remove from the cluster |






<a name="cluster_manager_api.AdjustClusterMsg.SshRemoveMachineSpec"/>

### AdjustClusterMsg.SshRemoveMachineSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  |  |






<a name="cluster_manager_api.AdjustClusterMsg.VMWareRemoveMachineSpec"/>

### AdjustClusterMsg.VMWareRemoveMachineSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  | The host for SSH access |






<a name="cluster_manager_api.AdjustClusterReply"/>

### AdjustClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.AzureClusterServiceAccount"/>

### AzureClusterServiceAccount
the account used by the cluster to create azure resources (ex: load balancer)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [string](#string) |  | The ClientId (aka: AppID) |
| client_secret | [string](#string) |  | The ClientSecret (aka: password) |






<a name="cluster_manager_api.AzureCredentials"/>

### AzureCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  | The AppId for API Access |
| tenant | [string](#string) |  | The Tenant for API access |
| password | [string](#string) |  | The Password for API access |
| subscription_id | [string](#string) |  | The Subscription for API access |






<a name="cluster_manager_api.Callback"/>

### Callback



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | The URL to call back to |
| request_id | [string](#string) |  | The ID of the request |






<a name="cluster_manager_api.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |
| status | [ClusterStatus](#cluster_manager_api.ClusterStatus) |  | The status of the cluster |
| bearertoken | [string](#string) |  | What is the admin bearer token |






<a name="cluster_manager_api.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| status | [ClusterStatus](#cluster_manager_api.ClusterStatus) |  | The status of the cluster |






<a name="cluster_manager_api.CreateClusterAKSSpec"/>

### CreateClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [string](#string) |  | The Azure Data Center |
| credentials | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Credentials to build the cluster |
| instance_groups | [CreateClusterAKSSpec.AKSInstanceGroup](#cluster_manager_api.CreateClusterAKSSpec.AKSInstanceGroup) | repeated | Instance groups |






<a name="cluster_manager_api.CreateClusterAKSSpec.AKSInstanceGroup"/>

### CreateClusterAKSSpec.AKSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the group |
| type | [string](#string) |  | Instance type (Standard_D2_v2, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |






<a name="cluster_manager_api.CreateClusterAWSSpec"/>

### CreateClusterAWSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data_center | [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter) |  | The AWS Data Center |
| credentials | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | Credentials to build the cluster |
| resources | [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems) |  | BYO items |
| instance_groups | [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup) | repeated | Instance groups |






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
| iam_role_arn | [string](#string) |  | The IAM role for the cluster (arn)ClusterAssociationdd |






<a name="cluster_manager_api.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec) |  | The provider specification |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - currently this is aws or maas |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| aws | [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec) |  | The AWS specification |
| azure | [CreateClusterAKSSpec](#cluster_manager_api.CreateClusterAKSSpec) |  |  |
| vmware | [CreateClusterVMWareSpec](#cluster_manager_api.CreateClusterVMWareSpec) |  |  |
| ssh | [CreateClusterSshSpec](#cluster_manager_api.CreateClusterSshSpec) |  |  |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cluster_manager_api.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cluster_manager_api.ClusterItem) |  | The details of the cluster request response |






<a name="cluster_manager_api.CreateClusterSshSpec"/>

### CreateClusterSshSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| control_plane_nodes | [SshMachineSpec](#cluster_manager_api.SshMachineSpec) | repeated | Machines which comprise the cluster |
| worker_nodes | [SshMachineSpec](#cluster_manager_api.SshMachineSpec) | repeated | Machines which comprise the cluster |
| api_endpoint | [string](#string) |  | This should be a value like ip:port that will be a virtual IP/port Passed back to external customers to be able to communicate to the cluster |






<a name="cluster_manager_api.CreateClusterVMWareSpec"/>

### CreateClusterVMWareSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| control_plane_nodes | [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec) | repeated | Machines which comprise the cluster |
| worker_nodes | [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec) | repeated | Machines which comprise the cluster |
| api_endpoint | [string](#string) |  | This should be a value like ip:port that will be a virtual IP/port Passed back to external customers to be able to communicate to the cluster |






<a name="cluster_manager_api.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






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
| package_manager | [string](#string) |  | What tiller should be used |
| chart | [string](#string) |  | Chart Name |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| request_id | [string](#string) |  | A unique id to indicate this request |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.DeleteHelmChartReply"/>

### DeleteHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.GenericHelmChart"/>

### GenericHelmChart



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the name of the deployment |
| namespace | [string](#string) |  | What is the namespace to deploy the application to |
| repo | [GenericHelmChart.ChartRepository](#cluster_manager_api.GenericHelmChart.ChartRepository) |  | What is the chart repository |
| chart | [string](#string) |  | What is the chart name |
| version | [string](#string) |  | What is the chart version |
| values | [string](#string) |  | What are the options (nested yaml - the Values.yaml contents) |
| chart_payload | [bytes](#bytes) |  | The payload of a chart (for airgapped solutions, etc) |






<a name="cluster_manager_api.GenericHelmChart.ChartRepository"/>

### GenericHelmChart.ChartRepository



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | What is the URL for the chart repo |
| name | [string](#string) |  | What is the repo name |






<a name="cluster_manager_api.GetClusterListMsg"/>

### GetClusterListMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |






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
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |






<a name="cluster_manager_api.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem) |  |  |






<a name="cluster_manager_api.GetUpgradeClusterInformationMsg"/>

### GetUpgradeClusterInformationMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |






<a name="cluster_manager_api.GetUpgradeClusterInformationReply"/>

### GetUpgradeClusterInformationReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Can the cluster be upgraded |
| versions | [string](#string) | repeated | What versions are possible right now |






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






<a name="cluster_manager_api.InstallHelmChartMsg"/>

### InstallHelmChartMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| package_manger | [string](#string) |  | What tiller should be used |
| chart | [GenericHelmChart](#cluster_manager_api.GenericHelmChart) |  | Chart Settings |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.InstallHelmChartReply"/>

### InstallHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.KubernetesLabel"/>

### KubernetesLabel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of a label |
| value | [string](#string) |  | The value of a label |






<a name="cluster_manager_api.ProvisionTillerMsg"/>

### ProvisionTillerMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the tiller / package manager |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| namespace | [string](#string) |  | Namespace tiller should be installed in |
| image | [string](#string) |  | image (if not the default) for tiller (quay.io/someguy/tiller) |
| version | [string](#string) |  | Version of tiller/helm to install / upgrade to (v2.11.0, etc) |
| cluster_wide | [bool](#bool) |  | Is the tiller a cluster-wide tiller? Should it have cluster-wide admin privileges? |
| admin_namespaces | [string](#string) | repeated | Namespaces that it should be able to admin on |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.ProvisionTillerReply"/>

### ProvisionTillerReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.SshMachineSpec"/>

### SshMachineSpec
The specification for a specific node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | The username for SSH access |
| host | [string](#string) |  | The host for SSH access |
| port | [int32](#int32) |  | The port for SSH access |
| password | [string](#string) |  | The k8s version for the control plane. This node is only a master if this field is defined. |
| labels | [KubernetesLabel](#cluster_manager_api.KubernetesLabel) | repeated | The labels for the machines |






<a name="cluster_manager_api.UpdateAWSCredentialsMsg"/>

### UpdateAWSCredentialsMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Cluster name |
| credentials | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | Credentials to build the cluster |






<a name="cluster_manager_api.UpdateAWSCredentialsReply"/>

### UpdateAWSCredentialsReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.UpdateAzureCredentialsMsg"/>

### UpdateAzureCredentialsMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Cluster name |
| credentials | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Credentials to build the cluster |






<a name="cluster_manager_api.UpdateAzureCredentialsReply"/>

### UpdateAzureCredentialsReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.UpgradeClusterMsg"/>

### UpgradeClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |
| version | [string](#string) |  | What version are we upgrading to? |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.UpgradeClusterReply"/>

### UpgradeClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.VMWareMachineSpec"/>

### VMWareMachineSpec
The specification for a specific node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | The username for SSH access |
| host | [string](#string) |  | The host for SSH access |
| port | [int32](#int32) |  | The port for SSH access |
| password | [string](#string) |  | The k8s version for the control plane. This node is only a master if this field is defined. |
| labels | [KubernetesLabel](#cluster_manager_api.KubernetesLabel) | repeated | The labels for the machines |





 


<a name="cluster_manager_api.ClusterStatus"/>

### ClusterStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | Not set |
| PROVISIONING | 1 | The PROVISIONING state indicates the cluster is being created. |
| RUNNING | 2 | The RUNNING state indicates the cluster has been created and is fully usable. |
| RECONCILING | 3 | The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software. |
| STOPPING | 4 | The STOPPING state indicates the cluster is being deleted |
| ERROR | 5 | The ERROR state indicates the cluster may be unusable |
| DEGRADED | 6 | The DEGRADED state indicates the cluster requires user action to restore full functionality |



<a name="cluster_manager_api.Provider"/>

### Provider


| Name | Number | Description |
| ---- | ------ | ----------- |
| undefined | 0 |  |
| aws | 1 |  |
| azure | 2 |  |
| vmware | 3 |  |
| ssh | 4 |  |


 

 


<a name="cluster_manager_api.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg) | [CreateClusterReply](#cluster_manager_api.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cluster_manager_api.GetClusterMsg) | [GetClusterReply](#cluster_manager_api.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg) | [DeleteClusterReply](#cluster_manager_api.DeleteClusterMsg) | Will delete a cluster |
| AdjustClusterNodes | [AdjustClusterMsg](#cluster_manager_api.AdjustClusterMsg) | [AdjustClusterReply](#cluster_manager_api.AdjustClusterMsg) | Will adjust a provision a cluster |
| GetClusterList | [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg) | [GetClusterListReply](#cluster_manager_api.GetClusterListMsg) | Will retrieve a list of clusters |
| ProvisionTiller | [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg) | [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerMsg) | Will install (or reinstall) tiller |
| InstallHelmChart | [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg) | [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartMsg) | Will install (or reinstall) helm chart This will be destructive if a chart has already been deployed with the same name |
| DeleteHelmChart | [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg) | [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartMsg) | Will delete deployed helm chart |
| GetVersionInformation | [GetVersionMsg](#cluster_manager_api.GetVersionMsg) | [GetVersionReply](#cluster_manager_api.GetVersionMsg) | Will return version information about api server |
| GetUpgradeClusterInformation | [GetUpgradeClusterInformationMsg](#cluster_manager_api.GetUpgradeClusterInformationMsg) | [GetUpgradeClusterInformationReply](#cluster_manager_api.GetUpgradeClusterInformationMsg) | Will return upgrade options for a given cluster |
| UpgradeCluster | [UpgradeClusterMsg](#cluster_manager_api.UpgradeClusterMsg) | [UpgradeClusterReply](#cluster_manager_api.UpgradeClusterMsg) | Will attempt to upgrade a cluster |
| UpdateAWSCredentials | [UpdateAWSCredentialsMsg](#cluster_manager_api.UpdateAWSCredentialsMsg) | [UpdateAWSCredentialsReply](#cluster_manager_api.UpdateAWSCredentialsMsg) | Will update aws credentials used for a cluster |
| UpdateAzureCredentials | [UpdateAzureCredentialsMsg](#cluster_manager_api.UpdateAzureCredentialsMsg) | [UpdateAzureCredentialsReply](#cluster_manager_api.UpdateAzureCredentialsMsg) | Will update azure credentials used for a cluster |

 



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cluster_manager_api.AWSCredentials"/>

### AWSCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |
| region | [string](#string) |  | The Region for API access |






<a name="cluster_manager_api.AdjustClusterMsg"/>

### AdjustClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AdjustClusterMsg.AdjustClusterAKSSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterAKSSpec) |  | The AWS specification AdjustClusterAWSSpec aws = 5; |
| vmware | [AdjustClusterMsg.AdjustClusterVMWareSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterVMWareSpec) |  |  |
| ssh | [AdjustClusterMsg.AdjustClusterSshSpec](#cluster_manager_api.AdjustClusterMsg.AdjustClusterSshSpec) |  |  |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.AdjustClusterMsg.AdjustClusterAKSSpec"/>

### AdjustClusterMsg.AdjustClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Credentials to build the cluster |
| node_pool | [string](#string) |  | Node Pool Name |
| add_count | [int32](#int32) |  | umber to increase by |
| remove_count | [int32](#int32) |  | number to decrease by |






<a name="cluster_manager_api.AdjustClusterMsg.AdjustClusterSshSpec"/>

### AdjustClusterMsg.AdjustClusterSshSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| add_nodes | [SshMachineSpec](#cluster_manager_api.SshMachineSpec) | repeated | Machines which we want to add to the cluster |
| remove_nodes | [AdjustClusterMsg.SshRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.SshRemoveMachineSpec) | repeated | Machines which we want to remove from the cluster |






<a name="cluster_manager_api.AdjustClusterMsg.AdjustClusterVMWareSpec"/>

### AdjustClusterMsg.AdjustClusterVMWareSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| add_nodes | [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec) | repeated | Machines which we want to add to the cluster |
| remove_nodes | [AdjustClusterMsg.VMWareRemoveMachineSpec](#cluster_manager_api.AdjustClusterMsg.VMWareRemoveMachineSpec) | repeated | Machines which we want to remove from the cluster |






<a name="cluster_manager_api.AdjustClusterMsg.SshRemoveMachineSpec"/>

### AdjustClusterMsg.SshRemoveMachineSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  |  |






<a name="cluster_manager_api.AdjustClusterMsg.VMWareRemoveMachineSpec"/>

### AdjustClusterMsg.VMWareRemoveMachineSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  | The host for SSH access |






<a name="cluster_manager_api.AdjustClusterReply"/>

### AdjustClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.AzureClusterServiceAccount"/>

### AzureClusterServiceAccount
the account used by the cluster to create azure resources (ex: load balancer)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [string](#string) |  | The ClientId (aka: AppID) |
| client_secret | [string](#string) |  | The ClientSecret (aka: password) |






<a name="cluster_manager_api.AzureCredentials"/>

### AzureCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  | The AppId for API Access |
| tenant | [string](#string) |  | The Tenant for API access |
| password | [string](#string) |  | The Password for API access |
| subscription_id | [string](#string) |  | The Subscription for API access |






<a name="cluster_manager_api.Callback"/>

### Callback



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | The URL to call back to |
| request_id | [string](#string) |  | The ID of the request |






<a name="cluster_manager_api.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |
| status | [ClusterStatus](#cluster_manager_api.ClusterStatus) |  | The status of the cluster |
| bearertoken | [string](#string) |  | What is the admin bearer token |






<a name="cluster_manager_api.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| status | [ClusterStatus](#cluster_manager_api.ClusterStatus) |  | The status of the cluster |






<a name="cluster_manager_api.CreateClusterAKSSpec"/>

### CreateClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [string](#string) |  | The Azure Data Center |
| credentials | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Credentials to build the cluster |
| instance_groups | [CreateClusterAKSSpec.AKSInstanceGroup](#cluster_manager_api.CreateClusterAKSSpec.AKSInstanceGroup) | repeated | Instance groups |






<a name="cluster_manager_api.CreateClusterAKSSpec.AKSInstanceGroup"/>

### CreateClusterAKSSpec.AKSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the group |
| type | [string](#string) |  | Instance type (Standard_D2_v2, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |






<a name="cluster_manager_api.CreateClusterAWSSpec"/>

### CreateClusterAWSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data_center | [CreateClusterAWSSpec.AWSDataCenter](#cluster_manager_api.CreateClusterAWSSpec.AWSDataCenter) |  | The AWS Data Center |
| credentials | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | Credentials to build the cluster |
| resources | [CreateClusterAWSSpec.AWSPreconfiguredItems](#cluster_manager_api.CreateClusterAWSSpec.AWSPreconfiguredItems) |  | BYO items |
| instance_groups | [CreateClusterAWSSpec.AWSInstanceGroup](#cluster_manager_api.CreateClusterAWSSpec.AWSInstanceGroup) | repeated | Instance groups |






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
| iam_role_arn | [string](#string) |  | The IAM role for the cluster (arn)ClusterAssociationdd |






<a name="cluster_manager_api.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cluster_manager_api.CreateClusterProviderSpec) |  | The provider specification |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - currently this is aws or maas |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| aws | [CreateClusterAWSSpec](#cluster_manager_api.CreateClusterAWSSpec) |  | The AWS specification |
| azure | [CreateClusterAKSSpec](#cluster_manager_api.CreateClusterAKSSpec) |  |  |
| vmware | [CreateClusterVMWareSpec](#cluster_manager_api.CreateClusterVMWareSpec) |  |  |
| ssh | [CreateClusterSshSpec](#cluster_manager_api.CreateClusterSshSpec) |  |  |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cluster_manager_api.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cluster_manager_api.ClusterItem) |  | The details of the cluster request response |






<a name="cluster_manager_api.CreateClusterSshSpec"/>

### CreateClusterSshSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| control_plane_nodes | [SshMachineSpec](#cluster_manager_api.SshMachineSpec) | repeated | Machines which comprise the cluster |
| worker_nodes | [SshMachineSpec](#cluster_manager_api.SshMachineSpec) | repeated | Machines which comprise the cluster |
| api_endpoint | [string](#string) |  | This should be a value like ip:port that will be a virtual IP/port Passed back to external customers to be able to communicate to the cluster |






<a name="cluster_manager_api.CreateClusterVMWareSpec"/>

### CreateClusterVMWareSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| control_plane_nodes | [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec) | repeated | Machines which comprise the cluster |
| worker_nodes | [VMWareMachineSpec](#cluster_manager_api.VMWareMachineSpec) | repeated | Machines which comprise the cluster |
| api_endpoint | [string](#string) |  | This should be a value like ip:port that will be a virtual IP/port Passed back to external customers to be able to communicate to the cluster |






<a name="cluster_manager_api.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






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
| package_manager | [string](#string) |  | What tiller should be used |
| chart | [string](#string) |  | Chart Name |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| request_id | [string](#string) |  | A unique id to indicate this request |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.DeleteHelmChartReply"/>

### DeleteHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.GenericHelmChart"/>

### GenericHelmChart



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the name of the deployment |
| namespace | [string](#string) |  | What is the namespace to deploy the application to |
| repo | [GenericHelmChart.ChartRepository](#cluster_manager_api.GenericHelmChart.ChartRepository) |  | What is the chart repository |
| chart | [string](#string) |  | What is the chart name |
| version | [string](#string) |  | What is the chart version |
| values | [string](#string) |  | What are the options (nested yaml - the Values.yaml contents) |
| chart_payload | [bytes](#bytes) |  | The payload of a chart (for airgapped solutions, etc) |






<a name="cluster_manager_api.GenericHelmChart.ChartRepository"/>

### GenericHelmChart.ChartRepository



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | What is the URL for the chart repo |
| name | [string](#string) |  | What is the repo name |






<a name="cluster_manager_api.GetClusterListMsg"/>

### GetClusterListMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |






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
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |






<a name="cluster_manager_api.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cluster_manager_api.ClusterDetailItem) |  |  |






<a name="cluster_manager_api.GetUpgradeClusterInformationMsg"/>

### GetUpgradeClusterInformationMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |






<a name="cluster_manager_api.GetUpgradeClusterInformationReply"/>

### GetUpgradeClusterInformationReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Can the cluster be upgraded |
| versions | [string](#string) | repeated | What versions are possible right now |






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






<a name="cluster_manager_api.InstallHelmChartMsg"/>

### InstallHelmChartMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| package_manger | [string](#string) |  | What tiller should be used |
| chart | [GenericHelmChart](#cluster_manager_api.GenericHelmChart) |  | Chart Settings |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.InstallHelmChartReply"/>

### InstallHelmChartReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.KubernetesLabel"/>

### KubernetesLabel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of a label |
| value | [string](#string) |  | The value of a label |






<a name="cluster_manager_api.ProvisionTillerMsg"/>

### ProvisionTillerMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the tiller / package manager |
| cluster | [string](#string) |  | Cluster tiller should be installed on |
| namespace | [string](#string) |  | Namespace tiller should be installed in |
| image | [string](#string) |  | image (if not the default) for tiller (quay.io/someguy/tiller) |
| version | [string](#string) |  | Version of tiller/helm to install / upgrade to (v2.11.0, etc) |
| cluster_wide | [bool](#bool) |  | Is the tiller a cluster-wide tiller? Should it have cluster-wide admin privileges? |
| admin_namespaces | [string](#string) | repeated | Namespaces that it should be able to admin on |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.ProvisionTillerReply"/>

### ProvisionTillerReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was the operation successful |
| message | [string](#string) |  | What messages were given |






<a name="cluster_manager_api.SshMachineSpec"/>

### SshMachineSpec
The specification for a specific node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | The username for SSH access |
| host | [string](#string) |  | The host for SSH access |
| port | [int32](#int32) |  | The port for SSH access |
| password | [string](#string) |  | The k8s version for the control plane. This node is only a master if this field is defined. |
| labels | [KubernetesLabel](#cluster_manager_api.KubernetesLabel) | repeated | The labels for the machines |






<a name="cluster_manager_api.UpdateAWSCredentialsMsg"/>

### UpdateAWSCredentialsMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Cluster name |
| credentials | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | Credentials to build the cluster |






<a name="cluster_manager_api.UpdateAWSCredentialsReply"/>

### UpdateAWSCredentialsReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.UpdateAzureCredentialsMsg"/>

### UpdateAzureCredentialsMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Cluster name |
| credentials | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Credentials to build the cluster |






<a name="cluster_manager_api.UpdateAzureCredentialsReply"/>

### UpdateAzureCredentialsReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.UpgradeClusterMsg"/>

### UpgradeClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [Provider](#cluster_manager_api.Provider) |  | Name of the providers (aws/azure/vmware/etc) |
| aws | [AWSCredentials](#cluster_manager_api.AWSCredentials) |  | AWS Credentials |
| azure | [AzureCredentials](#cluster_manager_api.AzureCredentials) |  | Azure Credentials |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |
| version | [string](#string) |  | What version are we upgrading to? |
| callback | [Callback](#cluster_manager_api.Callback) |  | Callback information |






<a name="cluster_manager_api.UpgradeClusterReply"/>

### UpgradeClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cluster_manager_api.VMWareMachineSpec"/>

### VMWareMachineSpec
The specification for a specific node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | The username for SSH access |
| host | [string](#string) |  | The host for SSH access |
| port | [int32](#int32) |  | The port for SSH access |
| password | [string](#string) |  | The k8s version for the control plane. This node is only a master if this field is defined. |
| labels | [KubernetesLabel](#cluster_manager_api.KubernetesLabel) | repeated | The labels for the machines |





 


<a name="cluster_manager_api.ClusterStatus"/>

### ClusterStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | Not set |
| PROVISIONING | 1 | The PROVISIONING state indicates the cluster is being created. |
| RUNNING | 2 | The RUNNING state indicates the cluster has been created and is fully usable. |
| RECONCILING | 3 | The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software. |
| STOPPING | 4 | The STOPPING state indicates the cluster is being deleted |
| ERROR | 5 | The ERROR state indicates the cluster may be unusable |
| DEGRADED | 6 | The DEGRADED state indicates the cluster requires user action to restore full functionality |



<a name="cluster_manager_api.Provider"/>

### Provider


| Name | Number | Description |
| ---- | ------ | ----------- |
| undefined | 0 |  |
| aws | 1 |  |
| azure | 2 |  |
| vmware | 3 |  |
| ssh | 4 |  |


 

 


<a name="cluster_manager_api.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cluster_manager_api.CreateClusterMsg) | [CreateClusterReply](#cluster_manager_api.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cluster_manager_api.GetClusterMsg) | [GetClusterReply](#cluster_manager_api.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cluster_manager_api.DeleteClusterMsg) | [DeleteClusterReply](#cluster_manager_api.DeleteClusterMsg) | Will delete a cluster |
| AdjustClusterNodes | [AdjustClusterMsg](#cluster_manager_api.AdjustClusterMsg) | [AdjustClusterReply](#cluster_manager_api.AdjustClusterMsg) | Will adjust a provision a cluster |
| GetClusterList | [GetClusterListMsg](#cluster_manager_api.GetClusterListMsg) | [GetClusterListReply](#cluster_manager_api.GetClusterListMsg) | Will retrieve a list of clusters |
| ProvisionTiller | [ProvisionTillerMsg](#cluster_manager_api.ProvisionTillerMsg) | [ProvisionTillerReply](#cluster_manager_api.ProvisionTillerMsg) | Will install (or reinstall) tiller |
| InstallHelmChart | [InstallHelmChartMsg](#cluster_manager_api.InstallHelmChartMsg) | [InstallHelmChartReply](#cluster_manager_api.InstallHelmChartMsg) | Will install (or reinstall) helm chart This will be destructive if a chart has already been deployed with the same name |
| DeleteHelmChart | [DeleteHelmChartMsg](#cluster_manager_api.DeleteHelmChartMsg) | [DeleteHelmChartReply](#cluster_manager_api.DeleteHelmChartMsg) | Will delete deployed helm chart |
| GetVersionInformation | [GetVersionMsg](#cluster_manager_api.GetVersionMsg) | [GetVersionReply](#cluster_manager_api.GetVersionMsg) | Will return version information about api server |
| GetUpgradeClusterInformation | [GetUpgradeClusterInformationMsg](#cluster_manager_api.GetUpgradeClusterInformationMsg) | [GetUpgradeClusterInformationReply](#cluster_manager_api.GetUpgradeClusterInformationMsg) | Will return upgrade options for a given cluster |
| UpgradeCluster | [UpgradeClusterMsg](#cluster_manager_api.UpgradeClusterMsg) | [UpgradeClusterReply](#cluster_manager_api.UpgradeClusterMsg) | Will attempt to upgrade a cluster |
| UpdateAWSCredentials | [UpdateAWSCredentialsMsg](#cluster_manager_api.UpdateAWSCredentialsMsg) | [UpdateAWSCredentialsReply](#cluster_manager_api.UpdateAWSCredentialsMsg) | Will update aws credentials used for a cluster |
| UpdateAzureCredentials | [UpdateAzureCredentialsMsg](#cluster_manager_api.UpdateAzureCredentialsMsg) | [UpdateAzureCredentialsReply](#cluster_manager_api.UpdateAzureCredentialsMsg) | Will update azure credentials used for a cluster |

 



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

