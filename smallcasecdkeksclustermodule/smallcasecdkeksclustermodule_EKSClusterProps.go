// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

type EKSClusterProps struct {
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	ClusterConfig *ClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
	KmsKey awskms.Key `field:"required" json:"kmsKey" yaml:"kmsKey"`
	Region *string `field:"required" json:"region" yaml:"region"`
	WorkerSecurityGroup awsec2.SecurityGroup `field:"required" json:"workerSecurityGroup" yaml:"workerSecurityGroup"`
	AddonProps *AddonProps `field:"optional" json:"addonProps" yaml:"addonProps"`
	ClusterVPC awsec2.IVpc `field:"optional" json:"clusterVPC" yaml:"clusterVPC"`
	CoreDnsAddonProps *CoreAddonValuesProps `field:"optional" json:"coreDnsAddonProps" yaml:"coreDnsAddonProps"`
	KubeProxyAddonProps *CoreAddonValuesProps `field:"optional" json:"kubeProxyAddonProps" yaml:"kubeProxyAddonProps"`
}

