#!/usr/bin/env bash

psql postgres --command "CREATE USER test_user WITH SUPERUSER PASSWORD 'test_user_password';"
createdb -O test_user promotioner_test_database