#!/bin/bash
go build -ldflags="-X 'github.com/kintohub/kinto-cli-go/internal/config.Version=$1'"



