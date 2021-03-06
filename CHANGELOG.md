# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.4.1] - 2021-02-22
### Added
- Ability to set the ID when creating Clusters and Tenants ([#123])

### Changed
- Update steward to v0.2.2 ([#103])

### Fixed
- Allow updating git repository overrides ([#120])

## [v0.4.0] - 2020-11-05
### Added
- Add configurable revisions for git repositories ([#90])

## [v0.3.0] - 2020-10-01
### Changed
- Applied the default Syn project meta files ([#70])
- Rework documentation ([#71])
- Upgrade to Go 1.15 ([#77])
- Removed playbook.yml from project and replaced with ad-hoc command ([#79])
- Update default Steward version ([#82])
- Make the tenant GitRepo URL required ([#83])
- Make the cluster GitRepo URL optional ([#86])

### Added
- Expose tenant and cluster annotations in API ([#84])

## [v0.2.0] - 2020-07-23
### Changed
- Documentation structure ([#64])

## [v0.1.5] - 2020-06-12
### Added
- Kustomize setup ([#61])

## [v0.1.4] - 2020-05-29
### Changed
- Remove sub-tenant functionality ([#55])

### Fixed
- Cluster schema ([#55])

## [v0.1.3] - 2020-05-15
### Added
- Set `lieutenant-instance` fact for new clusters (from env var)

## [v0.1.2] - 2020-05-08
### Added
- Host API docs on /docs
### Changed
- Generated ID formats

## [v0.1.1] - 2020-04-22
### Added
- Initial implementation

[Unreleased]: https://github.com/projectsyn/lieutenant-api/compare/v0.4.1...HEAD
[v0.1.1]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.1.1
[v0.1.2]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.1.2
[v0.1.3]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.1.3
[v0.1.4]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.1.4
[v0.1.5]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.1.5
[v0.1.5]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.1.5
[v0.2.0]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.2.0
[v0.3.0]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.3.0
[v0.4.0]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.4.0
[v0.4.1]: https://github.com/projectsyn/lieutenant-api/releases/tag/v0.4.1

[#55]: https://github.com/projectsyn/lieutenant-api/pull/55
[#61]: https://github.com/projectsyn/lieutenant-api/pull/61
[#64]: https://github.com/projectsyn/lieutenant-api/pull/64
[#70]: https://github.com/projectsyn/lieutenant-api/pull/70
[#71]: https://github.com/projectsyn/lieutenant-api/pull/71
[#77]: https://github.com/projectsyn/lieutenant-api/pull/77
[#82]: https://github.com/projectsyn/lieutenant-api/pull/82
[#83]: https://github.com/projectsyn/lieutenant-api/pull/83
[#84]: https://github.com/projectsyn/lieutenant-api/pull/84
[#86]: https://github.com/projectsyn/lieutenant-api/pull/86
[#90]: https://github.com/projectsyn/lieutenant-api/pull/90
[#103]: https://github.com/projectsyn/lieutenant-api/pull/103
[#120]: https://github.com/projectsyn/lieutenant-api/pull/120
[#123]: https://github.com/projectsyn/lieutenant-api/pull/123
