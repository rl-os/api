<a name="unreleased"></a>
## [Unreleased]


<a name="v1.1.21"></a>
## [v1.1.21] - 2020-05-03
### Bug Fixes
- **ayako:** fmt
- **deploy:** ayako path prefix + bump container version
- **deploy:** nulls in beatmapset + covers list
- **deploy:** nulls in beatmapset + covers list

### Chore
- remove useless systemd
- **deps:** bump github.com/lib/pq from 1.3.0 to 1.4.0 in /api ([#49](https://github.com/deissh/osu-api-server/issues/49))
- **deps:** bump github.com/deissh/go-utils from 1.0.0 to 1.1.0 in /api ([#47](https://github.com/deissh/osu-api-server/issues/47))
- **deps:** bump github.com/deissh/go-utils in /ayako ([#46](https://github.com/deissh/osu-api-server/issues/46))
- **deps:** bump github.com/lib/pq from 1.3.0 to 1.4.0 in /ayako ([#45](https://github.com/deissh/osu-api-server/issues/45))

### Ci
- disable codecov

### Code Refactoring
- **ayako:** change store.BeatmapSet.ComputeFields
- **ayako:** rename store methods
- **ayako:** main.go
- **ayako:** store logging (decorator)

### Features
- **ayako:** add docker-init.sh and auto migrations ([#51](https://github.com/deissh/osu-api-server/issues/51))
- **ayako:** simple search endpoint, beatmap/beatmapset lockup
- **ayako:** getting beatmaps + some fixs
- **ayako:** create beatmapset with beatmaps

### Test
- **ayako:** fix errors


[Unreleased]: https://github.com/deissh/osu-api-server/compare/v1.1.21...HEAD
[v1.1.21]: https://github.com/deissh/osu-api-server/compare/v1.1.14a...v1.1.21
