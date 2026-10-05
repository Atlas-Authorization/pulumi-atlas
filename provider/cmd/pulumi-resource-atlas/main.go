// Command pulumi-resource-atlas is the runtime Pulumi resource provider plugin
// for Atlas. Pulumi invokes it to create, read, update and delete Atlas
// resources; it bridges the Atlas Terraform provider via pulumi-terraform-bridge.
package main

import (
	"context"
	_ "embed"

	atlas "github.com/atlas/pulumi-atlas/provider"
	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
)

// schema is the Pulumi Package Schema for the bridged Atlas provider. It is
// generated from the Terraform provider's schema by `make schema`
// (pulumi-tfgen-atlas) and embedded here; regenerate it whenever the resource
// mapping or upstream schema changes.
//
//go:embed schema.json
var schema []byte

func main() {
	meta := pf.ProviderMetadata{PackageSchema: schema}
	pf.Main(context.Background(), "atlas", atlas.Provider(), meta)
}
