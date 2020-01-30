<a name="unreleased"></a>
## [Unreleased]


<a name="vv0.1.0"></a>
## [vv0.1.0] - 2020-01-30
### Feat
- add PM and loading private channels
- add user relations and getting friends
- add chat updates and getting message history
- sending message and fix users in channels
- add show users in channel
- add CreateMessageTable migration
- add git chglog generator

### Fix
- README.md logo
- message channel id and leaving from channels
- test running

### Refactor
- change entity struct

### Pull Requests
- Merge pull request [#23](https://github.com/deissh/osu-api-server/issues/23) from deissh/next


<a name="v0.0.1"></a>
## v0.0.1 - 2020-01-28
### Chore
- remove useless files and change struct

### Docs
- update docsify configuration
- init

### Feat
- add supporter expired at field
- rewrite auth middleware
- add nginx configurations
- add default avatars and backgrounds WARN: breaking changes
- add sentry support
- controllers and db migration files
- add more logging when startup
- fix user online status
- update db scheme
- add countries table with default values
- rewrite to CustomError
- auth request and check access_token
- add genesis user and oauth_client
- user login
- add client_id in oauth_token and change services struct
- add migration tool + dump
- custom error structure with nessosory information
- new error handler
- add oauth_client service and rework oauth_token
- add basic service
- add migrations
- add oauth routes
- add cors configuration + log format
- add logging
- rewrite from gin to echo (routing problem)
- add v2 routing
- add dockerfile
- add v1 routes
- add redis and rewrite cyclic imports
- add redis and rewrite cyclic imports
- add example config file
- database connecting
- add v2
- add logger + simple gin app

### Fix
- getting current user
- setting online status and verify user token
- getting profile
- code style
- entity datatypes
- error response style
- oauth handler
- crlf to lf

### Refacor
- rewrite services and move out entity

### Refactor
- using config

### Pull Requests
- Merge pull request [#22](https://github.com/deissh/osu-api-server/issues/22) from deissh/fix/user-profile
- Merge pull request [#21](https://github.com/deissh/osu-api-server/issues/21) from deissh/cleaning-1
- Merge pull request [#20](https://github.com/deissh/osu-api-server/issues/20) from deissh/deissh-patch-1
- Merge pull request [#19](https://github.com/deissh/osu-api-server/issues/19) from deissh/feature/issue-templates
- Merge pull request [#18](https://github.com/deissh/osu-api-server/issues/18) from deissh/feature/ci


[Unreleased]: https://github.com/deissh/osu-api-server/compare/vv0.1.0...HEAD
[vv0.1.0]: https://github.com/deissh/osu-api-server/compare/v0.0.1...vv0.1.0
