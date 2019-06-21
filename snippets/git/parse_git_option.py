#!/usr/bin/env python
# -*- coding: utf-8 -*-


from glob import glob
import re
import os
from typing import (  # noqa
        Any,
        List,
        Dict,
        Tuple,
        )


def extract_necessary_lines(lines: List[str]) -> List[str]:
    res_lines = []
    is_option = False
    for line in lines:
        if line.strip() == '':
            continue
        if re.match(r"^ *-", line):
            is_option = True
            res_lines.append(line.strip())
            continue
        if is_option:
            res_lines.append(line.strip())
    return res_lines


def collapse_line(lines: List[str]) -> List[str]:
    res_lines = []
    line_tmp = ''
    for line in lines:
        if re.match(r"^ *-", line):
            if line_tmp != '':
                res_lines.append(line_tmp)
            line_tmp = line
        else:
            line_tmp += " " + line
    if line_tmp != '':
        res_lines.append(line_tmp)
    return res_lines


def split_opt_and_desc(lines: List[str]) -> List[Tuple[str, str]]:
    res_tuples = []
    opt_reg_exp = r'(-[A-Za-z0-9\-:|]+){0,1},{0,1} *(--[A-Za-z0-9\-]+){0,1} *(<[A-Za-z0-9\-]+>){0,1} *(\[.*\]){0,1}'  # noqa
    for line in lines:
        m = re.match(opt_reg_exp, line)
        res_tuple = (
            line[m.start():m.end()].replace(',', ' ').replace('[', ' [').strip(),  # noqa
            line[m.end():].strip(),
        )
        res_tuples.append(res_tuple)
    return res_tuples


def organize_tuples(ts: List[Tuple[str, str]]) -> List[str]:
    res_lines: List[str] = []
    if len(ts) == 0:
        return res_lines
    max_option_len = max([len(t[0]) for t in ts])
    for t in ts:
        res_lines.append('{}{}{}'.format(
            t[0],
            ' ' * (max_option_len - len(t[0]) + 1),
            t[1],
        ))
    return res_lines


def process_one_file(path: str) -> None:
    with open(path) as f:
        lines = f.readlines()

    lines = extract_necessary_lines(lines)
    lines = collapse_line(lines)
    ts = split_opt_and_desc(lines)
    lines = organize_tuples(ts)
    os.remove(path)
    with open(path, 'w') as f:
        for line in lines:
            f.write(line + '\n')
    return


def main() -> None:
    paths = glob('./*_option.txt')
    for path in paths:
        process_one_file(path)

    return


if __name__ == '__main__':
    main()
