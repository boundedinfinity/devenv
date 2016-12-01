#!/usr/bin/env fish

set -gx ANYENV_ROOT $DEVENV_ROOT/anyenv
set -gx PATH $ANYENV_ROOT/bin $PATH

eval "(anyenv init -)"
