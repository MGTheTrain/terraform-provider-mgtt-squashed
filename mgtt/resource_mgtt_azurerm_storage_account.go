package mgtt

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAzurermStorageAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAzurermStorageAccountCreate,
		Read:   resourceMgttAzurermStorageAccountRead,
		Update: resourceMgttAzurermStorageAccountUpdate,
		Delete: resourceMgttAzurermStorageAccountDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sku_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sku_tier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"kind": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getStorageAccountHandler() *AzureStorageAccountHandler {
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	return NewAzureStorageAccountHandler(subscriptionID, accessToken)
}

func createStorageAccount(name, resourceGroupName, location, kind, skuName, skuTier string, handler *AzureStorageAccountHandler) error {
	createRequestBody := map[string]interface{}{
		"sku": map[string]interface{}{
			"name": skuName,
			"tier": skuTier,
		},
		"kind":     kind,
		"location": location,
	}

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	return handler.CreateStorageAccount(resourceGroupName, name, jsonString)
}

func deleteStorageAccount(resourceGroupName, name string, handler *AzureStorageAccountHandler) error {
	return handler.DeleteStorageAccount(resourceGroupName, name)
}

// Helper functions

func extractStorageAccountData(d *schema.ResourceData) (string, string, string, string, string, string) {
	return d.Get("name").(string), d.Get("resource_group_name").(string), d.Get("location").(string),
		d.Get("kind").(string), d.Get("sku_name").(string), d.Get("sku_tier").(string)
}

func setStorageAccountData(d *schema.ResourceData, name, resourceGroupName, location, kind, skuName, skuTier string) error {
	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}
	if err := d.Set("kind", kind); err != nil {
		return err
	}
	if err := d.Set("sku_name", skuName); err != nil {
		return err
	}
	if err := d.Set("sku_tier", skuTier); err != nil {
		return err
	}

	return nil
}

func extractOldStorageAccountData(d *schema.ResourceData) (string, string) {
	oldName, _ := d.GetChange("name")
	oldResourceGroupName, _ := d.GetChange("resource_group_name")

	return oldResourceGroupName.(string), oldName.(string)
}

func resourceMgttAzurermStorageAccountCreate(d *schema.ResourceData, m interface{}) error {
	name, resourceGroupName, location, kind, skuName, skuTier := extractStorageAccountData(d)
	handler := getStorageAccountHandler()

	err := createStorageAccount(name, resourceGroupName, location, kind, skuName, skuTier, handler)
	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())

	if err := setStorageAccountData(d, name, resourceGroupName, location, kind, skuName, skuTier); err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermStorageAccountRead(d *schema.ResourceData, m interface{}) error {
	name, resourceGroupName, _, _, _, _ := extractStorageAccountData(d)
	handler := getStorageAccountHandler()

	err := handler.GetStorageAccount(resourceGroupName, name)
	if err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermStorageAccountUpdate(d *schema.ResourceData, m interface{}) error {
	oldResourceGroupName, oldName := extractOldStorageAccountData(d)
	handler := getStorageAccountHandler()

	err := deleteStorageAccount(oldResourceGroupName, oldName, handler)
	if err != nil {
		return err
	}

	name, resourceGroupName, location, kind, skuName, skuTier := extractStorageAccountData(d)

	err = createStorageAccount(name, resourceGroupName, location, kind, skuName, skuTier, handler)
	if err != nil {
		return err
	}

	if err := setStorageAccountData(d, name, resourceGroupName, location, kind, skuName, skuTier); err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermStorageAccountDelete(d *schema.ResourceData, m interface{}) error {
	name, resourceGroupName := extractOldStorageAccountData(d)
	handler := getStorageAccountHandler()

	err := deleteStorageAccount(resourceGroupName, name, handler)
	if err != nil {
		return err
	}
	return nil
}
