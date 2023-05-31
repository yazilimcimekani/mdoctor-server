set shell := ["powershell.exe", "-c"]

readme:
  @echo "This project created to implemented with mdoctor."
  @echo "Writen in typescript. Please use 'just build' command to build it."
  @echo "You can use 'just -l' for more information."

build:
  if (Test-Path -Path dist) {
  @echo "Removing existing dist folder"
  @echo "You can press return to accept"
  rmdir dist 
  }
  @yarn tsc
  mkdir dist/public
  cp -r public/* dist/public

dev:
  @start http://localhost:3000
  @yarn dev

run:
  node dist/index.js
