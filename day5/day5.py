#!/usr/bin/env python3

from typing import TextIO


def day5(input: TextIO):
    score = 0

    for line in input:
        line = line.strip()

    print(score)


if __name__ == '__main__':
    with open('input.txt') as input:
        day5(input)
