// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/internal"
)

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
	// The configuration values that you provided.
	ConfigurationValues() *string
	SetConfigurationValues(val *string)
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
	// `AWS::EKS::Addon.PreserveOnDelete`.
	PreserveOnDelete() interface{}
	SetPreserveOnDelete(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// How to resolve field value conflicts for an Amazon EKS add-on.
	//
	// Conflicts are handled based on the value you choose:
	//
	// - *None* – If the self-managed version of the add-on is installed on your cluster, Amazon EKS doesn't change the value. Creation of the add-on might fail.
	// - *Overwrite* – If the self-managed version of the add-on is installed on your cluster and the Amazon EKS default value is different than the existing value, Amazon EKS changes the value to the Amazon EKS default value.
	// - *Preserve* – Not supported. You can set this value when updating an add-on though. For more information, see [UpdateAddon](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html) .
	//
	// If you don't currently have the self-managed version of the add-on installed on your cluster, the Amazon EKS add-on is installed. Amazon EKS sets all values to default values, regardless of the option that you specify.
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
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

func (j *jsiiProxy_VpcEniAddon) ConfigurationValues() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationValues",
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

func (j *jsiiProxy_VpcEniAddon) PreserveOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveOnDelete",
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

func (j *jsiiProxy_VpcEniAddon) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewVpcEniAddon(scope constructs.Construct, id *string, props *VpcCniAddonProps) VpcEniAddon {
	_init_.Initialize()

	if err := validateNewVpcEniAddonParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_VpcEniAddon)SetAddonName(val *string) {
	if err := j.validateSetAddonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addonName",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon)SetAddonVersion(val *string) {
	_jsii_.Set(
		j,
		"addonVersion",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon)SetConfigurationValues(val *string) {
	_jsii_.Set(
		j,
		"configurationValues",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon)SetPreserveOnDelete(val interface{}) {
	if err := j.validateSetPreserveOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveOnDelete",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon)SetResolveConflicts(val *string) {
	_jsii_.Set(
		j,
		"resolveConflicts",
		val,
	)
}

func (j *jsiiProxy_VpcEniAddon)SetServiceAccountRoleArn(val *string) {
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

	if err := validateVpcEniAddon_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateVpcEniAddon_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateVpcEniAddon_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := v.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddDependency(target awscdk.CfnResource) {
	if err := v.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addDependency",
		[]interface{}{target},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddDependsOn(target awscdk.CfnResource) {
	if err := v.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddMetadata(key *string, value interface{}) {
	if err := v.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddPropertyDeletionOverride(propertyPath *string) {
	if err := v.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (v *jsiiProxy_VpcEniAddon) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := v.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (v *jsiiProxy_VpcEniAddon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := v.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (v *jsiiProxy_VpcEniAddon) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := v.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		v,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) GetMetadata(key *string) interface{} {
	if err := v.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
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
	if err := v.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"inspect",
		[]interface{}{inspector},
	)
}

func (v *jsiiProxy_VpcEniAddon) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		v,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		v,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcEniAddon) RemoveDependency(target awscdk.CfnResource) {
	if err := v.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"removeDependency",
		[]interface{}{target},
	)
}

func (v *jsiiProxy_VpcEniAddon) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := v.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEniAddon) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := v.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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
	if err := v.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"validateProperties",
		[]interface{}{_properties},
	)
}

