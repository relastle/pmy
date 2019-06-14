#!/usr/bin/env zsh
local out="$(./bench 2>/dev/null)"
eval "${out}"
# echo $__pmy_out_sources | fzf
