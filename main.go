package main

import (
	"github.com/CptDelta/terraform-provider-mailtrap/mailtrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mailtrap.Provider,
	})
}
