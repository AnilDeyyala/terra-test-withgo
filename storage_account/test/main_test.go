package main

import (
	// "context"
	// "encoding/json"
	// "fmt"
	// "log"
	// "os"
	// "reflect"
	// "strings"
	"testing"

	// "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	// "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	// "github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func setTerraformVariables() (map[string]string, error) {

	// Getting enVars from environment variables
	ARM_CLIENT_ID := os.Getenv("AZURE_CLIENT_ID")
	ARM_CLIENT_SECRET := os.Getenv("AZURE_CLIENT_SECRET")
	ARM_TENANT_ID := os.Getenv("AZURE_TENANT_ID")
	ARM_SUBSCRIPTION_ID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	// Creating globalEnVars for terraform call through Terratest
	if ARM_CLIENT_ID != "" {
		globalEnvVars["ARM_CLIENT_ID"] = ARM_CLIENT_ID
		globalEnvVars["ARM_CLIENT_SECRET"] = ARM_CLIENT_SECRET
		globalEnvVars["ARM_SUBSCRIPTION_ID"] = ARM_SUBSCRIPTION_ID
		globalEnvVars["ARM_TENANT_ID"] = ARM_TENANT_ID
	}

	return globalEnvVars, nil
}

func TestTerraform_storage_account(t *testing.T) {
	t.Parallel()

	setTerraformVariables()

	expectedLocation := "East US"

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../provision",

		Reconfigure: true,
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(t, terraformOptions, "hello_world")
	assert.Equal(t, "Hello, World!", output)
}

func getResourceFromRESTAPI(out string) (armresources.ResourcesGetByIDResponse, error) {

	//expected variable
	//expectedVnetName := strings.ToLower(fmt.Sprintf("%s%s%s", prefix, separator, uniquePostfix))

	log.Printf("json output: %s\n", out)

	ctx := context.Background()
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}

	resourceId := fmt.Sprintf("%v", out) //result["resource_name"]
	//resourceId := fmt.Sprintf(resourceIdFormat, subscriptionId, resource_group_name, expectedResource_name)
	// Azure SDK Azure Resource Management clients accept the credential as a parameter
	client := armresources.NewResourcesClient(subscriptionId, cred, nil)

	resp, err := client.GetByID(ctx, resourceId, apiVersion, nil)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}

	return resp, err
}
