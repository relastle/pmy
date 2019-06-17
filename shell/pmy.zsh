#!/usr/bin/env zsh

# Export runnable pmy path
export PATH="${GOPATH:-${HOME}/go}/src/github.com/relastle/pmy:${PATH}"

# Export pmy configuration environment variable
export PMY_RULE_PATH="${PMY_RULE_PATH:-${GOPATH:-${HOME}/go}/src/github.com/relastle/pmy/resources/pmy_rules.json}"
export PMY_TAG_DELIMITER=${PMY_TAG_DELIMITER:-"\t"}
export PMY_FUZZY_FINDER_DEFAULT_CMD=${PMY_FUZZY_FINDER_DEFAULT_CMD:-"fzf -0 -1"}
export PMY_TRIGGER_KEY=${PMY_TRIGGER_KEY:-'^ '}

# Main Function of Pmy
# Args:
#     - Left buffer string
#     - Right buffer string
# Returns:
#     - Resulting left buffer string (by name of __pmy_res_lbuffer)
#     - Resulting right buffer string (by name of __pmy_res_rbuffer)
_pmy_main() {
    # get current buffer information
    local buffer_left=${1:-""}
    local buffer_right=${2:-""}
    local test_flag=${3:-""}

    # get output from pmy
    local out="$(pmy --bufferLeft=${buffer_left} --bufferRight=${buffer_right} 2>/dev/null)"

    if [[ -z $out  ]] then
        echo "No rule was matched"
        __pmy_res_lbuffer=${buffer_left}
        __pmy_res_rbuffer=${buffer_right}
    else
        eval ${out}

        local fuzzy_finder_cmd=${__pmy_out_fuzzy_finder_cmd:-${PMY_FUZZY_FINDER_DEFAULT_CMD}}

        if [[ -z ${__pmy_out_imm_cmd} ]] then
            # get result from fzf
            local fzf_res=$(echo -n "${__pmy_out_sources}" | eval ${fuzzy_finder_cmd})
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
            local fzf_res=$(eval "${__pmy_out_imm_cmd}" | eval ${fuzzy_finder_cmd})
            local after_cmd=${__pmy_out_imm_after_cmd}
            local res=$(echo ${fzf_res} | eval ${after_cmd})
        fi
        __pmy_res_lbuffer="${__pmy_out_buffer_left}${res}"
        __pmy_res_rbuffer="${__pmy_out_buffer_right}"
    fi

    if ! [[ -z $test_flag  ]] then
        echo $__pmy_res_lbuffer
        echo $__pmy_res_rbuffer
    fi
}

pmy-widget() {
    _pmy_main ${LBUFFER} ${RBUFFER}
    zle reset-prompt
    LBUFFER=${__pmy_res_lbuffer}
    RBUFFER=${__pmy_res_rbuffer}
}

zle -N pmy-widget
bindkey ${PMY_TRIGGER_KEY} pmy-widget
