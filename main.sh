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

install() {
  installGo
  installTempl
}; export -f install

#---------------
templGenerate() {
  templ generate ./cmd ./internal ./pkg ./tests
}; export -f templGenerate

goModTidy() {
  go mod tidy
}; export -f goModTidy

goModVendor() {
  go mod vendor
}; export -f goModVendor

#---------------
prerun() {
  templGenerate
  goModTidy
  goModVendor
}

build() {
  prerun

  go build -o ./bin/iamfeelingcody cmd/iamfeelingcody/*.go && \
  chmod u+x ./bin/iamfeelingcody
}; export -f build

run() {
  prerun

  go run cmd/iamfeelingcody/*.go
}; export -f run


#---------------
#if [ "${0}" = 'source' ] || [ ${0} = '.' ];then
#  echo Script must not be sourced
#  exit 1
#fi

# Pop goroot parameter
export GOPATH_PARENT="${IAMFEELINGCODY_GOPATH_PARENT:=${HOME}}"
echo "iamfeelingcody: ${IAMFEELINGCODY_GOPATH_PARENT}"
export GOPATH="${GOPATH_PARENT}/go"

export GO_VERSION="1.21.5"
export TEMPL_VERSION="v0.2.501"

export BIN_GO="${GOPATH}/bin/go"
export BIN_TEMPL="${GOPATH}/bin/templ"

export PATH="${GOPATH}/bin:${PATH}"

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
