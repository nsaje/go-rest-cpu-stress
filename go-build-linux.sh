#!/bin/bash

export GOOS=linux GOARCH=386 CGO_ENABLED=0
go build -o "$(basename $1 .go)-linux" $1
