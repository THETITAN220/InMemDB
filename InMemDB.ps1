#Powershell script to run on Windows Platforms

$ErrorActionPreference = "Stop"

$ScriptDir = $PSScriptRoot
if (-not $ScriptDir) { $ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition }

Set-Location -Path $ScriptDir

if (-not (Test-Path -Path ".\build" -PathType Container)) {
    Write-Host "Creating build directory..."
    New-Item -Path ".\build" -ItemType Directory | Out-Null
}

go build -o .\build\InMemDB.exe ./app/main.go
Write-Host "Build complete: .\build\InMemDB"

$ExecutablePath = ".\build\InMemDB.exe"

if (-not (Test-Path -Path $ExecutablePath -PathType Leaf)) {
    throw "Error: Could not find executable at $ExecutablePath. Build failed."
}

& $ExecutablePath $args
