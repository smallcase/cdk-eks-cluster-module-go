// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type FargateProfile struct {
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	Labels *InternalMap `field:"optional" json:"labels" yaml:"labels"`
	PodExecutionRole awsiam.Role `field:"optional" json:"podExecutionRole" yaml:"podExecutionRole"`
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

