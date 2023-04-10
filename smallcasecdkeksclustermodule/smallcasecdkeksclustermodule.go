package smallcasecdkeksclustermodule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.AddonProps",
		reflect.TypeOf((*AddonProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.ArgoCD",
		reflect.TypeOf((*ArgoCD)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.ClusterConfig",
		reflect.TypeOf((*ClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-eks-cluster-module.CommonHelmCharts",
		reflect.TypeOf((*CommonHelmCharts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CommonHelmCharts{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.CommonHelmChartsProps",
		reflect.TypeOf((*CommonHelmChartsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.CoreAddonProps",
		reflect.TypeOf((*CoreAddonProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.CoreAddonValuesProps",
		reflect.TypeOf((*CoreAddonValuesProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-eks-cluster-module.CoreDnsAddon",
		reflect.TypeOf((*CoreDnsAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "addonName", GoGetter: "AddonName"},
			_jsii_.MemberProperty{JsiiProperty: "addonVersion", GoGetter: "AddonVersion"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "configurationValues", GoGetter: "ConfigurationValues"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preserveOnDelete", GoGetter: "PreserveOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "resolveConflicts", GoGetter: "ResolveConflicts"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountRoleArn", GoGetter: "ServiceAccountRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CoreDnsAddon{}
			_jsii_.InitJsiiProxy(&j.Type__awseksCfnAddon)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.DefaultCommonComponents",
		reflect.TypeOf((*DefaultCommonComponents)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.DefaultCommonComponentsProps",
		reflect.TypeOf((*DefaultCommonComponentsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-eks-cluster-module.EKSCluster",
		reflect.TypeOf((*EKSCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalFargateProfile", GoGetter: "AdditionalFargateProfile"},
			_jsii_.MemberProperty{JsiiProperty: "additionalNodegroups", GoGetter: "AdditionalNodegroups"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAccountWithIamRole", GoMethod: "AddServiceAccountWithIamRole"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "fargateProfiles", GoGetter: "FargateProfiles"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EKSCluster{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.EKSClusterProps",
		reflect.TypeOf((*EKSClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.FargateProfile",
		reflect.TypeOf((*FargateProfile)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@smallcase/cdk-eks-cluster-module.ICommonComponentsProps",
		reflect.TypeOf((*ICommonComponentsProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "helm", GoGetter: "Helm"},
			_jsii_.MemberProperty{JsiiProperty: "iamPolicyPath", GoGetter: "IamPolicyPath"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccounts", GoGetter: "ServiceAccounts"},
		},
		func() interface{} {
			return &jsiiProxy_ICommonComponentsProps{}
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.InternalMap",
		reflect.TypeOf((*InternalMap)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-eks-cluster-module.KubeProxyAddon",
		reflect.TypeOf((*KubeProxyAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "addonName", GoGetter: "AddonName"},
			_jsii_.MemberProperty{JsiiProperty: "addonVersion", GoGetter: "AddonVersion"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "configurationValues", GoGetter: "ConfigurationValues"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preserveOnDelete", GoGetter: "PreserveOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "resolveConflicts", GoGetter: "ResolveConflicts"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountRoleArn", GoGetter: "ServiceAccountRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_KubeProxyAddon{}
			_jsii_.InitJsiiProxy(&j.Type__awseksCfnAddon)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.NamespaceSpec",
		reflect.TypeOf((*NamespaceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.NodeGroupConfig",
		reflect.TypeOf((*NodeGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.StandardHelmProps",
		reflect.TypeOf((*StandardHelmProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonProps",
		reflect.TypeOf((*VpcCniAddonProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-eks-cluster-module.VpcCniAddonVersion",
		reflect.TypeOf((*VpcCniAddonVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_VpcCniAddonVersion{}
		},
	)
	_jsii_.RegisterClass(
		"@smallcase/cdk-eks-cluster-module.VpcEniAddon",
		reflect.TypeOf((*VpcEniAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "addonName", GoGetter: "AddonName"},
			_jsii_.MemberProperty{JsiiProperty: "addonVersion", GoGetter: "AddonVersion"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "configurationValues", GoGetter: "ConfigurationValues"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preserveOnDelete", GoGetter: "PreserveOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "resolveConflicts", GoGetter: "ResolveConflicts"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountRoleArn", GoGetter: "ServiceAccountRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_VpcEniAddon{}
			_jsii_.InitJsiiProxy(&j.Type__awseksCfnAddon)
			return &j
		},
	)
}
