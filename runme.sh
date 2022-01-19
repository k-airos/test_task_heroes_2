#!/bin/sh
if [ `whoami` != root ]; then
    echo Please run this script as root or using sudo
    exit
fi
service mongod stop
systemctl stop mongo
docker run --name testmongodb -d -p 27017:27017 mongo
docker run --name KairosDocker -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
cd marvel
go build ./cmd/api/main.go
chmod +x ./main
cd ../DC
go build ./cmd/api/main.go
chmod +x ./main


