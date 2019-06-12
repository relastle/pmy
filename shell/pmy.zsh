# test environment path variable
export PATH="${GOPATH}/src/github.com/relastle/pmy:${PATH}"
export PMY_CONFIG_PATH="${GOPATH}/src/github.com/relastle/pmy/resources/pmy_rules_neo.json"
export PMY_DELIMITER=':::'

pmy-widget() {
    # get current buffer information
    local buffer_left=${LBUFFER}
    local buffer_right=${RBUFFER}

    # get output from pmy
    local out="$(pmy --bufferLeft=${buffer_left} --bufferRight=${buffer_right} 2>/dev/null)"

    if [[ -z $out  ]] then
        echo "No rule was matched"
    else
        eval ${out}

        # get result from fzf
        local fzf_res=$(echo "${__pmy_out_sources}" | fzf -0 -1)
        LBUFFER="${__pmy_out_buffer_left}${fzf_res}"
        RBUFFER="${__pmy_out_buffer_right}"
    fi
    zle reset-prompt
}

zle -N pmy-widget
bindkey '^ ' pmy-widget
