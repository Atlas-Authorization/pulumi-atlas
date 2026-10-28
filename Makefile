PACK            := atlas
PROJECT         := github.com/atlas/pulumi-atlas
PROVIDER        := pulumi-resource-$(PACK)
TFGEN           := pulumi-tfgen-$(PACK)
VERSION         ?= 0.15.0
PROVIDER_PATH   := provider
LDFLAGS         := -X $(PROJECT)/provider.Version=$(VERSION)

SCHEMA_FILE     := provider/cmd/$(PROVIDER)/schema.json
WORKING_DIR     := $(shell pwd)

# Languages produced by build_sdks. Each needs the Pulumi toolchain + the
# matching language SDK codegen (installed by `pulumi plugin`/`pulumi` CLI).
LANGS           := nodejs python go dotnet

.PHONY: default ensure provider tfgen schema build_sdks \
        build_nodejs build_python build_go build_dotnet clean help

default: provider

help:
	@echo "Targets:"
	@echo "  tfgen        build the pulumi-tfgen-$(PACK) codegen binary"
	@echo "  schema       (re)generate provider/cmd/$(PROVIDER)/schema.json"
	@echo "  provider     build the pulumi-resource-$(PACK) plugin binary"
	@echo "  build_sdks   generate SDKs for: $(LANGS)"
	@echo "  build_<lang> generate a single SDK (nodejs|python|go|dotnet)"
	@echo "  clean        remove build output (bin/, sdk/)"

ensure:
	cd $(PROVIDER_PATH) && go mod tidy

# Build the code generator. Pure Go — needs no external tools.
tfgen:
	cd $(PROVIDER_PATH) && go build -o $(WORKING_DIR)/bin/$(TFGEN) \
		-ldflags "$(LDFLAGS)" ./cmd/$(TFGEN)

# Regenerate the embedded Pulumi Package Schema from the bridged TF provider.
# Pure Go — needs no external tools.
schema: tfgen
	$(WORKING_DIR)/bin/$(TFGEN) schema --out $(PROVIDER_PATH)/cmd/$(PROVIDER)
	# tfgen drops the top-level "version" the Pulumi registry needs; put it back.
	node scripts/set-schema-version.mjs $(SCHEMA_FILE) $(VERSION)

# Build the runtime resource plugin. Pure Go — embeds schema.json (run `make
# schema` first if the mapping changed).
provider:
	cd $(PROVIDER_PATH) && go build -o $(WORKING_DIR)/bin/$(PROVIDER) \
		-ldflags "$(LDFLAGS)" ./cmd/$(PROVIDER)

# --- SDK generation --------------------------------------------------------
# These require the Pulumi CLI (`pulumi`) on PATH plus the per-language SDK
# codegen it ships. The tfgen binary emits the sources; language build tooling
# (npm/pip/go/dotnet) then packages them. See README "Generating the SDKs".
build_sdks: build_nodejs build_python build_go build_dotnet

build_nodejs: tfgen
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --out sdk/nodejs

build_python: tfgen
	$(WORKING_DIR)/bin/$(TFGEN) python --out sdk/python

build_go: tfgen
	$(WORKING_DIR)/bin/$(TFGEN) go --out sdk/go

build_dotnet: tfgen
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --out sdk/dotnet
	# The generated csproj embeds version.txt as a resource; `dotnet build`
	# fails without it.
	echo "$(VERSION)" > sdk/dotnet/version.txt

clean:
	rm -rf bin sdk
