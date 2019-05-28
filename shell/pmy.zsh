# test environment path variable
export PATH="${GOPATH}/src/github.com/relastle/pmy:${PATH}"
export PMY_CONFIG_PATH="${GOPATH}/src/github.com/relastle/pmy/resources/pmy_rules.json"
export PMY_DELIMITER=':::'

pmy-widget() {
    local buffer_left=${LBUFFER}
    local buffer_right=${RBUFFER}
    local out=$(pmy --bufferLeft=${buffer_left} --bufferRight=${buffer_right} 2>/dev/null)
    # local out=$(pmy --bufferLeft=${buffer_left} --bufferRight=${buffer_right})
    local lbuffer=$(echo -E ${out} | awk -F ${PMY_DELIMITER} '{print $1}')
    local rbuffer=$(echo -E ${out} | awk -F ${PMY_DELIMITER} '{print $2}')
    local cmd=$(echo -E ${out} | awk -F ${PMY_DELIMITER} '{print $3}')
    if [[ -z $out  ]] then
        echo "No rule was matched"
    else
        LBUFFER="${lbuffer}$(eval ${cmd} | fzf -0 -1)"
        RBUFFER="${rbuffer}"
    fi
    zle reset-prompt
}

zle -N pmy-widget
bindkey '^ ' pmy-widget
