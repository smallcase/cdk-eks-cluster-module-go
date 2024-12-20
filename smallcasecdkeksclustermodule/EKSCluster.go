package smallcasecdkeksclustermodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/internal"
)

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

	if err := validateNewEKSClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateEKSCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := e.validateAddServiceAccountWithIamRoleParameters(serviceAccountName, serviceAccountNamespace, policy); err != nil {
		panic(err)
	}
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

