// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/internal"
)

type AddonProps struct {
	VpnCniAddonVersion VpcCniAddonVersion `json:"vpnCniAddonVersion" yaml:"vpnCniAddonVersion"`
}

type ArgoCD struct {
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	ClusterRoleName *string `json:"clusterRoleName" yaml:"clusterRoleName"`
}

type ClusterConfig struct {
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	ClusterVersion awseks.KubernetesVersion `json:"clusterVersion" yaml:"clusterVersion"`
	DefaultCapacity *float64 `json:"defaultCapacity" yaml:"defaultCapacity"`
	NodeGroups *[]*NodeGroupConfig `json:"nodeGroups" yaml:"nodeGroups"`
	Subnets *InternalMap `json:"subnets" yaml:"subnets"`
	Tags *InternalMap `json:"tags" yaml:"tags"`
	TeamMembers *[]*string `json:"teamMembers" yaml:"teamMembers"`
	AddAutoscalerIam *bool `json:"addAutoscalerIam" yaml:"addAutoscalerIam"`
	AlbControllerVersion awseks.AlbControllerVersion `json:"albControllerVersion" yaml:"albControllerVersion"`
	ArgoCD *ArgoCD `json:"argoCD" yaml:"argoCD"`
	CommonComponents *map[string]ICommonComponentsProps `json:"commonComponents" yaml:"commonComponents"`
	FargateProfiles *[]*FargateProfile `json:"fargateProfiles" yaml:"fargateProfiles"`
	Namespaces *map[string]*NamespaceSpec `json:"namespaces" yaml:"namespaces"`
	PublicAllowAccess *[]*string `json:"publicAllowAccess" yaml:"publicAllowAccess"`
	TeamExistingRolePermission *map[string]*string `json:"teamExistingRolePermission" yaml:"teamExistingRolePermission"`
}

type CommonHelmCharts interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CommonHelmCharts
type jsiiProxy_CommonHelmCharts struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CommonHelmCharts) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCommonHelmCharts(scope constructs.Construct, id *string, props *CommonHelmChartsProps) CommonHelmCharts {
	_init_.Initialize()

	j := jsiiProxy_CommonHelmCharts{}

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.CommonHelmCharts",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCommonHelmCharts_Override(c CommonHelmCharts, scope constructs.Construct, id *string, props *CommonHelmChartsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.CommonHelmCharts",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CommonHelmCharts_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.CommonHelmCharts",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CommonHelmCharts) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CommonHelmChartsProps struct {
	Cluster awseks.ICluster `json:"cluster" yaml:"cluster"`
	HelmProps *StandardHelmProps `json:"helmProps" yaml:"helmProps"`
	IamPolicyPath *[]*string `json:"iamPolicyPath" yaml:"iamPolicyPath"`
	ServiceAccounts *[]*string `json:"serviceAccounts" yaml:"serviceAccounts"`
}

