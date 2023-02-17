//go:build !no_runtime_type_checking

// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

import (
	"fmt"
)

func validateVpcCniAddonVersion_OfParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateNewVpcCniAddonVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

