templ:
	templ generate .

build: templ
	go build -o ./bin/iamfeelingcody cmd/iamfeelingcody/main.go

run: templ
	go run cmd/iamfeelingcody/main.go
