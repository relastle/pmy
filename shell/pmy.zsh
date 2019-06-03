# test environment path variable
export PATH="${GOPATH}/src/github.com/relastle/pmy:${PATH}"
export PMY_CONFIG_PATH="${GOPATH}/src/github.com/relastle/pmy/resources/pmy_rules.json"
export PMY_DELIMITER=':::'

pmy-widget() {
    local buffer_left=${LBUFFER}
    local buffer_right=${RBUFFER}
    local out="$(pmy --bufferLeft=${buffer_left} --bufferRight=${buffer_right} 2>/dev/null)"
    if [[ -z $out  ]] then
        echo "No rule was matched"
    else
        eval ${out}
        LBUFFER="${__pmy_out_buffer_left}$(eval ${__pmy_out_command} | fzf -0 -1)"
        RBUFFER="${__pmy_out_buffer_right}"
    fi
    zle reset-prompt
}

zle -N pmy-widget
bindkey '^ ' pmy-widget
