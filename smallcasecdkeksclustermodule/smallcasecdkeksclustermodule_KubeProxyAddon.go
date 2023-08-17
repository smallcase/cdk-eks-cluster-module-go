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

type KubeProxyAddon interface {
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

// The jsii proxy struct for KubeProxyAddon
type jsiiProxy_KubeProxyAddon struct {
	internal.Type__awseksCfnAddon
}

func (j *jsiiProxy_KubeProxyAddon) AddonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) AddonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) ConfigurationValues() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) PreserveOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) ResolveConflicts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveConflicts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) ServiceAccountRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeProxyAddon) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewKubeProxyAddon(scope constructs.Construct, id *string, props *CoreAddonProps) KubeProxyAddon {
	_init_.Initialize()

	if err := validateNewKubeProxyAddonParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubeProxyAddon{}

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.KubeProxyAddon",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKubeProxyAddon_Override(k KubeProxyAddon, scope constructs.Construct, id *string, props *CoreAddonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@smallcase/cdk-eks-cluster-module.KubeProxyAddon",
		[]interface{}{scope, id, props},
		k,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetAddonName(val *string) {
	if err := j.validateSetAddonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addonName",
		val,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetAddonVersion(val *string) {
	_jsii_.Set(
		j,
		"addonVersion",
		val,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetConfigurationValues(val *string) {
	_jsii_.Set(
		j,
		"configurationValues",
		val,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetPreserveOnDelete(val interface{}) {
	if err := j.validateSetPreserveOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveOnDelete",
		val,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetResolveConflicts(val *string) {
	_jsii_.Set(
		j,
		"resolveConflicts",
		val,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetServiceAccountRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountRoleArn",
		val,
	)
}

func (j *jsiiProxy_KubeProxyAddon)SetTagsRaw(val *[]*awscdk.CfnTag) {
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
func KubeProxyAddon_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeProxyAddon_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.KubeProxyAddon",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func KubeProxyAddon_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKubeProxyAddon_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.KubeProxyAddon",
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
func KubeProxyAddon_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeProxyAddon_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@smallcase/cdk-eks-cluster-module.KubeProxyAddon",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubeProxyAddon_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.KubeProxyAddon",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubeProxyAddon) AddDeletionOverride(path *string) {
	if err := k.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (k *jsiiProxy_KubeProxyAddon) AddDependency(target awscdk.CfnResource) {
	if err := k.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addDependency",
		[]interface{}{target},
	)
}

func (k *jsiiProxy_KubeProxyAddon) AddDependsOn(target awscdk.CfnResource) {
	if err := k.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (k *jsiiProxy_KubeProxyAddon) AddMetadata(key *string, value interface{}) {
	if err := k.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (k *jsiiProxy_KubeProxyAddon) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubeProxyAddon) AddPropertyDeletionOverride(propertyPath *string) {
	if err := k.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (k *jsiiProxy_KubeProxyAddon) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := k.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (k *jsiiProxy_KubeProxyAddon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := k.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (k *jsiiProxy_KubeProxyAddon) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := k.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		k,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeProxyAddon) GetMetadata(key *string) interface{} {
	if err := k.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeProxyAddon) Inspect(inspector awscdk.TreeInspector) {
	if err := k.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"inspect",
		[]interface{}{inspector},
	)
}

func (k *jsiiProxy_KubeProxyAddon) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		k,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeProxyAddon) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		k,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeProxyAddon) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubeProxyAddon) RemoveDependency(target awscdk.CfnResource) {
	if err := k.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"removeDependency",
		[]interface{}{target},
	)
}

func (k *jsiiProxy_KubeProxyAddon) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := k.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeProxyAddon) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := k.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (k *jsiiProxy_KubeProxyAddon) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		k,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeProxyAddon) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeProxyAddon) ValidateProperties(_properties interface{}) {
	if err := k.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"validateProperties",
		[]interface{}{_properties},
	)
}

