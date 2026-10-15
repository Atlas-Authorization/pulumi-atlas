# Changelog

All notable changes to the Pulumi Atlas provider are documented here. The
format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.8.0]

Versioned in step with `terraform-provider-atlas` from this release on.

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
