// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule


type ArgoCD struct {
	AssumeRoleArn *string `field:"required" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	ClusterRoleName *string `field:"required" json:"clusterRoleName" yaml:"clusterRoleName"`
}

