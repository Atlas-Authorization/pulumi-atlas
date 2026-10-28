# Changelog

All notable changes to the Pulumi Atlas provider are documented here. The
format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Every resource and data source of the Terraform provider is now bridged (123 resources, 34 functions in the schema), including the 83 typed resources generated from the Backend API for the re-audit tracks and the remaining collections: MCP gateways, connection actions and triggers, agent lifecycle, conditional access and device assurance, workflows, chat integrations, IGA, administrative units, FGA assertions and materializations, coupons, audit retention, customer-managed keys, SCIM targets, WS-Fed connections, groups, machines, actions, flows, feature flags, network zones and more. Resources without an explicit token get the standard one (`atlas_mcp_gateway` becomes `atlas:index/mcpGateway:McpGateway`, its data source `getMcpGateway`).
- New properties on existing resources: `OauthClient.oidcMetadata`, `BillingPlan.graceDays`, `OrganizationPolicy.ssoBypass` and `itContacts`, `SsoConnection.requireConsentScreen`.

### Changed

- The `integration` attribute of `atlas.Integration` is named `integrationId` in the SDKs (a .NET member cannot share its class's name).

## [0.8.0]

Versioned in step with `terraform-provider-atlas` from this release on. First
release published to the package registries.

### Changed

- SDK package names: Node.js `@atlasauth/pulumi` (was
  `@atlas-authorization/pulumi-atlas`), Python `pulumi-atlasauth`, imported as
  `pulumi_atlasauth` (was `atlas_authorization_pulumi_atlas`). Go
  (`github.com/Atlas-Authorization/pulumi-atlas/sdk/go/atlas`) and .NET
  (`AtlasAuthorization.Atlas`) are unchanged.
- The schema sets `pluginDownloadURL` to
  `github://api.github.com/Atlas-Authorization`, so the SDKs install the
  `pulumi-resource-atlas` plugin from this repository's GitHub releases.

### Added

- Inputs and outputs from `terraform-provider-atlas` 0.8.0:
  - `OauthClient.skipLogoutConfirmation` (boolean, default `false`;
    first-party clients only, so a preview fails when the client is
    third-party) and `getOauthClient` output `skipLogoutConfirmation`.
  - `CookieConsent.settingsLink` (`CookieConsentSettingsLink`: `label`,
    `preset`, `attributes`, and output `renderedAttributes`) — a "Cookie
    settings" link in the hosted pages' footer.
  - `EmailProvider.forceDestroy` (boolean, default `false`) — delete removes
    the provider even while email-dependent features are on; without it,
    Atlas's `409 EMAIL_PROVIDER_IN_USE` now says how to proceed.
- Mappings for the resources and data source added in
  `terraform-provider-atlas` 0.7.0: `IdApp` (Sign in with Atlas apps — set the
  provider's `platformKey`), `InstanceDesign`, `OrganizationDesign`,
  `CookieConsent`, `EntitlementOverride`, `EmailProvider`, `SmsProvider`, and
  the `getInstancePreflight` function. The new `platformKey` provider
  configuration and the extended `WebhookEndpoint`, `OauthProvider` and
  `OrganizationPolicy` inputs come through the bridge automatically.
- `NotificationCategory` resource mapping (17 bridged resources before this
  release).

### Initial release

- Initial Pulumi provider for Atlas, bridged from `terraform-provider-atlas`
  via `pulumi-terraform-bridge` v3 (plugin-framework path).
- 16 bridged resources: `OauthClient`, `OauthProvider`, `SsoConnection`,
  `ResourceServer`, `JwtTemplate`, `WebhookEndpoint`, `Role`, `Organization`,
  `Domain`, `InstanceConfig`, `RedirectUrl`, `BillingPlan`,
  `OrganizationPolicy`, `OrgSettingsSchema`, `NotificationTemplate`,
  `EmailTemplate`.
- 3 bridged data sources: `getOauthClient`, `getOauthProvider`,
  `getResourceServer`.
- Pulumi Package Schema generation (`pulumi-tfgen-atlas`) and the runtime
  resource plugin (`pulumi-resource-atlas`).
- Per-language SDK package mappings for Node.js, Python, Go and .NET.
- `make` targets: `tfgen`, `schema`, `provider`, `build_sdks`.

[0.8.0]: https://github.com/Atlas-Authorization/pulumi-atlas
