package mgtt

import (
	"os"
	"testing"

	"github.com/MGTheTrain/terraform-provider-mgtt-squashed/mgtt"
	"github.com/stretchr/testify/assert"
)

func TestStorageAccountContainerHandler(t *testing.T) {
	// Read parameters from environment variables
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	accountName := "testaccount54321"
	resourceGroupName := "rg-test-100"
	containerName := "test-container-8092"

	if subscriptionID == "" || resourceGroupName == "" || accountName == "" || accessToken == "" {
		t.Fatal("Missing required environment variables")
	}

	resource_group_handler := mgtt.NewAzureResourceGroupHandler(subscriptionID, accessToken)
	storage_account_handler := mgtt.NewAzureStorageAccountHandler(subscriptionID, accessToken)
	storage_account_container_handler := mgtt.NewAzureStorageAccountContainerHandler(subscriptionID, accessToken)

	createResourceGroupRequestBody := `{
		"location": "West Europe"
	}`

	createStorageAccountRequestBody := `{
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

	err = storage_account_handler.CreateStorageAccount(resourceGroupName, accountName, createStorageAccountRequestBody)
	assert.NoError(t, err, "CreateStorageAccount should not return an error")

	err = storage_account_container_handler.CreateStorageAccountContainer(resourceGroupName, accountName, containerName, `{}`)
	assert.NoError(t, err, "CreateStorageAccount should not return an error")

	// [D]elete
	err = storage_account_container_handler.DeleteStorageAccountContainer(resourceGroupName, accountName, containerName)
	assert.NoError(t, err, "DeleteStorageAccountContainer should not return an error")

	err = storage_account_handler.DeleteStorageAccount(resourceGroupName, accountName)
	assert.NoError(t, err, "DeleteStorageAccount should not return an error")

	err = resource_group_handler.DeleteResourceGroup(resourceGroupName)
	assert.NoError(t, err, "DeleteResourceGroup should not return an error")
}
