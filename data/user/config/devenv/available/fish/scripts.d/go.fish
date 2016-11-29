#!/usr/bin/env fish

set -gx GOPATH {{ .GoPath }}
set -gx PATH $PATH $GOPATH/bin