type EKSCluster interface {
	constructs.Construct
	AdditionalFargateProfile() *[]awseks.FargateProfile
	AdditionalNodegroups() *[]awseks.Nodegroup
	Cluster() awseks.Cluster
	FargateProfiles() *[]*FargateProfile
	// The tree node.
	Node() constructs.Node
	AddServiceAccountWithIamRole(serviceAccountName *string, serviceAccountNamespace *string, policy interface{})
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EKSCluster
type jsiiProxy_EKSCluster struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EKSCluster) AdditionalFargateProfile() *[]awseks.FargateProfile {
	var returns *[]awseks.FargateProfile
	_jsii_.Get(
		j,
		"additionalFargateProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EKSCluster) AdditionalNodegroups() *[]awseks.Nodegroup {
	var returns *[]awseks.Nodegroup
	_jsii_.Get(
		j,
		"additionalNodegroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EKSCluster) Cluster() awseks.Cluster {
	var returns awseks.Cluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EKSCluster) FargateProfiles() *[]*FargateProfile {
	var returns *[]*FargateProfile
	_jsii_.Get(
		j,
		"fargateProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EKSCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEKSCluster(scope constructs.Construct, id *string, props *EKSClusterProps) EKSCluster {
	_init_.Initialize()

	j := jsiiProxy_EKSCluster{}

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.EKSCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEKSCluster_Override(e EKSCluster, scope constructs.Construct, id *string, props *EKSClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.EKSCluster",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func EKSCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.EKSCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EKSCluster) AddServiceAccountWithIamRole(serviceAccountName *string, serviceAccountNamespace *string, policy interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addServiceAccountWithIamRole",
		[]interface{}{serviceAccountName, serviceAccountNamespace, policy},
	)
}

func (e *jsiiProxy_EKSCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EKSClusterProps struct {
	AvailabilityZones *[]*string `json:"availabilityZones" yaml:"availabilityZones"`
	ClusterConfig *ClusterConfig `json:"clusterConfig" yaml:"clusterConfig"`
	KmsKey awskms.Key `json:"kmsKey" yaml:"kmsKey"`
	Region *string `json:"region" yaml:"region"`
	WorkerSecurityGroup awsec2.SecurityGroup `json:"workerSecurityGroup" yaml:"workerSecurityGroup"`
	AddonProps *AddonProps `json:"addonProps" yaml:"addonProps"`
	ClusterVPC awsec2.IVpc `json:"clusterVPC" yaml:"clusterVPC"`
}

type FargateProfile struct {
	Namespaces *[]*string `json:"namespaces" yaml:"namespaces"`
	ProfileName *string `json:"profileName" yaml:"profileName"`
	Labels *InternalMap `json:"labels" yaml:"labels"`
}

type ICommonComponentsProps interface {
	Helm() *StandardHelmProps
	SetHelm(h *StandardHelmProps)
	IamPolicyPath() *[]*string
	SetIamPolicyPath(i *[]*string)
	ServiceAccounts() *[]*string
	SetServiceAccounts(s *[]*string)
}

// The jsii proxy for ICommonComponentsProps
type jsiiProxy_ICommonComponentsProps struct {
	_ byte // padding
}

func (j *jsiiProxy_ICommonComponentsProps) Helm() *StandardHelmProps {
	var returns *StandardHelmProps
	_jsii_.Get(
		j,
		"helm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICommonComponentsProps) SetHelm(val *StandardHelmProps) {
	_jsii_.Set(
		j,
		"helm",
		val,
	)
}

func (j *jsiiProxy_ICommonComponentsProps) IamPolicyPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamPolicyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICommonComponentsProps) SetIamPolicyPath(val *[]*string) {
	_jsii_.Set(
		j,
		"iamPolicyPath",
		val,
	)
}

func (j *jsiiProxy_ICommonComponentsProps) ServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICommonComponentsProps) SetServiceAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"serviceAccounts",
		val,
	)
}

type InternalMap struct {
}

type NamespaceSpec struct {
	Annotations *InternalMap `json:"annotations" yaml:"annotations"`
	Labels *InternalMap `json:"labels" yaml:"labels"`
}

type NodeGroupConfig struct {
	InstanceTypes *[]awsec2.InstanceType `json:"instanceTypes" yaml:"instanceTypes"`
	Labels *InternalMap `json:"labels" yaml:"labels"`
	MaxSize *float64 `json:"maxSize" yaml:"maxSize"`
	MinSize *float64 `json:"minSize" yaml:"minSize"`
	Name *string `json:"name" yaml:"name"`
	SubnetGroupName *string `json:"subnetGroupName" yaml:"subnetGroupName"`
	Taints *InternalMap `json:"taints" yaml:"taints"`
	CapacityType awseks.CapacityType `json:"capacityType" yaml:"capacityType"`
	DesiredSize *float64 `json:"desiredSize" yaml:"desiredSize"`
	DiskSize *float64 `json:"diskSize" yaml:"diskSize"`
	LaunchTemplateSpec *awseks.LaunchTemplateSpec `json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	NodeAMIVersion *string `json:"nodeAMIVersion" yaml:"nodeAMIVersion"`
	SshKeyName *string `json:"sshKeyName" yaml:"sshKeyName"`
	SubnetAz *[]*string `json:"subnetAz" yaml:"subnetAz"`
	Tags *InternalMap `json:"tags" yaml:"tags"`
}

type StandardHelmProps struct {
	ChartName *string `json:"chartName" yaml:"chartName"`
	ChartReleaseName *string `json:"chartReleaseName" yaml:"chartReleaseName"`
	ChartVersion *string `json:"chartVersion" yaml:"chartVersion"`
	CreateNamespace *bool `json:"createNamespace" yaml:"createNamespace"`
	HelmRepository *string `json:"helmRepository" yaml:"helmRepository"`
	HelmValues *map[string]interface{} `json:"helmValues" yaml:"helmValues"`
	LocalHelmChart *string `json:"localHelmChart" yaml:"localHelmChart"`
	Namespace *string `json:"namespace" yaml:"namespace"`
}

type VpcCniAddonProps struct {
	Cluster awseks.Cluster `json:"cluster" yaml:"cluster"`
	AddonVersion VpcCniAddonVersion `json:"addonVersion" yaml:"addonVersion"`
	Namespace *string `json:"namespace" yaml:"namespace"`
	ResolveConflicts *bool `json:"resolveConflicts" yaml:"resolveConflicts"`
}

type VpcCniAddonVersion interface {
	Version() *string
}

// The jsii proxy struct for VpcCniAddonVersion
type jsiiProxy_VpcCniAddonVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_VpcCniAddonVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewVpcCniAddonVersion(version *string) VpcCniAddonVersion {
	_init_.Initialize()

	j := jsiiProxy_VpcCniAddonVersion{}

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		[]interface{}{version},
		&j,
	)

	return &j
}

func NewVpcCniAddonVersion_Override(v VpcCniAddonVersion, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		[]interface{}{version},
		v,
	)
}

// Custom add-on version.
func VpcCniAddonVersion_Of(version *string) VpcCniAddonVersion {
	_init_.Initialize()

	var returns VpcCniAddonVersion

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func VpcCniAddonVersion_V1_10_1() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_10_1",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_10_2() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_10_2",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_10_3() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_10_3",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_11_0() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_11_0",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_6_3() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_6_3",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_7_10() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_7_10",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_7_5() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_7_5",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_7_6() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_7_6",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_7_9() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_7_9",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_8_0() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_8_0",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_9_0() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_9_0",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_9_1() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_9_1",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_9_3() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_9_3",
		&returns,
	)
	return returns
}

type VpcEniAddon interface {
	awseks.CfnAddon
	// The name of the add-on.
	AddonName() *string
	SetAddonName(val *string)
	// The version of the add-on.
	AddonVersion() *string
	SetAddonVersion(val *string)
	// The ARN of the add-on, such as `arn:aws:eks:us-west-2:111122223333:addon/1-19/vpc-cni/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The name of the cluster.
	ClusterName() *string
	SetClusterName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on.
	ResolveConflicts() *string
	SetResolveConflicts(val *string)
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.
	//
	// The role must be assigned the IAM permissions required by the add-on. If you don't specify an existing IAM role, then the add-on uses the permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the *Amazon EKS User Guide* .
	//
	// > To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC) provider created for your cluster. For more information, see [Enabling IAM roles for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html) in the *Amazon EKS User Guide* .
	ServiceAccountRoleArn() *string
	SetServiceAccountRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The metadata that you apply to the add-on to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Add-on tags do not propagate to any other resources associated with the cluster.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for VpcEniAddon
type jsiiProxy_VpcEniAddon struct {
	internal.Type__awseksCfnAddon
}

func (j *jsiiProxy_VpcEniAddon) AddonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) AddonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) ResolveConflicts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveConflicts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) ServiceAccountRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEniAddon) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


