// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

type NodeGroupConfig struct {
	InstanceTypes *[]awsec2.InstanceType `field:"required" json:"instanceTypes" yaml:"instanceTypes"`
	Labels *InternalMap `field:"required" json:"labels" yaml:"labels"`
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
	Name *string `field:"required" json:"name" yaml:"name"`
	SubnetGroupName *string `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
	Taints *InternalMap `field:"required" json:"taints" yaml:"taints"`
	CapacityType awseks.CapacityType `field:"optional" json:"capacityType" yaml:"capacityType"`
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	LaunchTemplateSpec *awseks.LaunchTemplateSpec `field:"optional" json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	NodeAMIVersion *string `field:"optional" json:"nodeAMIVersion" yaml:"nodeAMIVersion"`
	SshKeyName *string `field:"optional" json:"sshKeyName" yaml:"sshKeyName"`
	SubnetAz *[]*string `field:"optional" json:"subnetAz" yaml:"subnetAz"`
	Tags *InternalMap `field:"optional" json:"tags" yaml:"tags"`
}

