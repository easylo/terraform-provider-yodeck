package main

import (
	"github.com/easylo/terraform-provider-yodeck/yodeck"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: yodeck.Provider})
}
