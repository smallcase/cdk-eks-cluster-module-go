package smallcasecdkeksclustermodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/internal"
)

type AwsEfsCsiAddon interface {
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
	// The name of your cluster.
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
	// An array of Pod Identity Assocations owned by the Addon.
	PodIdentityAssociations() interface{}
	SetPodIdentityAssociations(val interface{})
	// Specifying this option preserves the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.
	PreserveOnDelete() interface{}
	SetPreserveOnDelete(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// How to resolve field value conflicts for an Amazon EKS add-on.
	ResolveConflicts() *string
	SetResolveConflicts(val *string)
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.
	ServiceAccountRoleArn() *string
	SetServiceAccountRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The metadata that you apply to the add-on to assist with categorization and organization.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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

// The jsii proxy struct for AwsEfsCsiAddon
type jsiiProxy_AwsEfsCsiAddon struct {
	internal.Type__awseksCfnAddon
}

func (j *jsiiProxy_AwsEfsCsiAddon) AddonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) AddonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) ConfigurationValues() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) PodIdentityAssociations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"podIdentityAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) PreserveOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) ResolveConflicts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveConflicts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) ServiceAccountRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsCsiAddon) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewAwsEfsCsiAddon(scope constructs.Construct, id *string, props *CoreAddonProps) AwsEfsCsiAddon {
	_init_.Initialize()

	if err := validateNewAwsEfsCsiAddonParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEfsCsiAddon{}

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.AwsEfsCsiAddon",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAwsEfsCsiAddon_Override(a AwsEfsCsiAddon, scope constructs.Construct, id *string, props *CoreAddonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.AwsEfsCsiAddon",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetAddonName(val *string) {
	if err := j.validateSetAddonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addonName",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetAddonVersion(val *string) {
	_jsii_.Set(
		j,
		"addonVersion",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetConfigurationValues(val *string) {
	_jsii_.Set(
		j,
		"configurationValues",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetPodIdentityAssociations(val interface{}) {
	if err := j.validateSetPodIdentityAssociationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podIdentityAssociations",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetPreserveOnDelete(val interface{}) {
	if err := j.validateSetPreserveOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveOnDelete",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetResolveConflicts(val *string) {
	_jsii_.Set(
		j,
		"resolveConflicts",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetServiceAccountRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEfsCsiAddon)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func AwsEfsCsiAddon_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEfsCsiAddon_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.AwsEfsCsiAddon",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func AwsEfsCsiAddon_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEfsCsiAddon_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.AwsEfsCsiAddon",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AwsEfsCsiAddon_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEfsCsiAddon_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.AwsEfsCsiAddon",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEfsCsiAddon_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.AwsEfsCsiAddon",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) AddDeletionOverride(path *string) {
	if err := a.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) AddDependency(target awscdk.CfnResource) {
	if err := a.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addDependency",
		[]interface{}{target},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) AddDependsOn(target awscdk.CfnResource) {
	if err := a.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) AddMetadata(key *string, value interface{}) {
	if err := a.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) AddPropertyDeletionOverride(propertyPath *string) {
	if err := a.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := a.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := a.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := a.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		a,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) GetMetadata(key *string) interface{} {
	if err := a.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) Inspect(inspector awscdk.TreeInspector) {
	if err := a.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"inspect",
		[]interface{}{inspector},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		a,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		a,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) RemoveDependency(target awscdk.CfnResource) {
	if err := a.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"removeDependency",
		[]interface{}{target},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := a.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := a.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (a *jsiiProxy_AwsEfsCsiAddon) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsCsiAddon) ValidateProperties(_properties interface{}) {
	if err := a.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"validateProperties",
		[]interface{}{_properties},
	)
}

