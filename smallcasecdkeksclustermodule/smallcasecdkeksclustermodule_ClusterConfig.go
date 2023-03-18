// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type ClusterConfig struct {
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	ClusterVersion awseks.KubernetesVersion `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	DefaultCapacity *float64 `field:"required" json:"defaultCapacity" yaml:"defaultCapacity"`
	NodeGroups *[]*NodeGroupConfig `field:"required" json:"nodeGroups" yaml:"nodeGroups"`
	Subnets *InternalMap `field:"required" json:"subnets" yaml:"subnets"`
	Tags *InternalMap `field:"required" json:"tags" yaml:"tags"`
	TeamMembers *[]*string `field:"required" json:"teamMembers" yaml:"teamMembers"`
	AddAutoscalerIam *bool `field:"optional" json:"addAutoscalerIam" yaml:"addAutoscalerIam"`
	AlbControllerVersion awseks.AlbControllerVersion `field:"optional" json:"albControllerVersion" yaml:"albControllerVersion"`
	ArgoCD *ArgoCD `field:"optional" json:"argoCD" yaml:"argoCD"`
	CommonComponents *map[string]ICommonComponentsProps `field:"optional" json:"commonComponents" yaml:"commonComponents"`
	DebugLogs *bool `field:"optional" json:"debugLogs" yaml:"debugLogs"`
	DefaultCommonComponents *DefaultCommonComponents `field:"optional" json:"defaultCommonComponents" yaml:"defaultCommonComponents"`
	DeprecateClusterAutoScaler *bool `field:"optional" json:"deprecateClusterAutoScaler" yaml:"deprecateClusterAutoScaler"`
	FargateProfiles *[]*FargateProfile `field:"optional" json:"fargateProfiles" yaml:"fargateProfiles"`
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	Namespaces *map[string]*NamespaceSpec `field:"optional" json:"namespaces" yaml:"namespaces"`
	PublicAllowAccess *[]*string `field:"optional" json:"publicAllowAccess" yaml:"publicAllowAccess"`
	TeamExistingRolePermission *map[string]*string `field:"optional" json:"teamExistingRolePermission" yaml:"teamExistingRolePermission"`
}

