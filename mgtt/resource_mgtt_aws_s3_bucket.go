package mgtt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAwsS3Bucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAwsS3BucketCreate,
		Read:   resourceMgttAwsS3BucketRead,
		Update: resourceMgttAwsS3BucketUpdate,
		Delete: resourceMgttAwsS3BucketDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMgttAwsS3BucketCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	if err := d.Set("name", name); err != nil {
		return err
	}
	d.SetId(name)
	return nil
}

func resourceMgttAwsS3BucketRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAwsS3BucketUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAwsS3BucketDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
