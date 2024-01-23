package mgtt

import (
	"bytes"
	"fmt"
	"net/http"
)

// AzureResourceGroupHandler represents a manager for handling Azure Storage operations.
type AzureResourceGroupHandler struct {
	SubscriptionID string
	AccessToken    string
}

// NewAzureResourceGroupHandler creates a new instance of AzureResourceGroupHandler.
func NewAzureResourceGroupHandler(subscriptionID, accessToken string) *AzureResourceGroupHandler {
	return &AzureResourceGroupHandler{
		SubscriptionID: subscriptionID,
		AccessToken:    accessToken,
	}
}

// CreateResourceGroup creates an Azure Storage account.
func (m *AzureResourceGroupHandler) CreateResourceGroup(resourceGroupName, requestBody string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s?api-version=2021-04-01",
		m.SubscriptionID, resourceGroupName)

	return m.sendHTTPRequest("PUT", url, []byte(requestBody), m.AccessToken)
}

// GetResourceGroup reads information about an Azure Storage account.
func (m *AzureResourceGroupHandler) GetResourceGroup(resourceGroupName string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s?api-version=2021-04-01",
		m.SubscriptionID, resourceGroupName)

	return m.sendHTTPRequest("GET", url, nil, m.AccessToken)
}

// // UpdateResourceGroup updates an Azure Storage account.
// NOTE: We can not rename the resource group with patch. Prefer deleting the resource group and recreating it.
// 		"Strive for immutability in your infrastructure deployments. Instead of making in-place updates, destroy and recreate resources when changes are required."
// func (m *AzureResourceGroupHandler) UpdateResourceGroup(resourceGroupName, requestBody string) error {
// 	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s?api-version=2021-04-01",
// 		m.SubscriptionID, resourceGroupName)

// 	return m.sendHTTPRequest("PATCH", url, []byte(requestBody), m.AccessToken)
// }

// DeleteResourceGroup deletes an Azure Storage account.
func (m *AzureResourceGroupHandler) DeleteResourceGroup(resourceGroupName string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s?api-version=2021-04-01",
		m.SubscriptionID, resourceGroupName)

	return m.sendHTTPRequest("DELETE", url, nil, m.AccessToken)
}

// sendHTTPRequest sends an HTTP request.
func (m *AzureResourceGroupHandler) sendHTTPRequest(method, url string, requestBody []byte, accessToken string) error {
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
