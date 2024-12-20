//go:build !no_runtime_type_checking

package smallcasecdkeksclustermodule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_ICommonComponentsProps) validateSetHelmParameters(val *StandardHelmProps) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

