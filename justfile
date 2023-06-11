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

dev:
  @start http://localhost:3000
  @yarn dev

start:
  @yarn start
