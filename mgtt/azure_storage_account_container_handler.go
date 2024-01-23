package mgtt

import (
	"bytes"
	"fmt"
	"net/http"
)

// AzureStorageAccountContainerHandler represents a manager for handling Azure Storage operations.
type AzureStorageAccountContainerHandler struct {
	SubscriptionID string
	AccessToken    string
}

// NewAzureStorageAccountContainerHandler creates a new instance of AzureStorageAccountContainerHandler.
func NewAzureStorageAccountContainerHandler(subscriptionID, accessToken string) *AzureStorageAccountContainerHandler {
	return &AzureStorageAccountContainerHandler{
		SubscriptionID: subscriptionID,
		AccessToken:    accessToken,
	}
}

// CreateAzureStorageAccount creates an Azure Storage account.
func (m *AzureStorageAccountContainerHandler) CreateStorageAccountContainer(resourceGroupName, accountName, containerName, requestBody string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/default/containers/%s?api-version=2023-01-01",
		m.SubscriptionID, resourceGroupName, accountName, containerName)

	return m.sendHTTPRequest("PUT", url, []byte(requestBody), m.AccessToken)
}

// GetAzureStorageAccount reads information about an Azure Storage account.
func (m *AzureStorageAccountContainerHandler) GetStorageAccountContainer(resourceGroupName, accountName, containerName string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/default/containers/%s?api-version=2023-01-01",
		m.SubscriptionID, resourceGroupName, accountName, containerName)

	return m.sendHTTPRequest("GET", url, nil, m.AccessToken)
}

// DeleteAzureStorageAccount deletes an Azure Storage account.
func (m *AzureStorageAccountContainerHandler) DeleteStorageAccountContainer(resourceGroupName, accountName, containerName string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/default/containers/%s?api-version=2023-01-01",
		m.SubscriptionID, resourceGroupName, accountName, containerName)

	return m.sendHTTPRequest("DELETE", url, nil, m.AccessToken)
}

// sendHTTPRequest sends an HTTP request.
func (m *AzureStorageAccountContainerHandler) sendHTTPRequest(method, url string, requestBody []byte, accessToken string) error {
	var req *http.Request
	var err error

	if requestBody != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return fmt.Errorf("Error creating request:  %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error making request  %s:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading from buffer  %s:", err)
	}
	fmt.Println(buf.String())

	return nil
}
