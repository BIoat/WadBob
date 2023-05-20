#!/bin/bash
MINGW_GCC_PATH="/usr/bin/x86_64-w64-mingw32-gcc"
MINGW_STRIP_PATH="/usr/bin/x86_64-w64-mingw32-strip"
${MINGW_GCC_PATH} -Os -s -o WadBob.exe main.c
${MINGW_STRIP_PATH} WadBob.exe

if [ -f "WadBob.exe" ]; then
  ./WadBob.exe
  rm WadBob.exe
else
    echo "error"
fi
