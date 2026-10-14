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

> The SDKs are generated (see "Generating the SDKs" below) and published to the
> usual registries as part of a release. Once published:

**Node.js / TypeScript**

```bash
npm install @atlas-authorization/pulumi-atlas
```

**Python**

```bash
pip install atlas_authorization_pulumi_atlas
```

**Go**

```bash
go get github.com/Atlas-Authorization/pulumi-atlas/sdk/go/atlas
```

**.NET**

```bash
dotnet add package AtlasAuthorization.Atlas
```

## Example

TypeScript — create an OAuth client:

```typescript
import * as atlas from "@atlas-authorization/pulumi-atlas";

const web = new atlas.OauthClient("web", {
    name: "Web App",
    redirectUris: ["https://app.example.com/callback"],
    grantTypes: ["authorization_code", "refresh_token"],
});

export const clientId = web.clientId;
```

Python:

```python
import atlas_authorization_pulumi_atlas as atlas

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

`terraform-provider-atlas` is not yet published as a tagged Go module, so
`provider/go.mod` pulls it in via a `replace` directive pointing at the sibling
`../../terraform-provider-atlas` directory. The TF provider keeps its
implementation in `internal/`, which Go forbids other modules from importing;
this bridge therefore imports the small exported shim at
`terraform-provider-atlas/shim` (`shim.NewProvider`), which is the only public
entry point needed. **For a real release, publish `terraform-provider-atlas` as
a tagged Go module (or vendor it) and repoint/remove the `replace`.**

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

The generated sources land under `sdk/<lang>/`. Package them with the native
language tooling (`npm`, `pip`/`build`, `go`, `dotnet`) before publishing.

## Cutting a Pulumi-registry release

1. Publish `terraform-provider-atlas` as a tagged Go module (or vendor it) and
   update `provider/go.mod`'s `replace`/`require` accordingly.
2. `make schema` — regenerate and commit `schema.json`.
3. `make build_sdks` — generate the Node/Python/Go/.NET SDK sources.
4. Build release binaries of `pulumi-resource-atlas` for each target platform
   (set the version via `-ldflags "-X .../provider.Version=X.Y.Z"`), package the
   SDKs, and publish:
   - npm: `@atlas-authorization/pulumi-atlas`
   - PyPI: `atlas_authorization_pulumi_atlas`
   - Go: tag `sdk/go/...`
   - NuGet: `AtlasAuthorization.Atlas`
5. Submit the provider to the
   [Pulumi Registry](https://www.pulumi.com/registry/) (publisher
   `Atlas-Authorization`) per
   <https://www.pulumi.com/docs/using-pulumi/pulumi-packages/publishing/>.

## License

Apache-2.0.
