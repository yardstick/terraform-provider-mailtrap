package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/yardstick/terraform-provider-mailtrap/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
