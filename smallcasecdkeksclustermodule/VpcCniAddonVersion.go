package smallcasecdkeksclustermodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smallcase/cdk-eks-cluster-module-go/smallcasecdkeksclustermodule/jsii"
)

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

	if err := validateNewVpcCniAddonVersionParameters(version); err != nil {
		panic(err)
	}
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

	if err := validateVpcCniAddonVersion_OfParameters(version); err != nil {
		panic(err)
	}
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

func VpcCniAddonVersion_V1_11_2() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_11_2",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_11_3() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_11_3",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_11_4() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_11_4",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_12_0() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_12_0",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_12_1() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_12_1",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_12_2() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_12_2",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_12_5() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_12_5",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_12_5_2() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_12_5_2",
		&returns,
	)
	return returns
}

func VpcCniAddonVersion_V1_17_1_1() VpcCniAddonVersion {
	_init_.Initialize()
	var returns VpcCniAddonVersion
	_jsii_.StaticGet(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		"V1_17_1_1",
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

