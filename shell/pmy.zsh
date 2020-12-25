#!/usr/bin/env zsh

# Export runnable pmy path

# Export pmy configuration environment variable
export PMY_FUZZY_FINDER_DEFAULT_CMD=${PMY_FUZZY_FINDER_DEFAULT_CMD:-"fzf -0 -1 --ansi"}
export PMY_TRIGGER_KEY=${PMY_TRIGGER_KEY:-'^ '}

# Declarations of exit status code constants
_PMY_SUCCESS_EXIT_CODE=0
_PMY_NOT_FOUND_EXIT_CODE=204
_PMY_FATAL_EXIT_CODE=205
_PMY_FUZZY_FINDER_FAIL_EXIT_CODE=206

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
    # local out="$(pmy --bufferLeft=${buffer_left} --bufferRight=${buffer_right} 2>/dev/null)"
    local out="$(pmy main --buffer-left=${buffer_left} --buffer-right=${buffer_right})"

    # ${out} is empty, which indicates
    # there was no matching
    if [[ -z $out  ]] ; then
        return ${_PMY_NOT_FOUND_EXIT_CODE}
    fi

    # Here, evaluate the output of pmy,
    # following local variable will be declared
    # - ${__pmy_out_buffer_left}
    # - ${__pmy_out_buffer_right}
    # - ${__pmy_out_command}
    # - ${__pmy_out_<tagname>_after}
    # - ${__pmy_out_fuzzy_finder_cmd}
    # - ${__pmy_out_tag_all_empty}
    # - ${__pmy_out_tag_delimiter}
    # - ${__pmy_out_error_message}
    eval ${out}

    # Check if error occurred.
    if [[ ${__pmy_out_error_message} != '' ]] ; then
        command echo ${__pmy_out_error_message}
        return ${_PMY_FATAL_EXIT_CODE}
    fi

    local fuzzy_finder_cmd=${__pmy_out_fuzzy_finder_cmd:-${PMY_FUZZY_FINDER_DEFAULT_CMD}}
    local fzf_res_tag_included
    fzf_res_tag_included="$(eval ${__pmy_out_command} | eval ${fuzzy_finder_cmd})"
    if [[ "$?" -ne 0 ]]; then
        # when ${fuzzy_finder_cmd} is failed (or canceled)
        return ${_PMY_FUZZY_FINDER_FAIL_EXIT_CODE}
    fi

    # get result from fzf
    # get tag
    if [[ -z ${__pmy_out_tag_all_empty} ]] ; then
        local tag="$(command echo -n ${fzf_res_tag_included} | awk -F ${__pmy_out_tag_delimiter} 'BEGIN{ORS = ""}{print $1}' | base64)"
        tag=${tag//\//a_a} # original escape of base64 `/`
        tag=${tag//+/b_b} # original escape of base64 `+`
        tag=${tag//=/c_c} # original escape of base64 `+`
        # get rest statement
        local fzf_res="$(command echo ${fzf_res_tag_included} | awk -F ${__pmy_out_tag_delimiter} '{for(i=2;i<NF;i++){printf("%s%s",$i,OFS=" ")}print $NF}')"
    else
        # tag was not specified, so use line as it is
        local fzf_res="${fzf_res_tag_included}"
        local tag=""
    fi
    # get after command
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
    # Switch by the exit status code of `_pmy_main`
    case $exit_status in
        $_PMY_SUCCESS_EXIT_CODE)
            # When there was match
            zle reset-prompt
            LBUFFER=${__pmy_res_lbuffer}
            RBUFFER=${__pmy_res_rbuffer}
            ;;
        $_PMY_NOT_FOUND_EXIT_CODE)
            # When there was not match
            if [[ ${PMY_TRIGGER_KEY} == "^I" ]] then;
                # invole zsh's original completion
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
            # When error occurred in pmy-core.
            ;;
        $_PMY_FUZZY_FINDER_FAIL_EXIT_CODE)
            # When fuzzy finder was failed
            zle reset-prompt
            ;;
    esac
}

# If PMY_TRIGGER_KEY is set to `tab`
# make the key signal down to pre-refined completion.
# The main strategy is the same as
# https://github.com/junegunn/fzf/blob/master/shell/completion.zsh
[[ ${PMY_TRIGGER_KEY} == "^I" ]] && [[ -z "$pmy_default_completion" ]] && {
  binding=$(bindkey '^I')
  [[ $binding =~ 'undefined-key' ]] || pmy_default_completion=$binding[(s: :w)2]
  unset binding
}

zle -N pmy-widget

bindkey ${PMY_TRIGGER_KEY} pmy-widget
