#!/bin/bash

#---------------
installGo() {
  mkdir -p "${GOPATH_PARENT}"

  if [ ! -d "${GOPATH_PARENT}" ]; then
    echo "Could not find directory to install go"
    exit 1
  fi

  if [ ! -f "${BIN_GO}" ]; then
    cd "${GOPATH_PARENT}" && ( \
      curl -L "https://go.dev/dl/go1.21.5.linux-amd64.tar.gz" | tar xzf -
      chmod u+x ./go/bin/go
      echo ${PWD}
    ) && cd -
  fi

  go version
}; export -f installGo

installTempl() {
  if [ ! -f "${BIN_TEMPL}" ]; then
    go install "github.com/a-h/templ/cmd/templ@${TEMPL_VERSION}"
  fi

  templ version
}; export -f installTempl

installGoTelemetry() {
  go run golang.org/x/telemetry/cmd/gotelemetry@latest on
}

installGopls() {
  go install golang.org/x/tools/gopls@latest
}

installSqlc() {
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
}

installGosqlite3() {
  go get github.com/mattn/go-sqlite3@latest
}

install() {
  installGo
  installTempl
  installGoTelemetry
  installGopls
  installSqlc
  installGosqlite3
}; export -f install

#---------------
templGenerate() {
  templ generate ./cmd ./internal ./pkg ./tests
}; export -f templGenerate

sqlcGenerate() {
  sqlc vet -f internal/sqlc/sqlc.yaml
  sqlc generate -f internal/sqlc/sqlc.yaml
}; export -f sqlcGenerate

goModTidy() {
  go mod tidy
}; export -f goModTidy

goModVendor() {
  go mod vendor
}; export -f goModVendor

#---------------
prerun() {
  ( \
    echo "[TEMPL]"; \
    templGenerate; \
  ) && ( \
    echo "[SQLC]"; \
    sqlcGenerate; \
  ) && ( \
    echo "[TIDY]"; \
    goModTidy; \
  ) && ( \
    echo "[VENDOR]"; \
    goModVendor; \
  )

  echo
}

clean() {
  rm -f ./internal/sqlc/{db,models,query.sql}.go
  rm -f ./internal/templ/*.go
  
  # `tail -n+2` excludes directory itself
  # `grep -v` used to exclude by regex
  find ./bin | tail -n+2 | \
    grep -v '^./bin/public.*$' | \
    xargs -i rm -rf {}
}

build() {
  prerun

  go build -o ../bin/iamfeelingcody cmd/iamfeelingcody/*.go && \
  chmod u+x ../bin/iamfeelingcody
}; export -f build

run() {
  prerun

  go run cmd/iamfeelingcody/*.go
}; export -f run

runbin() {
  build && \
  (cd ../bin && ./iamfeelingcody)
}

cleanrunbin() {
  clean && runbin
}

#---------------
source "$(dirname "$(realpath "${0}")")/env.sh"

# Not injection safe
case "${1}" in
  "")
    echo "Provide a target"
    ;;
  *)
    echo "${@}"
    eval "${@}"
    ;;
esac
