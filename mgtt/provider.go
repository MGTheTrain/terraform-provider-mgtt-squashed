package mgtt

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ConvertMapToJSON(input map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"mgtt_aws_s3_bucket":                     resourceMgttAwsS3Bucket(),
			"mgtt_azurerm_storage_account":           resourceMgttAzurermStorageAccount(),
			"mgtt_azurerm_storage_account_container": resourceMgttAzurermStorageAccountContainer(),
			"mgtt_azurerm_rg":                        resourceMgttAzurermRg(),
		},
	}
}
