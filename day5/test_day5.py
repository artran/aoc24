import pytest

from day5.day5 import (convert_defs_to_rules, extract_rule_definitions,
                       extract_updates)


def test_ordering_rules_length():
    with open('test.txt') as input:
        rules = extract_rule_definitions(input)

    assert len(rules) == 21


def test_rule_defs_type():
    with open('test.txt') as input:
        rules = extract_rule_definitions(input)

    for rule in rules:
        assert type(rule[0]) == int
        assert type(rule[1]) == int


def test_rule_values():
    with open('test.txt') as input:
        defs = extract_rule_definitions(input)
        rules = convert_defs_to_rules(defs)

    # Reversed ordering is denied
    assert rules[(53, 47)] is False
    assert rules[(13, 53)] is False
    assert rules[(29, 61)] is False

    # Forward ordering is allowed
    assert rules[(47, 53)] is True
    assert rules[(53, 13)] is True
    assert rules[(61, 29)] is True

    # Not mentioned ordering is allowed
    assert rules[(12, 29)] is True


def test_updates_length():
    with open('test.txt') as input:
        updates = extract_updates(input)

    assert len(updates) == 6


def test_updates_types():
    with open('test.txt') as input:
        updates = extract_updates(input)

    for update in updates:
        for page in update:
            assert type(page) == int
