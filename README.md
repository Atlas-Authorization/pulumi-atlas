# Pulumi Atlas Provider

A [Pulumi](https://www.pulumi.com) provider for [Atlas](https://atlasauth.net),
the Clerk-class authentication platform. Manage your Atlas instance
configuration — OAuth clients, SSO connections, resource servers, JWT
templates, webhook endpoints, roles, organizations, custom domains, billing
plans and more — as code, from TypeScript/JavaScript, Python, Go or .NET.

This provider is a **bridge** over the
[`terraform-provider-atlas`](../terraform-provider-atlas) Terraform provider,
built with [`pulumi-terraform-bridge`](https://github.com/pulumi/pulumi-terraform-bridge).
The resource and data-source schemas are reused verbatim, so the Pulumi and
Terraform providers always expose the same surface and talk to the same Atlas
Backend API.

## Configuration

The provider authenticates to the Atlas Backend API with an instance secret
key. Set it (and, optionally, a non-default API URL) via environment variables
so the key never lands in state or source:

```bash
export ATLAS_SECRET_KEY="sk_live_..."
export ATLAS_API_URL="https://api.atlasauth.net"   # optional; this is the default
```

They can also be set as provider config (`secretKey`, `apiUrl`), but the
environment variables are preferred.

## Install

**Node.js / TypeScript**

```bash
npm install @atlasauth/pulumi
```

**Python**

```bash
pip install pulumi-atlasauth
```

**Go**

```bash
go get github.com/Atlas-Authorization/pulumi-atlas/sdk/go/atlas
```

**.NET**

```bash
dotnet add package AtlasAuthorization.Atlas
```

The SDKs install the matching `pulumi-resource-atlas` plugin automatically from
this repository's GitHub releases. To install it by hand:

```bash
pulumi plugin install resource atlas v0.8.0 --server github://api.github.com/Atlas-Authorization
```

## Example

TypeScript — create an OAuth client:

```typescript
import * as atlas from "@atlasauth/pulumi";

const web = new atlas.OauthClient("web", {
    name: "Web App",
    redirectUris: ["https://app.example.com/callback"],
    grantTypes: ["authorization_code", "refresh_token"],
});

export const clientId = web.clientId;
```

Python:

```python
import pulumi_atlasauth as atlas

web = atlas.OauthClient("web",
    name="Web App",
    redirect_uris=["https://app.example.com/callback"],
    grant_types=["authorization_code", "refresh_token"])
```

Go:

```go
import atlas "github.com/Atlas-Authorization/pulumi-atlas/sdk/go/atlas"

web, err := atlas.NewOauthClient(ctx, "web", &atlas.OauthClientArgs{
    Name:         pulumi.String("Web App"),
    RedirectUris: pulumi.StringArray{pulumi.String("https://app.example.com/callback")},
})
```

## Resources & data sources

24 resources and 4 data sources are bridged. Pulumi tokens follow the canonical
`atlas:index/<camelName>:<PascalName>` form.

| Terraform type                | Pulumi resource                   |
| ----------------------------- | --------------------------------- |
| `atlas_oauth_client`          | `atlas:index:OauthClient`         |
| `atlas_oauth_provider`        | `atlas:index:OauthProvider`       |
| `atlas_sso_connection`        | `atlas:index:SsoConnection`       |
| `atlas_resource_server`       | `atlas:index:ResourceServer`      |
| `atlas_jwt_template`          | `atlas:index:JwtTemplate`         |
| `atlas_webhook_endpoint`      | `atlas:index:WebhookEndpoint`     |
| `atlas_role`                  | `atlas:index:Role`                |
| `atlas_organization`          | `atlas:index:Organization`        |
| `atlas_domain`                | `atlas:index:Domain`              |
| `atlas_instance_config`       | `atlas:index:InstanceConfig`      |
| `atlas_redirect_url`          | `atlas:index:RedirectUrl`         |
| `atlas_billing_plan`          | `atlas:index:BillingPlan`         |
| `atlas_organization_policy`   | `atlas:index:OrganizationPolicy`  |
| `atlas_org_settings_schema`   | `atlas:index:OrgSettingsSchema`   |
| `atlas_notification_template` | `atlas:index:NotificationTemplate`|
| `atlas_email_template`        | `atlas:index:EmailTemplate`       |
| `atlas_notification_category` | `atlas:index:NotificationCategory`|
| `atlas_id_app`                | `atlas:index:IdApp`               |
| `atlas_instance_design`       | `atlas:index:InstanceDesign`      |
| `atlas_organization_design`   | `atlas:index:OrganizationDesign`  |
| `atlas_cookie_consent`        | `atlas:index:CookieConsent`       |
| `atlas_entitlement_override`  | `atlas:index:EntitlementOverride` |
| `atlas_email_provider`        | `atlas:index:EmailProvider`       |
| `atlas_sms_provider`          | `atlas:index:SmsProvider`         |

Data sources: `atlas_oauth_client` → `getOauthClient`, `atlas_oauth_provider` →
`getOauthProvider`, `atlas_resource_server` → `getResourceServer`,
`atlas_instance_preflight` → `getInstancePreflight`.

`IdApp` (Sign in with Atlas apps) needs the provider's `platformKey`
(`psk_...`, or `ATLAS_PLATFORM_KEY`).

## Building from source

The provider plugin and its code generator are pure Go and need no external
tools:

```bash
make tfgen      # build the pulumi-tfgen-atlas codegen binary
make schema     # regenerate provider/cmd/pulumi-resource-atlas/schema.json
make provider   # build the pulumi-resource-atlas plugin
```

Or directly:

```bash
cd provider && go build ./...
```

### How the TF provider is referenced

`provider/go.mod` pulls in `terraform-provider-atlas` via a `replace` directive
pointing at a sibling `../../terraform-provider-atlas` directory. To build from
a clone of this repository, check out
[`Atlas-Authorization/terraform-provider-atlas`](https://github.com/Atlas-Authorization/terraform-provider-atlas)
next to it at the same version tag (the two providers are versioned in step):

```bash
git clone https://github.com/Atlas-Authorization/pulumi-atlas
git clone --branch v0.8.0 https://github.com/Atlas-Authorization/terraform-provider-atlas
cd pulumi-atlas && make provider
```

The TF provider keeps its implementation in `internal/`, which Go forbids other
modules from importing; this bridge therefore imports the small exported shim at
`terraform-provider-atlas/shim` (`shim.NewProvider`), which is the only public
entry point needed.

## Generating the SDKs

SDK generation needs the [Pulumi CLI](https://www.pulumi.com/docs/install/) on
`PATH` (it provides the per-language codegen) in addition to the Go toolchain:

```bash
make build_sdks          # all languages
make build_nodejs        # or one at a time: nodejs | python | go | dotnet
```

`make schema` also restores the schema's top-level `"version"` (tfgen drops it)
from `VERSION` via `scripts/set-schema-version.mjs`, so it needs `node` on `PATH`.
The `Pulumi SDKs` GitHub Actions workflow runs the same targets on Linux with the
HCL examples converted into the docs, and uploads the result as an artifact.

On Windows, `pulumi-tfgen-atlas <lang>` cannot create the SDK's nested
folders; generate from the committed schema with the Pulumi CLI instead:
`pulumi package gen-sdk provider/cmd/pulumi-resource-atlas/schema.json --language <lang> --out sdk`.

The generated sources land under `sdk/<lang>/`. Package them with the native
language tooling (`npm`, `pip`/`build`, `go`, `dotnet`) before publishing.

## Cutting a release

Releases are built by GitHub Actions (`.github/workflows/release.yml`). Bump
`VERSION` in the `Makefile` and `Version` in `provider/resources.go`, add a
`CHANGELOG.md` entry, make sure `terraform-provider-atlas` has the same tag, then
push a `vX.Y.Z` tag. The workflow:

1. regenerates `schema.json` and the Node.js, Python, Go and .NET SDKs;
2. builds `pulumi-resource-atlas` for linux/darwin/windows on amd64/arm64 and
   attaches `pulumi-resource-atlas-vX.Y.Z-<os>-<arch>.tar.gz` to the GitHub
   release;
3. publishes `@atlasauth/pulumi` to npm, `pulumi-atlasauth` to PyPI and
   `AtlasAuthorization.Atlas` to NuGet (skipped when `NUGET_API_KEY` is not
   set), and tags the Go SDK as `sdk/vX.Y.Z`.

Every publish step is idempotent, so re-running a failed release only does what
is still missing. `.github/workflows/build.yml` builds the provider and compiles
every SDK on each pull request and push to `main`.

## License

Apache-2.0.
