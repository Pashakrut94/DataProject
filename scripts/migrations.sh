#!/usr/bin/env sh

set -e

docker run -v /home/anduser/Documents/DataProject/dataCenter/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://Pasha:12345678@localhost:54320/mydb?sslmode=disable up