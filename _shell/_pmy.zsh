

export PMY_FUZZY_FINDER_DEFAULT_CMD=${PMY_FUZZY_FINDER_DEFAULT_CMD:-"fzf -0 -1 --ansi"}
export PMY_TRIGGER_KEY=${PMY_TRIGGER_KEY:-'^ '}

_pmy_main() {
    local buffer_left=${1:-""}
    local buffer_right=${2:-""}
    local test_flag=${3:-""}

    local out="$(pmy main --buffer-left=${buffer_left} --buffer-right=${buffer_right})"

    if [[ -z $out  ]] ; then
        echo "No rule was matched"
        __pmy_res_lbuffer=${buffer_left}
        __pmy_res_rbuffer=${buffer_right}
        return
    fi

    eval ${out}

    if [[ ${__pmy_out_error_message} != '' ]] ; then
        echo ${__pmy_out_error_message}
        return
    fi

    local fuzzy_finder_cmd=${__pmy_out_fuzzy_finder_cmd:-${PMY_FUZZY_FINDER_DEFAULT_CMD}}
    local fzf_res_tag_included="$(eval ${__pmy_out_command} | eval ${fuzzy_finder_cmd})"
    if [[ -z ${__pmy_out_tag_all_empty} ]] ; then
        local tag="$(echo -n ${fzf_res_tag_included} | awk -F ${__pmy_out_tag_delimiter} 'BEGIN{ORS = ""}{print $1}' | base64)"
        tag=${tag//\//a_a} 
        tag=${tag//+/b_b} 
        tag=${tag//=/c_c} 
        local fzf_res="$(echo ${fzf_res_tag_included} | awk -F ${__pmy_out_tag_delimiter} '{for(i=2;i<NF;i++){printf("%s%s",$i,OFS=" ")}print $NF}')"
    else
        local fzf_res="${fzf_res_tag_included}"
        local tag=""
    fi
    local after_cmd_variable="__pmy_out_${tag}_after"
    local after_cmd="$(eval echo \$$after_cmd_variable)"
    local res="$(echo ${fzf_res} | eval ${after_cmd})"
    __pmy_res_lbuffer="${__pmy_out_buffer_left}${res}"
    __pmy_res_rbuffer="${__pmy_out_buffer_right}"

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
