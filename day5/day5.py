#!/usr/bin/env python3

from collections import defaultdict
from typing import List, TextIO


def main(input: TextIO):
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


def extract_updates(input: TextIO) -> list[list[int]]:
    updates = []

    for line in input:
        update = line.split(',')
        if len(update) > 1:
            updates.append(update)

    updates = map(lambda x: [int(y) for y in x], updates)

    return list(updates)


if __name__ == '__main__':
    with open('input.txt') as input:
        main(input)
