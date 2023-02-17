// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule


type StandardHelmProps struct {
	ChartName *string `field:"required" json:"chartName" yaml:"chartName"`
	ChartReleaseName *string `field:"optional" json:"chartReleaseName" yaml:"chartReleaseName"`
	ChartVersion *string `field:"optional" json:"chartVersion" yaml:"chartVersion"`
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	HelmRepository *string `field:"optional" json:"helmRepository" yaml:"helmRepository"`
	HelmValues *map[string]interface{} `field:"optional" json:"helmValues" yaml:"helmValues"`
	LocalHelmChart *string `field:"optional" json:"localHelmChart" yaml:"localHelmChart"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

