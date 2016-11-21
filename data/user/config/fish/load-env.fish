#!/usr/bin/env fish

for script in $CONFIG_ROOT/fish/scripts.d/*.fish
    source $script
end

