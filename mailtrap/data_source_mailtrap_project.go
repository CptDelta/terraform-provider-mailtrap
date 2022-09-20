package mailtrap

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func projectDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "",
		ReadContext: projectList,

		Schema: map[string]*schema.Schema{
			"Id": {
				Type:        schema.TypeInt,
				Description: "Id of a Mailtrap Project",
			},
		},
	}
}

func projectList(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}
