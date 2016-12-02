#!/usr/bin/env fish

set script_dir (cd (dirname (status -f)); and pwd)

for script in $script_dir/enabled/*.fish
    source $script
end

