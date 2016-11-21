#!/usr/bin/env fish

set -gx ANYENV_ROOT $CONFIG_ROOT/anyenv
set -gx PATH $ANYENV_ROOT/bin $PATH

eval "(anyenv init -)"
