//go:build !no_runtime_type_checking

package smallcasecdkeksclustermodule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EKSCluster) validateAddServiceAccountWithIamRoleParameters(serviceAccountName *string, serviceAccountNamespace *string, policy interface{}) error {
	if serviceAccountName == nil {
		return fmt.Errorf("parameter serviceAccountName is required, but nil was provided")
	}

	if serviceAccountNamespace == nil {
		return fmt.Errorf("parameter serviceAccountNamespace is required, but nil was provided")
	}

	if policy == nil {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func validateEKSCluster_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewEKSClusterParameters(scope constructs.Construct, id *string, props *EKSClusterProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

