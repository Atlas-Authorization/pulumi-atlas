# Changelog

All notable changes to the Pulumi Atlas provider are documented here. The
format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

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

[Unreleased]: https://github.com/Atlas-Authorization/pulumi-atlas
