// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule


type DefaultCommonComponents struct {
	AwsEbsCsiDriver *DefaultCommonComponentsProps `field:"optional" json:"awsEbsCsiDriver" yaml:"awsEbsCsiDriver"`
	AwsEfsCsiDriver *DefaultCommonComponentsProps `field:"optional" json:"awsEfsCsiDriver" yaml:"awsEfsCsiDriver"`
	ClusterAutoscaler *DefaultCommonComponentsProps `field:"optional" json:"clusterAutoscaler" yaml:"clusterAutoscaler"`
	ExternalDns *DefaultCommonComponentsProps `field:"optional" json:"externalDns" yaml:"externalDns"`
}

