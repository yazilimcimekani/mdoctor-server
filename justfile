set shell := ["powershell.exe", "-c"]

readme:
  @echo "This project created to implemented with MDoctor."
  @echo "Writen in typescript. Please use 'just build' command to build it."
  @echo "Then use 'just start' command to start compiled server."
  @echo " "
  @echo "You can use 'just -l' for more information."

install:
  @yarn

build: install
  if (Test-Path -Path dist) { rmdir dist }
  @yarn tsc
  @mkdir dist/public
  @cp -r public/* dist/public

dev:
  @start http://localhost:3000
  @yarn dev

start:
  @yarn start
