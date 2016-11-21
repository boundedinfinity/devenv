#!/usr/bin/env bash

for script in $CONFIG_ROOT/bash/scripts.d/*.bash; do
    source $script
done
