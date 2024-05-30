@echo off
@if errorlevel 1 pause

SET GOARCH=amd64
SET GOOS=windows
go build -o build/m3u8_windows.exe ./m3u8.go


set GOARCH=amd64
set GOOS=linux
go build -o build/m3u8_linux ./m3u8.go
