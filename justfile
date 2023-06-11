set shell := ["powershell.exe", "-c"]
alias b :=  build
alias ca :=  clean
alias d := dev
alias f := format

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

clean-dist:
  @echo "> Cleaning 'dist' if it's exists"
  if (Test-Path -Path dist) { rmdir dist }
  @echo " "

clean-deps:
  @echo "> Cleaning 'node_modules' if it's exists"
  if (Test-Path -Path node_modules) { rmdir node_modules }
  @echo " "

clean: clean-dist clean-deps
