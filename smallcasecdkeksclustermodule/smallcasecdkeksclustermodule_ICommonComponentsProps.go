// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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

func (j *jsiiProxy_ICommonComponentsProps)SetHelm(val *StandardHelmProps) {
	if err := j.validateSetHelmParameters(val); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_ICommonComponentsProps)SetIamPolicyPath(val *[]*string) {
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

func (j *jsiiProxy_ICommonComponentsProps)SetServiceAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"serviceAccounts",
		val,
	)
}

