#!/bin/bash
if [ "${0}" == "${BASH_SOURCE}" ]; then
	echo 'Script must be sourced (${0}: '"${BASH_SOURCE[0]}"')'
	return 1
fi

export GOPATH_PARENT="${IAMFEELINGCODY_GOPATH_PARENT:=${HOME}/iamfeelingcody}"
export GOPATH="${GOPATH_PARENT}/go"

export GO_VERSION="1.21.5"
export TEMPL_VERSION="v0.2.501"

export BIN_GO="${GOPATH}/bin/go"
export BIN_TEMPL="${GOPATH}/bin/templ"

export PATH="${GOPATH}/bin:${PATH}"