#!/usr/bin/env sh

set -e

docker run -v /home/anduser/Documents/DataProject/dataCenter/migrations:/migrations --network host migrate/migrate create -dir migrations -ext sql -seq create_regions_schema

docker run -v /home/anduser/Documents/DataProject/dataCenter/migrations:/migrations --network host migrate/migrate create -dir migrations -ext sql -seq create_regions_table

docker run -v /home/anduser/Documents/DataProject/dataCenter/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://Pasha:12345678@localhost:54320/mydb?sslmode=disable up