package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AddOnManagerConfiguration = map[string]string{
	"featureGates": "FeatureGates represents the list of feature gates for addon manager If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:\n  1. If featuregate/Foo does not exist, registration-operator will discard it\n  2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]\n  3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,\n \the can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
}

func (AddOnManagerConfiguration) SwaggerDoc() map[string]string {
	return map_AddOnManagerConfiguration
}

var map_AwsIrsaConfig = map[string]string{
	"hubClusterArn":          "This represents the hub cluster ARN Example - arn:eks:us-west-2:12345678910:cluster/hub-cluster1",
	"autoApprovedIdentities": "AutoApprovedIdentities represent a list of approved arn patterns",
	"tags":                   "List of tags to be added to AWS resources created by hub while processing awsirsa registration request Example - \"product:v1:tenant:app-name=My-App\"",
}

func (AwsIrsaConfig) SwaggerDoc() map[string]string {
	return map_AwsIrsaConfig
}

var map_CSRConfig = map[string]string{
	"autoApprovedIdentities": "AutoApprovedIdentities represent a list of approved users",
}

func (CSRConfig) SwaggerDoc() map[string]string {
	return map_CSRConfig
}

var map_ClusterManager = map[string]string{
	"":       "ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets. In Default mode, ClusterManager will only be deployed in open-cluster-management-hub namespace. In Hosted mode, ClusterManager will be deployed in the namespace with the same name as cluster manager.",
	"spec":   "Spec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.",
	"status": "Status represents the current status of controllers that govern the lifecycle of managed clusters.",
}

func (ClusterManager) SwaggerDoc() map[string]string {
	return map_ClusterManager
}

var map_ClusterManagerDeployOption = map[string]string{
	"":       "ClusterManagerDeployOption describes the deployment options for cluster-manager",
	"mode":   "Mode can be Default or Hosted. In Default mode, the Hub is installed as a whole and all parts of Hub are deployed in the same cluster. In Hosted mode, only crd and configurations are installed on one cluster(defined as hub-cluster). Controllers run in another cluster (defined as management-cluster) and connect to the hub with the kubeconfig in secret of \"external-hub-kubeconfig\"(a kubeconfig of hub-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
	"hosted": "Hosted includes configurations we need for clustermanager in the Hosted mode.",
}

func (ClusterManagerDeployOption) SwaggerDoc() map[string]string {
	return map_ClusterManagerDeployOption
}

var map_ClusterManagerList = map[string]string{
	"":         "ClusterManagerList is a collection of deployment configurations for registration and work distribution controllers.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of deployment configurations for registration and work distribution controllers.",
}

func (ClusterManagerList) SwaggerDoc() map[string]string {
	return map_ClusterManagerList
}

var map_ClusterManagerSpec = map[string]string{
	"":                          "ClusterManagerSpec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.",
	"registrationImagePullSpec": "RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.",
	"workImagePullSpec":         "WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.",
	"placementImagePullSpec":    "PlacementImagePullSpec represents the desired image configuration of placement controller/webhook installed on hub.",
	"addOnManagerImagePullSpec": "AddOnManagerImagePullSpec represents the desired image configuration of addon manager controller/webhook installed on hub.",
	"nodePlacement":             "NodePlacement enables explicit control over the scheduling of the deployed pods.",
	"deployOption":              "DeployOption contains the options of deploying a cluster-manager Default mode is used if DeployOption is not set.",
	"registrationConfiguration": "RegistrationConfiguration contains the configuration of registration",
	"workConfiguration":         "WorkConfiguration contains the configuration of work",
	"addOnManagerConfiguration": "AddOnManagerConfiguration contains the configuration of addon manager",
	"resourceRequirement":       "ResourceRequirement specify QoS classes of deployments managed by clustermanager. It applies to all the containers in the deployments.",
}

func (ClusterManagerSpec) SwaggerDoc() map[string]string {
	return map_ClusterManagerSpec
}

