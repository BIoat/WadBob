#!/bin/bash
# Linux
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o out/bin


# Windows
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o out/bin.exe

# Finish
cd out && strip bin bin.exe && upx --lzma bin bin.exe && sync
# lzma -qv9T 8 bin bin.exe


cp bin.exe /home/anon/wadbot/
# fyne-cross windows -arch=amd64
