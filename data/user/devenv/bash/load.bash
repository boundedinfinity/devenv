#!/usr/bin/env bash

script_dir=$( cd $( dirname "${BASH_SOURCE[0]}" ) && pwd )

for script in $script_dir/enabled/*.bash; do
    source $script
done
