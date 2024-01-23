package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	subscriptionID    string
	resourceGroupName string
	accountName       string
	accessToken       string
	requestBodyFile   string
)

func main() {
	var rootCmd = &cobra.Command{Use: "azure_storage_account_handler"}

	rootCmd.PersistentFlags().StringVarP(&subscriptionID, "subscriptionID", "s", "", "Your Azure subscription ID")
	rootCmd.PersistentFlags().StringVarP(&resourceGroupName, "resourceGroupName", "g", "", "Your Azure resource group name")
	rootCmd.PersistentFlags().StringVarP(&accountName, "accountName", "a", "", "Your Azure Storage account name")
	rootCmd.PersistentFlags().StringVarP(&accessToken, "accessToken", "t", "", "Your Bearer access token")
	rootCmd.PersistentFlags().StringVarP(&requestBodyFile, "requestBodyFile", "r", "assets/create_azure_storage_account_request_body.json", "Path to JSON file containing the request body")

	// Create a new storage_account command
	storageAccountCmd := &cobra.Command{
		Use:   "storage_account",
		Short: "Manage Azure Storage Account",
	}

	// Add subcommands to the storage_account command
	storageAccountCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create Azure Storage Account",
		Run: func(cmd *cobra.Command, args []string) {
			handleCreateAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken, requestBodyFile)
		},
	})

	storageAccountCmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get an Azure Storage Account by account name",
		Run: func(cmd *cobra.Command, args []string) {
			handleReadAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken)
		},
	})

	storageAccountCmd.AddCommand(&cobra.Command{
		Use:   "update",
		Short: "Update an Azure Storage Account by account name",
		Run: func(cmd *cobra.Command, args []string) {
			handleUpdateAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken, requestBodyFile)
		},
	})

	storageAccountCmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete an Azure Storage Account by account name",
		Run: func(cmd *cobra.Command, args []string) {
			handleDeleteAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken)
		},
	})

	// Add storage_account command to the root command
	rootCmd.AddCommand(storageAccountCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleCreateAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken, requestBodyFile string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	requestBody, err := readRequestBodyFromFile(requestBodyFile)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return
	}

	sendHTTPRequest("PUT", url, requestBody, accessToken)
}

func handleReadAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	sendHTTPRequest("GET", url, nil, accessToken)
}

func handleUpdateAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken, requestBodyFile string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	requestBody, err := readRequestBodyFromFile(requestBodyFile)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return
	}

	sendHTTPRequest("PATCH", url, requestBody, accessToken)
}

func handleDeleteAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	sendHTTPRequest("DELETE", url, nil, accessToken)
}

func readRequestBodyFromFile(filePath string) (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var requestBody map[string]interface{}
	if err := json.Unmarshal(content, &requestBody); err != nil {
		return nil, err
	}

	return requestBody, nil
}

func sendHTTPRequest(method, url string, requestBody map[string]interface{}, accessToken string) {
	var req *http.Request
	var err error

	if requestBody != nil {
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading from buffer:", err)
		return
	}
	fmt.Println(buf.String())
}
