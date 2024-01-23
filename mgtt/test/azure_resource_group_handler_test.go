package mgtt

import (
	"os"
	"testing"

	"github.com/MGTheTrain/terraform-provider-mgtt/mgtt"
	"github.com/stretchr/testify/assert"
)

func TestAzureResourceGroupHandler(t *testing.T) {
	// Read parameters from environment variables
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	resourceGroupName := "rg-test-100"

	if subscriptionID == "" || resourceGroupName == "" || accessToken == "" {
		t.Fatal("Missing required environment variables")
	}

	handler := mgtt.NewAzureResourceGroupHandler(subscriptionID, accessToken)

	createRequestBody := `{
		"location": "West Europe"
	}`

	// [C]reate
	err := handler.CreateResourceGroup(resourceGroupName, createRequestBody)
	assert.NoError(t, err, "CreateResourceGroup should not return an error")

	// [R]ead
	err = handler.GetResourceGroup(resourceGroupName)
	assert.NoError(t, err, "GetResourceGroup should not return an error")

	// [U]pdate -> Strive for immutability in your infrastructure deployments. Instead of making in-place updates, destroy and recreate resources when changes are required.
	err = handler.DeleteResourceGroup(resourceGroupName)
	assert.NoError(t, err, "DeleteResourceGroup should not return an error")

	newResourceGroupName := "rg-test1000"
	err = handler.CreateResourceGroup(newResourceGroupName, createRequestBody)
	assert.NoError(t, err, "CreateResourceGroup should not return an error")

	// [D]elete
	err = handler.DeleteResourceGroup(newResourceGroupName)
	assert.NoError(t, err, "DeleteResourceGroup should not return an error")
}
