#!/usr/bin/env zsh

# Export runnable pmy path
export PATH="${GOPATH}/src/github.com/relastle/pmy:${PATH}"

# Export pmy configuration environment variable
export PMY_RULE_PATH="${GOPATH}/src/github.com/relastle/pmy/resources/pmy_rules.json"
export PMY_TAG_DELIMITER="\t"
export PMY_FUZZY_FINDER_CMD="fzf -0 -1"

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

        if [[ -z ${__pmy_out_imm_cmd} ]] then
            # get result from fzf
            local fzf_res=$(echo -n "${__pmy_out_sources}" | eval ${PMY_FUZZY_FINDER_CMD})
            # get tag
            local tag="$(echo -n ${fzf_res} | awk -F ${PMY_TAG_DELIMITER} 'BEGIN{ORS = ""}{print $1}' | base64)"
            tag=${tag//\//a_a} # original escape of base64 `/`
            tag=${tag//+/b_b} # original escape of base64 `+`
            tag=${tag//=/c_c} # original escape of base64 `+`
            # get rest statement
            local fzf_res_rest=$(echo ${fzf_res} | awk -F ${PMY_TAG_DELIMITER} '{for(i=2;i<NF;i++){printf("%s%s",$i,OFS=" ")}print $NF}')
            # get after command
            local after_cmd_variable="__pmy_out_${tag}_after"
            local after_cmd=$(eval echo \$$after_cmd_variable)
            local res=$(echo ${fzf_res_rest} | eval ${after_cmd})
        else
            # get result from fzf
            local fzf_res=$(eval "${__pmy_out_imm_cmd}" | eval ${PMY_FUZZY_FINDER_CMD})
            local after_cmd=${__pmy_out_imm_after_cmd}
            local res=$(echo ${fzf_res} | eval ${after_cmd})
        fi
        LBUFFER="${__pmy_out_buffer_left}${res}"
        RBUFFER="${__pmy_out_buffer_right}"
    fi
    zle reset-prompt
}

zle -N pmy-widget
bindkey '^ ' pmy-widget
