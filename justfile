set shell := ["powershell.exe", "-c"]

readme:
  @just -l

install:
  @yarn

build:
  if (Test-Path -Path dist) { rmdir dist }
  @yarn tsc
  @mkdir dist/public
  @cp -r public/* dist/public
  @mkdir dist/views
  @cp -r views/* dist/views 

dev:
  @start http://localhost:3000
  @yarn dev

start:
  @yarn start

format-check:
  @yarn format:check

format:
  @yarn format
