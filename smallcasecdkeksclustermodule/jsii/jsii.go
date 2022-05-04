// Package jsii contains the functionaility needed for jsii packages to
// initialize their dependencies and themselves. Users should never need to use this package
// directly. If you find you need to - please report a bug at
// https://github.com/aws/jsii/issues/new/choose
package jsii

import (
	_                                  "embed"

	_jsii_                             "github.com/aws/jsii-runtime-go/runtime"

	awscdk                             "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	constructs                         "github.com/aws/constructs-go/constructs/v10/jsii"
	cdk8s                              "github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2/jsii"
	opencdk8scdk8sclusterautoscaleraws "github.com/opencdk8s/cdk8s-cluster-autoscaler-aws-go/opencdk8scdk8sclusterautoscaleraws/jsii"
	opencdk8scdk8sexternaldnsroute53   "github.com/opencdk8s/cdk8s-external-dns-route53-go/opencdk8scdk8sexternaldnsroute53/jsii"
)

//go:embed smallcase-cdk-eks-cluster-module-0.0.2.tgz
var tarball []byte

// Initialize loads the necessary packages in the @jsii/kernel to support the enclosing module.
// The implementation is idempotent (and hence safe to be called over and over).
func Initialize() {
	// Ensure all dependencies are initialized
	opencdk8scdk8sclusterautoscaleraws.Initialize()
	opencdk8scdk8sexternaldnsroute53.Initialize()
	awscdk.Initialize()
	cdk8s.Initialize()
	constructs.Initialize()

	// Load this library into the kernel
	_jsii_.Load("@smallcase/cdk-eks-cluster-module", "0.0.2", tarball)
}
