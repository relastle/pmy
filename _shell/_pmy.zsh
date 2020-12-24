

export PMY_FUZZY_FINDER_DEFAULT_CMD=${PMY_FUZZY_FINDER_DEFAULT_CMD:-"fzf -0 -1 --ansi"}
export PMY_TRIGGER_KEY=${PMY_TRIGGER_KEY:-'^ '}

_PMY_SUCCESS_EXIT_CODE=0
_PMY_NOT_FOUND_EXIT_CODE=204
_PMY_FATAL_EXIT_CODE=205

_pmy_main() {
    local buffer_left=${1:-""}
    local buffer_right=${2:-""}
    local test_flag=${3:-""}

    local out="$(pmy main --buffer-left=${buffer_left} --buffer-right=${buffer_right})"

    if [[ -z $out  ]] ; then
        return ${_PMY_NOT_FOUND_EXIT_CODE}
    fi

    eval ${out}

    if [[ ${__pmy_out_error_message} != '' ]] ; then
        command echo ${__pmy_out_error_message}
        return ${_PMY_FATAL_EXIT_CODE}
    fi

    local fuzzy_finder_cmd=${__pmy_out_fuzzy_finder_cmd:-${PMY_FUZZY_FINDER_DEFAULT_CMD}}
    local fzf_res_tag_included="$(eval ${__pmy_out_command} | eval ${fuzzy_finder_cmd})"
    if [[ -z ${__pmy_out_tag_all_empty} ]] ; then
        local tag="$(command echo -n ${fzf_res_tag_included} | awk -F ${__pmy_out_tag_delimiter} 'BEGIN{ORS = ""}{print $1}' | base64)"
        tag=${tag//\//a_a} 
        tag=${tag//+/b_b} 
        tag=${tag//=/c_c} 
        local fzf_res="$(command echo ${fzf_res_tag_included} | awk -F ${__pmy_out_tag_delimiter} '{for(i=2;i<NF;i++){printf("%s%s",$i,OFS=" ")}print $NF}')"
    else
        local fzf_res="${fzf_res_tag_included}"
        local tag=""
    fi
    local after_cmd_variable="__pmy_out_${tag}_after"
    local after_cmd="$(eval command echo \$$after_cmd_variable)"
    local res="$(command echo ${fzf_res} | eval ${after_cmd})"
    __pmy_res_lbuffer="${__pmy_out_buffer_left}${res}"
    __pmy_res_rbuffer="${__pmy_out_buffer_right}"

    if ! [[ -z $test_flag  ]] then
        command echo $__pmy_res_lbuffer
        command echo $__pmy_res_rbuffer
    fi

    return ${_PMY_SUCCESS_EXIT_CODE}
}

pmy-widget() {
    local buffer_left=${LBUFFER}
    local buffer_right=${RBUFFER}
    _pmy_main ${LBUFFER} ${RBUFFER}
    local exit_status=$?
    case $exit_status in
        $_PMY_SUCCESS_EXIT_CODE)
            zle reset-prompt
            LBUFFER=${__pmy_res_lbuffer}
            RBUFFER=${__pmy_res_rbuffer}
            ;;
        $_PMY_NOT_FOUND_EXIT_CODE)
            if [[ ${PMY_TRIGGER_KEY} == "^I" ]] then;
                zle ${pmy_default_completion:-expand-or-complete}
            else
                command echo "No rule was matched"
                __pmy_res_lbuffer=${buffer_left}
                __pmy_res_rbuffer=${buffer_right}
                LBUFFER=${__pmy_res_lbuffer}
                RBUFFER=${__pmy_res_rbuffer}
                zle reset-prompt
            fi
            ;;
        $_PMY_FATAL_EXIT_CODE)
            ;;
    esac
}

[[ ${PMY_TRIGGER_KEY} == "^I" ]] && [[ -z "$pmy_default_completion" ]] && {
  binding=$(bindkey '^I')
  [[ $binding =~ 'undefined-key' ]] || pmy_default_completion=$binding[(s: :w)2]
  unset binding
}

zle -N pmy-widget

bindkey ${PMY_TRIGGER_KEY} pmy-widget
