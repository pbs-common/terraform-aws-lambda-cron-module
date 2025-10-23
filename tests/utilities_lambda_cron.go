package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testLambdaCron(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
		Upgrade:      true,
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	region := getAWSRegion(t)
	accountID := getAWSAccountID(t)

	expectedLambdaName := fmt.Sprintf("ex-tf-lambda-cron-%s", variant)
	expectedLambdaARN := fmt.Sprintf("arn:aws:lambda:%s:%s:function:%s", region, accountID, expectedLambdaName)
	expectedCron := "00 7 * * ? *"

	lambdaARN := terraform.Output(t, terraformOptions, "lambda_arn")
	cron := terraform.Output(t, terraformOptions, "cron")

	assert.Equal(t, expectedLambdaARN, lambdaARN)
	assert.Equal(t, expectedCron, cron)
}
