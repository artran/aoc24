#!/usr/bin/env python3

from collections import defaultdict
from typing import List, TextIO


def day5(input: TextIO):
    score = 0

    rule_definitions = extract_rule_definitions(input)
    rules = convert_defs_to_rules(rule_definitions)

    print(score)

def extract_rule_definitions(input: TextIO) -> list[tuple[int, int]]:
    rules = []

    for line in input:
        rule = line.split('|')
        if len(rule) == 2:
            rules.append(rule)

    rules = map(lambda r: (int(r[0]), int(r[1])), rules)
    return list(rules)

def convert_defs_to_rules(defs: list[tuple[int,int]]) -> dict[tuple[int,int], bool]:
    rules = defaultdict(lambda: True)

    for definition in defs:
        rules[(definition[1], definition[0])] = False

    return rules


if __name__ == '__main__':
    with open('input.txt') as input:
        day5(input)
