package smallcasecdkeksclustermodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

type CommonHelmChartsProps struct {
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	HelmProps *StandardHelmProps `field:"required" json:"helmProps" yaml:"helmProps"`
	DependentNamespaces *[]awseks.KubernetesManifest `field:"optional" json:"dependentNamespaces" yaml:"dependentNamespaces"`
	IamPolicyPath *[]*string `field:"optional" json:"iamPolicyPath" yaml:"iamPolicyPath"`
	LogCharts *bool `field:"optional" json:"logCharts" yaml:"logCharts"`
	ServiceAccounts *[]*string `field:"optional" json:"serviceAccounts" yaml:"serviceAccounts"`
}