var map_ClusterManagerStatus = map[string]string{
	"":                   "ClusterManagerStatus represents the current status of the registration and work distribution controllers running on the hub.",
	"observedGeneration": "ObservedGeneration is the last generation change you've dealt with",
	"conditions":         "Conditions contain the different condition statuses for this ClusterManager. Valid condition types are: Applied: Components in hub are applied. Available: Components in hub are available and ready to serve. Progressing: Components in hub are in a transitioning state. Degraded: Components in hub do not match the desired configuration and only provide degraded service.",
	"generations":        "Generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.",
	"relatedResources":   "RelatedResources are used to track the resources that are related to this ClusterManager.",
}

func (ClusterManagerStatus) SwaggerDoc() map[string]string {
	return map_ClusterManagerStatus
}

var map_FeatureGate = map[string]string{
	"feature": "Feature is the key of feature gate. e.g. featuregate/Foo.",
	"mode":    "Mode is either Enable, Disable, \"\" where \"\" is Disable by default. In Enable mode, a valid feature gate `featuregate/Foo` will be set to \"--featuregate/Foo=true\". In Disable mode, a valid feature gate `featuregate/Foo` will be set to \"--featuregate/Foo=false\".",
}

func (FeatureGate) SwaggerDoc() map[string]string {
	return map_FeatureGate
}

var map_GenerationStatus = map[string]string{
	"":               "GenerationStatus keeps track of the generation for a given resource so that decisions about forced updates can be made. The definition matches the GenerationStatus defined in github.com/openshift/api/v1",
	"group":          "group is the group of the resource that you're tracking",
	"version":        "version is the version of the resource that you're tracking",
	"resource":       "resource is the resource type of the resource that you're tracking",
	"namespace":      "namespace is where the resource that you're tracking is",
	"name":           "name is the name of the resource that you're tracking",
	"lastGeneration": "lastGeneration is the last generation of the resource that controller applies",
}

func (GenerationStatus) SwaggerDoc() map[string]string {
	return map_GenerationStatus
}

var map_HostedClusterManagerConfiguration = map[string]string{
	"":                                 "HostedClusterManagerConfiguration represents customized configurations we need to set for clustermanager in the Hosted mode.",
	"registrationWebhookConfiguration": "RegistrationWebhookConfiguration represents the customized webhook-server configuration of registration.",
	"workWebhookConfiguration":         "WorkWebhookConfiguration represents the customized webhook-server configuration of work.",
}

func (HostedClusterManagerConfiguration) SwaggerDoc() map[string]string {
	return map_HostedClusterManagerConfiguration
}

var map_NodePlacement = map[string]string{
	"":             "NodePlacement describes node scheduling configuration for the pods.",
	"nodeSelector": "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
	"tolerations":  "Tolerations are attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
}

func (NodePlacement) SwaggerDoc() map[string]string {
	return map_NodePlacement
}

var map_RegistrationDriverHub = map[string]string{
	"authType": "Type of the authentication used by hub to initialize the Hub cluster. Possible values are csr and awsirsa.",
	"csr":      "CSR represents the configuration for csr driver.",
	"awsisra":  "AwsIrsa represents the configuration for awsisra driver.",
}

func (RegistrationDriverHub) SwaggerDoc() map[string]string {
	return map_RegistrationDriverHub
}

var map_RegistrationHubConfiguration = map[string]string{
	"autoApproveUsers":    "AutoApproveUser represents a list of users that can auto approve CSR and accept client. If the credential of the bootstrap-hub-kubeconfig matches to the users, the cluster created by the bootstrap-hub-kubeconfig will be auto-registered into the hub cluster. This takes effect only when ManagedClusterAutoApproval feature gate is enabled.",
	"featureGates":        "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:\n  1. If featuregate/Foo does not exist, registration-operator will discard it\n  2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]\n  3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,\n \the can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
	"registrationDrivers": "RegistrationDrivers represent the list of hub registration drivers that contain information used by hub to initialize the hub cluster A RegistrationDriverHub contains details of authentication type and the hub cluster ARN",
}

func (RegistrationHubConfiguration) SwaggerDoc() map[string]string {
	return map_RegistrationHubConfiguration
}

var map_RelatedResourceMeta = map[string]string{
	"":          "RelatedResourceMeta represents the resource that is managed by an operator",
	"group":     "group is the group of the resource that you're tracking",
	"version":   "version is the version of the thing you're tracking",
	"resource":  "resource is the resource type of the resource that you're tracking",
	"namespace": "namespace is where the thing you're tracking is",
	"name":      "name is the name of the resource that you're tracking",
}

