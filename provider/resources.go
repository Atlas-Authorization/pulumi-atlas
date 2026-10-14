// Package atlas contains the Pulumi provider for Atlas. It is a thin bridge
// over the Atlas Terraform provider (terraform-provider-atlas) built with
// pulumi-terraform-bridge: the Terraform resource/data-source schemas are
// reused verbatim and surfaced to Pulumi's Go/Node/Python/.NET SDKs.
package atlas

import (
	tfshim "github.com/atlas/terraform-provider-atlas/shim"
	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

// Package/module identifiers for the generated Pulumi tokens.
const (
	mainPkg = "atlas"
	mainMod = "index" // all Atlas resources live in the root module.
)

// Version is the released version of the Pulumi Atlas provider. It is a valid
// semver string (pf/tfbridge.Main and pf/tfgen.Main both reject an empty or
// non-semver version). Override at release time with
// -ldflags "-X github.com/atlas/pulumi-atlas/provider.Version=x.y.z".
var Version = "0.1.0"

// atlasResource is a small helper that produces a canonical Pulumi resource
// token, e.g. ("OauthClient") -> "atlas:index/oauthClient:OauthClient".
func atlasResource(name string) *tfbridge.ResourceInfo {
	return &tfbridge.ResourceInfo{Tok: tfbridge.MakeResource(mainPkg, mainMod, name)}
}

// atlasDataSource produces a canonical Pulumi data-source (function) token,
// e.g. ("getOauthClient") -> "atlas:index/getOauthClient:getOauthClient".
func atlasDataSource(name string) *tfbridge.DataSourceInfo {
	return &tfbridge.DataSourceInfo{Tok: tfbridge.MakeDataSource(mainPkg, mainMod, name)}
}

// Provider returns the tfbridge ProviderInfo describing how the Atlas
// Terraform provider maps onto Pulumi. tfgen consumes it to emit the Pulumi
// Package Schema and the per-language SDKs; the resource plugin consumes it at
// runtime to serve the bridged provider.
func Provider() tfbridge.ProviderInfo {
	return tfbridge.ProviderInfo{
		// P wraps the plugin-framework provider. ShimProvider adapts a
		// terraform-plugin-framework provider into the bridge's shim.Provider.
		P: pf.ShimProvider(tfshim.NewProvider(Version)),

		Name:        "atlas",
		DisplayName: "Atlas",
		Publisher:   "atlas-authorization",
		Version:     Version,
		Description: "A Pulumi package for managing Atlas authentication-platform resources " +
			"(OAuth clients, SSO connections, resource servers, JWT templates, webhooks, " +
			"roles, organizations, custom domains, sign-in page design, messaging providers, " +
			"Sign in with Atlas apps and more) as code.",
		Keywords: []string{
			"pulumi", "atlas", "auth", "authentication", "identity",
			"category/network", "category/infrastructure",
		},
		License:    "Apache-2.0",
		Homepage:   "https://atlasauth.net",
		Repository: "https://github.com/Atlas-Authorization/pulumi-atlas",
		GitHubOrg:  "Atlas-Authorization",
		// The upstream provider lives in this monorepo (go.mod replace), not in
		// the module cache under the GitHub path tfgen would derive, so point the
		// docs lookup at it directly. Relative to the pulumi-atlas root, where
		// tfgen runs.
		UpstreamRepoPath: "../terraform-provider-atlas",
		// Runtime bridge metadata, written next to schema.json by tfgen and
		// required by the generator.
		MetadataInfo: tfbridge.NewProviderMetadata(nil),

		// --- Resource token mappings (TF type -> Pulumi token) ---
		Resources: map[string]*tfbridge.ResourceInfo{
			"atlas_oauth_client":          atlasResource("OauthClient"),
			"atlas_oauth_provider":        atlasResource("OauthProvider"),
			"atlas_sso_connection":        atlasResource("SsoConnection"),
			"atlas_resource_server":       atlasResource("ResourceServer"),
			"atlas_jwt_template":          atlasResource("JwtTemplate"),
			"atlas_webhook_endpoint":      atlasResource("WebhookEndpoint"),
			"atlas_role":                  atlasResource("Role"),
			"atlas_organization":          atlasResource("Organization"),
			"atlas_domain":                atlasResource("Domain"),
			"atlas_instance_config":       atlasResource("InstanceConfig"),
			"atlas_redirect_url":          atlasResource("RedirectUrl"),
			"atlas_billing_plan":          atlasResource("BillingPlan"),
			"atlas_organization_policy":   atlasResource("OrganizationPolicy"),
			"atlas_org_settings_schema":   atlasResource("OrgSettingsSchema"),
			"atlas_notification_template": atlasResource("NotificationTemplate"),
			"atlas_email_template":        atlasResource("EmailTemplate"),
			"atlas_notification_category": atlasResource("NotificationCategory"),
			"atlas_id_app":                atlasResource("IdApp"),
			"atlas_instance_design":       atlasResource("InstanceDesign"),
			"atlas_organization_design":   atlasResource("OrganizationDesign"),
			"atlas_cookie_consent":        atlasResource("CookieConsent"),
			"atlas_entitlement_override":  atlasResource("EntitlementOverride"),
			"atlas_email_provider":        atlasResource("EmailProvider"),
			"atlas_sms_provider":          atlasResource("SmsProvider"),
		},

		// --- Data-source token mappings (TF type -> Pulumi getter token) ---
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"atlas_oauth_client":       atlasDataSource("getOauthClient"),
			"atlas_oauth_provider":     atlasDataSource("getOauthProvider"),
			"atlas_resource_server":    atlasDataSource("getResourceServer"),
			"atlas_instance_preflight": atlasDataSource("getInstancePreflight"),
		},

		// --- Per-language SDK package paths ---
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@atlas-authorization/pulumi-atlas",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^18",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "atlas_authorization_pulumi_atlas",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath:                 "github.com/Atlas-Authorization/pulumi-atlas/sdk/go/atlas",
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "AtlasAuthorization",
			Namespaces: map[string]string{
				"atlas": "Atlas",
			},
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}
}
