templ:
	templ generate .

tidy:
	go mod tidy

build: templ tidy
	go build -o ./bin/iamfeelingcody cmd/iamfeelingcody/*.go

run: templ tidy
	go run cmd/iamfeelingcody/*.go
