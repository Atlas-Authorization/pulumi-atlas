// Command pulumi-tfgen-atlas is the build-time code generator for the Pulumi
// Atlas provider. It emits the Pulumi Package Schema (schema.json) and the
// per-language SDK sources (Go, Node, Python, .NET) from the bridged Terraform
// provider's schema. Invoked by `make tfgen` and `make build_sdks`.
package main

import (
	atlas "github.com/atlas/pulumi-atlas/provider"
	pftfgen "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"
)

func main() {
	pftfgen.Main("atlas", atlas.Provider())
}
