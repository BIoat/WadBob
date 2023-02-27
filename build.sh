#!/bin/bash
# Linux
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin

# Windows
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin.exe

# Finish
strip bin bin.exe && mv bin bin.exe out/
cp out/bin.exe /home/anon/wadbot/

# fyne-cross windows -arch=amd64
