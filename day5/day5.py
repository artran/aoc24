#!/usr/bin/env python3

from collections import defaultdict
from functools import cmp_to_key
from typing import TextIO


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


def process_updates(updates: list[list[int]], rules: dict[tuple[int, int], bool]) -> int:
    total = 0
    resorted_total = 0

    def _cmp_to_key(x: int, y: int) -> int:
        return -1 if rules[x, y] else 1

    for update in updates:
        value = _process_update(update, rules)
        if value > 0:
            total += value
        else:
            resorted = sorted(update, key=cmp_to_key(_cmp_to_key))
            resorted_total += _process_update(resorted, rules)

    print(f'resorted: {resorted_total}')
    return total


def _process_update(update: list[int], rules: dict[tuple[int, int], bool]) -> int:
    for i, x in enumerate(update):
        for y in update[i:]:
            if not rules[(x, y)]:
                return 0
    
    mid_idx = len(update) // 2
    return update[mid_idx]

if __name__ == '__main__':
    with open('input.txt') as input:
        score = 0

        defs = extract_rule_definitions(input)
        input.seek(0)
        rules = convert_defs_to_rules(defs)
        updates = extract_updates(input)

    result = process_updates(updates, rules)

    print(result)
