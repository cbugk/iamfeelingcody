#!/bin/bash

#---------------
installGo() {
  mkdir -p "${GOPATH}"

  if [ ! -f "${BIN_GO}" ]; then
    cd "${GOPATH}/.." && ( \
      curl -L "https://go.dev/dl/go1.21.5.linux-amd64.tar.gz" | tar xzf -
      chmod u+x ./go/bin/go
    ) && cd -
  fi

  go version
}

installTempl() {
  if [ ! -f "${BIN_TEMPL}" ]; then
    go install "github.com/a-h/templ/cmd/templ@${TEMPL_VERSION}"
  fi

  templ version
}

install() {
  installGo
  installTempl
}

#---------------
templGenerate() {
  templ generate ./cmd ./internal ./pkg ./tests
}

goModTidy() {
  go mod tidy
}

goModVendor() {
  go mod vendor
}

#---------------
build() {
  templGenerate
  goModTidy
  goModVendor

  go build -o ./bin/iamfeelingcody cmd/iamfeelingcody/*.go
}

run() {
  templGenerate
  goModTidy
  goModVendor

  build && ./bin/iamfeelingcody
}

go() {
  "${BIN_GO}" "${@}"
}; export -f go

templ() {
  "${BIN_TEMPL}" "${@}"
}; export templ

main() {
  DIR_SCRIPT="$(dirname -- "$(readlink -f -- "${1}")")"

  export GOROOT="${DIR_SCRIPT}/bin/go"
  export GOPATH="${DIR_SCRIPT}/bin/go"
  export GO_VERSION="1.21.5"
  export BIN_GO="${GOPATH}/bin/go"
  export BIN_TEMPL="${GOPATH}/bin/templ"
  export TEMPL_VERSION="v0.2.501"

  #---------------
  # Not injection safe
  case "${2}" in
    "")
      echo "Provide a target"
      ;;
    *)
      echo "${2}"
      eval "${2}"
      ;;
  esac
}

if [ "${0}" = 'source' ] || [ ${0} = '.' ];then
  echo Script must not be sourced
else
  main "${0}" "${1}"
fi
