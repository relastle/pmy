#!/usr/bin/env zsh
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
    # echo $out

    if [[ -z $out  ]] then
        echo "No rule was matched"
    else
        eval ${out}

        # get result from fzf
        local fzf_res=$(echo "${__pmy_out_sources}" | fzf -0 -1)
        # get tag
        local tag=$(echo ${fzf_res} | awk '{print $1}')
        # get rest statement
        local fzf_res_rest=$(echo ${fzf_res} | awk '{for(i=2;i<NF;i++){printf("%s%s",$i,OFS=" ")}print $NF}')
        # get after command
        local after_cmd_variable="__pmy_out_${tag}_after"
        local after_cmd=$(eval echo \$$after_cmd_variable)
        local res=$(echo ${fzf_res_rest} | eval ${after_cmd})
        LBUFFER="${__pmy_out_buffer_left}${res}"
        RBUFFER="${__pmy_out_buffer_right}"
    fi
    zle reset-prompt
}

zle -N pmy-widget
bindkey '^ ' pmy-widget
