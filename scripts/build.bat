@echo off
setlocal enabledelayedexpansion

if "%1"=="" (
    echo Usage: %0 [api/grpc]
    exit /b 1
)

cd ..
set param=%1
set server_exe=bin\server_%param%.exe
set client_exe=bin\client_%param%.exe

if not exist bin (
    echo Creating bin directory...
    mkdir bin
)

echo Compiling server...
go build -o "%server_exe%" src/%param%/server.go
echo Server compiled!

echo Compiling client...
go build -o "%client_exe%" src/%param%/client.go
echo Client compiled!

cd scripts
endlocal