#! /bin/bash
echo "Run Scheduler"
rm *.sst *.vlog MANIFEST
go run tasksch/tasksch.go

