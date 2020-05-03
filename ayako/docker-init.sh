#!/bin/bash

dbmate --wait --env CONFIG__DATABASE__DSN --no-dump-schema -d migrations -s schema.sql up || exit 1