func (RelatedResourceMeta) SwaggerDoc() map[string]string {
	return map_RelatedResourceMeta
}

var map_WebhookConfiguration = map[string]string{
	"":        "WebhookConfiguration has two properties: Address and Port.",
	"address": "Address represents the address of a webhook-server. It could be in IP format or fqdn format. The Address must be reachable by apiserver of the hub cluster.",
	"port":    "Port represents the port of a webhook-server. The default value of Port is 443.",
}

func (WebhookConfiguration) SwaggerDoc() map[string]string {
	return map_WebhookConfiguration
}

var map_WorkConfiguration = map[string]string{
	"featureGates": "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:\n  1. If featuregate/Foo does not exist, registration-operator will discard it\n  2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]\n  3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,\n \the can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
	"workDriver":   "WorkDriver represents the type of work driver. Possible values are \"kube\", \"mqtt\", or \"grpc\". If not provided, the default value is \"kube\". If set to non-\"kube\" drivers, the klusterlet need to use the same driver. and the driver configuration must be provided in a secret named \"work-driver-config\" in the namespace where the cluster manager is running, adhering to the following structure: config.yaml: |\n  <driver-config-in-yaml>\n\nFor detailed driver configuration, please refer to the sdk-go documentation: https://github.com/open-cluster-management-io/sdk-go/blob/main/pkg/cloudevents/README.md#supported-protocols-and-drivers",
}

func (WorkConfiguration) SwaggerDoc() map[string]string {
	return map_WorkConfiguration
}

var map_AwsIrsa = map[string]string{
	"hubClusterArn":     "The arn of the hub cluster (ie: an EKS cluster). This will be required to pass information to hub, which hub will use to create IAM identities for this klusterlet. Example - arn:eks:us-west-2:12345678910:cluster/hub-cluster1.",
	"managedClusterArn": "The arn of the managed cluster (ie: an EKS cluster). This will be required to generate the md5hash which will be used as a suffix to create IAM role on hub as well as used by kluslerlet-agent, to assume role suffixed with the md5hash, on startup. Example - arn:eks:us-west-2:12345678910:cluster/managed-cluster1.",
}

func (AwsIrsa) SwaggerDoc() map[string]string {
	return map_AwsIrsa
}

var map_BootstrapKubeConfigs = map[string]string{
	"type":               "Type specifies the type of priority bootstrap kubeconfigs. By default, it is set to None, representing no priority bootstrap kubeconfigs are set.",
	"localSecretsConfig": "LocalSecretsConfig include a list of secrets that contains the kubeconfigs for ordered bootstrap kubeconifigs. The secrets must be in the same namespace where the agent controller runs.",
}

func (BootstrapKubeConfigs) SwaggerDoc() map[string]string {
	return map_BootstrapKubeConfigs
}

var map_HubApiServerHostAlias = map[string]string{
	"":         "HubApiServerHostAlias holds the mapping between IP and hostname that will be injected as an entry in the pod's hosts file.",
	"ip":       "IP address of the host file entry.",
	"hostname": "Hostname for the above IP address.",
}

func (HubApiServerHostAlias) SwaggerDoc() map[string]string {
	return map_HubApiServerHostAlias
}

var map_Klusterlet = map[string]string{
	"":       "Klusterlet represents controllers to install the resources for a managed cluster. When configured, the Klusterlet requires a secret named bootstrap-hub-kubeconfig in the agent namespace to allow API requests to the hub for the registration protocol. In Hosted mode, the Klusterlet requires an additional secret named external-managed-kubeconfig in the agent namespace to allow API requests to the managed cluster for resources installation.",
	"spec":   "Spec represents the desired deployment configuration of Klusterlet agent.",
	"status": "Status represents the current status of Klusterlet agent.",
}

func (Klusterlet) SwaggerDoc() map[string]string {
	return map_Klusterlet
}

