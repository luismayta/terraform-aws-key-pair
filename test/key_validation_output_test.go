package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/stretchr/testify/assert"

	"github.com/hadenlabs/terraform-aws-key-pair/config"
	"github.com/hadenlabs/terraform-aws-key-pair/internal/app/external/faker"
	"github.com/hadenlabs/terraform-aws-key-pair/internal/common/log"
)

func TestValidationOutputSuccess(t *testing.T) {
	t.Parallel()
	conf := config.Must()
	name := faker.Key().Name()
	publicKey := "../fixtures/keys/testing.pub"
	logger := log.Factory(*conf)
	logger.Debugf(
		"values for test terraform-aws-key-pair is",
		"name", name,
	)

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "key-validation-output",
		Upgrade:      true,
		Vars: map[string]interface{}{
			"name":       name,
			"public_key": publicKey,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
	outputInstance := terraform.Output(t, terraformOptions, "instance")
	assert.NotEmpty(t, outputInstance, outputInstance)
	outputName := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, outputName, name)
}
