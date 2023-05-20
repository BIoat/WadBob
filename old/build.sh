#!/bin/bash
#Android 
fyne install -os android -appID com.wadbot.wadbob -icon Icon.png

# Windows
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-s -w -H=windowsgui" -o out/bin.exe && sync && cp out/bin.exe /home/anon/wadbot/

# Linux
GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o out/bin



# Finish
cd out && strip bin bin.exe && upx --lzma bin bin.exe && sync

# lzma -qv9T 8 bin bin.exe
# fyne-cross windows -arch=amd64
