package smallcasecdkeksclustermodule


type DefaultCommonComponents struct {
	AwsEbsCsiDriver *DefaultCommonComponentsProps `field:"optional" json:"awsEbsCsiDriver" yaml:"awsEbsCsiDriver"`
	ClusterAutoscaler *DefaultCommonComponentsProps `field:"optional" json:"clusterAutoscaler" yaml:"clusterAutoscaler"`
	ExternalDns *DefaultCommonComponentsProps `field:"optional" json:"externalDns" yaml:"externalDns"`
}

