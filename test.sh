#!/bin/bash
if test "$#" -ne 1; then
  echo "Please set the path where test files are located"
  exit 0
fi

path=$1
go test $path -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
