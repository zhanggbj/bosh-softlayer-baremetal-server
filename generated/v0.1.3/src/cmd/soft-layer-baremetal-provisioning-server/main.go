package main

import (
	"log"
	"os"

	spec "github.com/go-swagger/go-swagger/spec"
	flags "github.com/jessevdk/go-flags"

	"restapi"
	"restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {
	swaggerSpec, err := spec.New(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSoftLayerBaremetalProvisioningAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = `SoftLayer Baremetal Provisioning`
	parser.LongDescription = `BOSH SoftLayer baremetal provisioning (BMP) server APIs`

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
	}

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
