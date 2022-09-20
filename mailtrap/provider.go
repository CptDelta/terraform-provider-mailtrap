package mailtrap

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mailtrap-project": resourceProject(),
			"mailtrap-inbox":   resourceInbox(),
		},
		ConfigureContextFunc: configureFunc,
	}
}

func configureFunc(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return nil, nil
}

func resourceInbox() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
				ValidateFunc: func(i interface{}, s string) ([]string, []error) {
					return nil, nil
				},
			},
		},
	}
}

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
				ForceNew:    true,
				ValidateFunc: func(i interface{}, s string) ([]string, []error) {
					return nil, nil
				},
			},
		},
	}
}
