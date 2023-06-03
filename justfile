set shell := ["powershell.exe", "-c"]

readme:
  @echo "This project created to implemented with mdoctor."
  @echo "Writen in typescript. Please use 'just build' command to build it."
  @echo "You can use 'just -l' for more information."

install:
  @yarn

build: install
  if (Test-Path -Path dist) { rmdir dist }
  @yarn tsc
  mkdir dist/public
  cp -r public/* dist/public

dev:
  @start http://localhost:3000
  @yarn dev

run:
  node dist/index.js