var map_KlusterletDeployOption = map[string]string{
	"":     "KlusterletDeployOption describes the deployment options for klusterlet",
	"mode": "Mode can be Default, Hosted, Singleton or SingletonHosted. It is Default mode if not specified In Default mode, all klusterlet related resources are deployed on the managed cluster. In Hosted mode, only crd and configurations are installed on the spoke/managed cluster. Controllers run in another cluster (defined as management-cluster) and connect to the mangaged cluster with the kubeconfig in secret of \"external-managed-kubeconfig\"(a kubeconfig of managed-cluster with cluster-admin permission). In Singleton mode, registration/work agent is started as a single deployment. In SingletonHosted mode, agent is started as a single deployment in hosted mode. Note: Do not modify the Mode field once it's applied.",
}

func (KlusterletDeployOption) SwaggerDoc() map[string]string {
	return map_KlusterletDeployOption
}

var map_KlusterletList = map[string]string{
	"":         "KlusterletList is a collection of Klusterlet agents.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of Klusterlet agents.",
}

func (KlusterletList) SwaggerDoc() map[string]string {
	return map_KlusterletList
}

var map_KlusterletSpec = map[string]string{
	"":                          "KlusterletSpec represents the desired deployment configuration of Klusterlet agent.",
	"namespace":                 "Namespace is the namespace to deploy the agent on the managed cluster. The namespace must have a prefix of \"open-cluster-management-\", and if it is not set, the namespace of \"open-cluster-management-agent\" is used to deploy agent. In addition, the add-ons are deployed to the namespace of \"{Namespace}-addon\". In the Hosted mode, this namespace still exists on the managed cluster to contain necessary resources, like service accounts, roles and rolebindings, while the agent is deployed to the namespace with the same name as klusterlet on the management cluster.",
	"registrationImagePullSpec": "RegistrationImagePullSpec represents the desired image configuration of registration agent. quay.io/open-cluster-management.io/registration:latest will be used if unspecified.",
	"workImagePullSpec":         "WorkImagePullSpec represents the desired image configuration of work agent. quay.io/open-cluster-management.io/work:latest will be used if unspecified.",
	"imagePullSpec":             "ImagePullSpec represents the desired image configuration of agent, it takes effect only when singleton mode is set. quay.io/open-cluster-management.io/registration-operator:latest will be used if unspecified",
	"clusterName":               "ClusterName is the name of the managed cluster to be created on hub. The Klusterlet agent generates a random name if it is not set, or discovers the appropriate cluster name on OpenShift.",
	"externalServerURLs":        "ExternalServerURLs represents a list of apiserver urls and ca bundles that is accessible externally If it is set empty, managed cluster has no externally accessible url that hub cluster can visit.",
	"nodePlacement":             "NodePlacement enables explicit control over the scheduling of the deployed pods.",
	"deployOption":              "DeployOption contains the options of deploying a klusterlet",
	"registrationConfiguration": "RegistrationConfiguration contains the configuration of registration",
	"workConfiguration":         "WorkConfiguration contains the configuration of work",
	"hubApiServerHostAlias":     "HubApiServerHostAlias contains the host alias for hub api server. registration-agent and work-agent will use it to communicate with hub api server.",
	"resourceRequirement":       "ResourceRequirement specify QoS classes of deployments managed by klusterlet. It applies to all the containers in the deployments.",
	"priorityClassName":         "PriorityClassName is the name of the PriorityClass that will be used by the deployed klusterlet agent. It will be ignored when the PriorityClass/v1 API is not available on the managed cluster.",
}

func (KlusterletSpec) SwaggerDoc() map[string]string {
	return map_KlusterletSpec
}

var map_KlusterletStatus = map[string]string{
	"":                   "KlusterletStatus represents the current status of Klusterlet agent.",
	"observedGeneration": "ObservedGeneration is the last generation change you've dealt with",
	"conditions":         "Conditions contain the different condition statuses for this Klusterlet. Valid condition types are: Applied: Components have been applied in the managed cluster. Available: Components in the managed cluster are available and ready to serve. Progressing: Components in the managed cluster are in a transitioning state. Degraded: Components in the managed cluster do not match the desired configuration and only provide degraded service.",
	"generations":        "Generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.",
	"relatedResources":   "RelatedResources are used to track the resources that are related to this Klusterlet.",
}

func (KlusterletStatus) SwaggerDoc() map[string]string {
	return map_KlusterletStatus
}

var map_KubeConfigSecret = map[string]string{
	"name": "Name is the name of the secret.",
}

