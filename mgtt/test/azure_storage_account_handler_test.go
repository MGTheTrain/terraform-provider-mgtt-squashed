package mgtt

import (
	"os"
	"testing"

	"github.com/MGTheTrain/terraform-provider-mgtt-squashed/mgtt"
	"github.com/stretchr/testify/assert"
)

func TestStorageAccountHandler(t *testing.T) {
	// Read parameters from environment variables
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	accountName := "testaccount54321"
	resourceGroupName := "rg-test-100"

	if subscriptionID == "" || resourceGroupName == "" || accountName == "" || accessToken == "" {
		t.Fatal("Missing required environment variables")
	}

	resource_group_handler := mgtt.NewAzureResourceGroupHandler(subscriptionID, accessToken)
	handler := mgtt.NewAzureStorageAccountHandler(subscriptionID, accessToken)

	createResourceGroupRequestBody := `{
		"location": "West Europe"
	}`

	createRequestBody := `{
		"sku": {
			"name": "Standard_LRS",
			"tier": "Standard"
		},
		"kind": "StorageV2",
		"location": "West Europe"
	}`

	// [C]reate
	err := resource_group_handler.CreateResourceGroup(resourceGroupName, createResourceGroupRequestBody)
	assert.NoError(t, err, "CreateResourceGroup should not return an error")

	err = handler.CreateStorageAccount(resourceGroupName, accountName, createRequestBody)
	assert.NoError(t, err, "CreateStorageAccount should not return an error")

	// [R]ead
	err = handler.GetStorageAccount(resourceGroupName, accountName)
	assert.NoError(t, err, "GetStorageAccount should not return an error")

	// [U]pdate -> Strive for immutability in your infrastructure deployments. Instead of making in-place updates, destroy and recreate resources when changes are required.
	err = handler.DeleteStorageAccount(resourceGroupName, accountName)
	assert.NoError(t, err, "DeleteStorageAccount should not return an error")

	newAccountName := "testaccount09876"
	err = handler.CreateStorageAccount(resourceGroupName, newAccountName, createRequestBody)
	assert.NoError(t, err, "CreateStorageAccount should not return an error")

	// [D]elete
	err = handler.DeleteStorageAccount(resourceGroupName, newAccountName)
	assert.NoError(t, err, "DeleteStorageAccount should not return an error")

	err = resource_group_handler.DeleteResourceGroup(resourceGroupName)
	assert.NoError(t, err, "DeleteResourceGroup should not return an error")
}
