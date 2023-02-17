// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule


type AddonProps struct {
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	VpnCniAddonVersion VpcCniAddonVersion `field:"optional" json:"vpnCniAddonVersion" yaml:"vpnCniAddonVersion"`
}

