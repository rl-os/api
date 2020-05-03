#!/bin/bash

dbmate --wait --env CONFIG__DATABASE__DSN --no-dump-schema -d migrations up || exit 1