# Setting up the RisuLife

## Requirements

Before you get started, make sure you have the following dependencies installed on your machine:

* GNU Make >= 4.2.1
* Golang >= 1.13
* NodeJS >= 8 with yarn or npm 3.
* Python 3.7+ (using async/await)
* Redis
* PostgreSQL 11+

## Local Installation

```bash
$ git clone git@github.com:deissh/rl.git && cd rl
```

First you need to prepare the database, you can create and migrate database using the `dbmate` command with the `db` folder.

```bash
# install dbmate tool
$ go get -u github.com/amacneil/dbmate

# setup database uri
$ export DATABASE_URL=postgresql://USER:PASS@localhost/DATABASE?sslmode=disable

# and migrate database
$ cd db
$ dbmate -d migrations -s schema.sql up
```

#### API

All commands are executed in the `api` directory.
```bash
$ cd api
$ make install
```

For the server to work, you need to fill out the config.yaml file or setup env: `SERVER__DATABASE__DSN`, `SERVER__REDIS__HOST` and etc.

```bash
$ vim config.yaml
```

To run a specific command you need to run.

```bash
$ make build
$ ./bin/www # or ./bin/cron or ./bin/statsd
```

Example GoLand configuration. For other cmds, the configuration will be similar

![www](https://i.imgur.com/IfTLma7.png)

#### Ayako

Ayako tries to hold a replica of the osu! database as updated as possible.

All commands are executed in the `ayako` directory.
```bash
$ cd ayako
$ make install
```

For the server to work, you need to fill out the config.yaml file or setup env: `SERVER__DATABASE__DSN`, `SERVER__REDIS__HOST` and etc.

```bash
$ vim config.yaml
```

To run a specific command you need to run.

```bash
$ make build
$ ./bin/ayako
```

#### Frontend

> todo
