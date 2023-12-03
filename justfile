readme:
  @just -l

tidy:
  @go mod tidy

run:
  @go run main.go

build:
  @go build main.go

clean:
  if [ -e main ]; then rm main; fi
