//go:build no_runtime_type_checking

// @smallcase/cdk-eks-cluster-module
package smallcasecdkeksclustermodule

// Building without runtime type checking enabled, so all the below just return nil

func validateCommonHelmCharts_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCommonHelmChartsParameters(scope constructs.Construct, id *string, props *CommonHelmChartsProps) error {
	return nil
}

