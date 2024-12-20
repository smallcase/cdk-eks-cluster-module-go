package smallcasecdkeksclustermodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

type CoreAddonProps struct {
	Cluster awseks.Cluster `field:"required" json:"cluster" yaml:"cluster"`
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	ResolveConflicts *bool `field:"optional" json:"resolveConflicts" yaml:"resolveConflicts"`
}

