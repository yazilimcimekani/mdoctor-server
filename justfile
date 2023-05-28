set shell := ["powershell.exe", "-c"]

build:
  @echo "Removing existing dist folder"
  @echo "You can press return to accept"
  @rmdir dist
  @yarn tsc
  mkdir dist/public
  cp -r public/* dist/public

run:
  node dist/index.js
