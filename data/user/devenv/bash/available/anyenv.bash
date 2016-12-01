#!/usr/bin/env bash

export ANYENV_ROOT=$DEVENV_ROOT/anyenv
export PATH="$ANYENV_ROOT/bin:$PATH"

eval "$(anyenv init -)"
