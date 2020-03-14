# Database migrations

Setup database as target

```bash
# bash shell
$ export DATABASE_URL="postgresql://postgres:postgres@localhost/osuserver?sslmode=disable"
# fish shell
$ set -x DATABASE_URL "postgresql://postgres:postgres@localhost/osuserver?sslmode=disable"
```

and migrate

```bash
$ go get -u github.com/amacneil/dbmate
$ dbmate up
```