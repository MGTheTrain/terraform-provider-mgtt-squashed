package mgtt

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAzurermRg() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAzurermRgCreate,
		Read:   resourceMgttAzurermRgRead,
		Update: resourceMgttAzurermRgUpdate,
		Delete: resourceMgttAzurermRgDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getResourceHandler() *AzureResourceGroupHandler {
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	return NewAzureResourceGroupHandler(subscriptionID, accessToken)
}

func createResourceGroup(name, location string, handler *AzureResourceGroupHandler) error {
	createRequestBody := map[string]interface{}{
		"location": location,
	}

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	return handler.CreateResourceGroup(name, jsonString)
}

func deleteResourceGroup(name string, handler *AzureResourceGroupHandler) error {
	return handler.DeleteResourceGroup(name)
}

// Helper functions

func extractResourceData(d *schema.ResourceData) (string, string) {
	return d.Get("name").(string), d.Get("location").(string)
}

func setResourceData(d *schema.ResourceData, name, location string) error {
	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}

	return nil
}

func extractOldResourceData(d *schema.ResourceData) string {
	oldName, _ := d.GetChange("name")
	return oldName.(string)
}

func resourceMgttAzurermRgCreate(d *schema.ResourceData, m interface{}) error {
	name, location := extractResourceData(d)
	handler := getResourceHandler()

	err := createResourceGroup(name, location, handler)
	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())

	if err := setResourceData(d, name, location); err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermRgRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	handler := getResourceHandler()

	err := handler.GetResourceGroup(name)
	if err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermRgUpdate(d *schema.ResourceData, m interface{}) error {
	oldName := extractOldResourceData(d)
	handler := getResourceHandler()

	err := deleteResourceGroup(oldName, handler)
	if err != nil {
		return err
	}

	name, location := extractResourceData(d)

	err = createResourceGroup(name, location, handler)
	if err != nil {
		return err
	}

	if err := setResourceData(d, name, location); err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermRgDelete(d *schema.ResourceData, m interface{}) error {
	name := extractOldResourceData(d)

	handler := getResourceHandler()

	err := deleteResourceGroup(name, handler)
	if err != nil {
		return err
	}
	return nil
}