func NewVpcEniAddon(scope constructs.Construct, id *string, props *VpcCniAddonProps) VpcEniAddon {
	_init_.Initialize()

	j := jsiiProxy_VpcEniAddon{}

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.VpcEniAddon",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVpcEniAddon_Override(v VpcEniAddon, scope constructs.Construct, id *string, props *VpcCniAddonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.VpcEniAddon",
		[]interface{}{scope, id, props},
		v,
	)
}

func (j *jsiiProxy_VpcEniAddon) SetAddonName(val *string) {
	_jsii_.Set(
		j,
		"addonName",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon) SetAddonVersion(val *string) {
	_jsii_.Set(
		j,
		"addonVersion",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon) SetResolveConflicts(val *string) {
	_jsii_.Set(
		j,
		"resolveConflicts",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon) SetServiceAccountRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountRoleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func VpcEniAddon_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.VpcEniAddon",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func VpcEniAddon_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.VpcEniAddon",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func VpcEniAddon_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.VpcEniAddon",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcEniAddon_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcEniAddon",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcEniAddon) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		v,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		v,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		v,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (v *jsiiProxy_VpcEniAddon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (v *jsiiProxy_VpcEniAddon) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		v,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		v,
		"inspect",
		[]interface{}{inspector},
	)
}

func (v *jsiiProxy_VpcEniAddon) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcEniAddon) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		v,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		v,
		"validateProperties",
		[]interface{}{_properties},
	)
}

