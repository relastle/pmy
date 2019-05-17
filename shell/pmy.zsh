# test environment path variable
export PATH="${GOPATH}/src/github.com/relastle/pmy:${PATH}"
export PMY_CONFIG_PATH="${GOPATH}/src/github.com/relastle/pmy/resources/test_setting.json"

pmy-widget() {
    local buffer_left=${LBUFFER}
    local buffer_right=${RBUFFER}
    local cmd=$(pmy --bufferLeft=${buffer_left} --bufferRight=${buffer_right} 2>/dev/null)
    if [[ -z $cmd  ]] then
        echo "No rule was matched"
    else
        LBUFFER="${LBUFFER}$(eval ${cmd} | fzf -0 -1)"
    fi
    zle reset-prompt
}
zle -N pmy-widget
bindkey '^ ' pmy-widget
