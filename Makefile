GO_VERSION=1.21.5
TEMPL_VERSION=v0.2.501
GO=go${GO_VERSION}

# Depends on there being a go installation present.
# Install with package manager or follow official docs:
# https://go.dev/doc/install
goinstall:
	go install golang.org/dl/go${GO_VERSION}@latest
	go${GO_VERSION} download

templinstall:
	$(GO) install github.com/a-h/templ/cmd/templ@${TEMPL_VERSION}

install: goinstall templinstall

#-----------
templ:
	templ generate .

tidy:
	$(GO) mod tidy

#-----------
build: templ tidy
	$(GO) build -o ./bin/iamfeelingcody cmd/iamfeelingcody/*.go

run: templ tidy
	$(GO) run cmd/iamfeelingcody/*.go