func (KubeConfigSecret) SwaggerDoc() map[string]string {
	return map_KubeConfigSecret
}

var map_LocalSecretsConfig = map[string]string{
	"kubeConfigSecrets":           "KubeConfigSecrets is a list of secret names. The secrets are in the same namespace where the agent controller runs.",
	"hubConnectionTimeoutSeconds": "HubConnectionTimeoutSeconds is used to set the timeout of connecting to the hub cluster. When agent loses the connection to the hub over the timeout seconds, the agent do a rebootstrap. By default is 10 mins.",
}

func (LocalSecretsConfig) SwaggerDoc() map[string]string {
	return map_LocalSecretsConfig
}

var map_RegistrationConfiguration = map[string]string{
	"clientCertExpirationSeconds": "clientCertExpirationSeconds represents the seconds of a client certificate to expire. If it is not set or 0, the default duration seconds will be set by the hub cluster. If the value is larger than the max signing duration seconds set on the hub cluster, the max signing duration seconds will be set.",
	"featureGates":                "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:\n  1. If featuregate/Foo does not exist, registration-operator will discard it\n  2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]\n  3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,\n \the can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
	"clusterAnnotations":          "ClusterAnnotations is annotations with the reserve prefix \"agent.open-cluster-management.io\" set on ManagedCluster when creating only, other actors can update it afterwards.",
	"kubeAPIQPS":                  "KubeAPIQPS indicates the maximum QPS while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 50",
	"kubeAPIBurst":                "KubeAPIBurst indicates the maximum burst of the throttle while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 100",
	"bootstrapKubeConfigs":        "BootstrapKubeConfigs defines the ordered list of bootstrap kubeconfigs. The order decides which bootstrap kubeconfig to use first when rebootstrap.\n\nWhen the agent loses the connection to the current hub over HubConnectionTimeoutSeconds, or the managedcluster CR is set `hubAcceptsClient=false` on the hub, the controller marks the related bootstrap kubeconfig as \"failed\".\n\nA failed bootstrapkubeconfig won't be used for the duration specified by SkipFailedBootstrapKubeConfigSeconds. But if the user updates the content of a failed bootstrapkubeconfig, the \"failed\" mark will be cleared.",
	"registrationDriver":          "This provides driver details required to register with hub",
}

func (RegistrationConfiguration) SwaggerDoc() map[string]string {
	return map_RegistrationConfiguration
}

var map_RegistrationDriver = map[string]string{
	"authType": "Type of the authentication used by managedcluster to register as well as pull work from hub. Possible values are csr and awsirsa.",
	"awsIrsa":  "Contain the details required for registering with hub cluster (ie: an EKS cluster) using AWS IAM roles for service account. This is required only when the authType is awsirsa.",
}

func (RegistrationDriver) SwaggerDoc() map[string]string {
	return map_RegistrationDriver
}

var map_ServerURL = map[string]string{
	"":         "ServerURL represents the apiserver url and ca bundle that is accessible externally",
	"url":      "URL is the url of apiserver endpoint of the managed cluster.",
	"caBundle": "CABundle is the ca bundle to connect to apiserver of the managed cluster. System certs are used if it is not set.",
}

func (ServerURL) SwaggerDoc() map[string]string {
	return map_ServerURL
}

var map_WorkAgentConfiguration = map[string]string{
	"featureGates":                           "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates:\n  1. If featuregate/Foo does not exist, registration-operator will discard it\n  2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true]\n  3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false,\n \the can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
	"kubeAPIQPS":                             "KubeAPIQPS indicates the maximum QPS while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 50",
	"kubeAPIBurst":                           "KubeAPIBurst indicates the maximum burst of the throttle while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 100",
	"appliedManifestWorkEvictionGracePeriod": "AppliedManifestWorkEvictionGracePeriod is the eviction grace period the work agent will wait before evicting the AppliedManifestWorks, whose corresponding ManifestWorks are missing on the hub cluster, from the managed cluster. If not present, the default value of the work agent will be used.",
}

func (WorkAgentConfiguration) SwaggerDoc() map[string]string {
	return map_WorkAgentConfiguration
}

// AUTO-GENERATED FUNCTIONS END HERE
