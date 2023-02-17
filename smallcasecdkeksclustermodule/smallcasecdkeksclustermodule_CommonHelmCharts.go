// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/internal"
)

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

	if err := validateNewCommonHelmChartsParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateCommonHelmCharts_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

