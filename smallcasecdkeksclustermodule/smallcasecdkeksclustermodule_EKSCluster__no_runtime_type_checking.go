//go:build no_runtime_type_checking

// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EKSCluster) validateAddServiceAccountWithIamRoleParameters(serviceAccountName *string, serviceAccountNamespace *string, policy interface{}) error {
	return nil
}

func validateEKSCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEKSClusterParameters(scope constructs.Construct, id *string, props *EKSClusterProps) error {
	return nil
}

