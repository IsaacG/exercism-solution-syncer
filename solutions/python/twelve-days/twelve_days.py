#!/bin/python

OPENING = "On the %s day of Christmas my true love gave to me: "
ITEM = [
    "zero",
    "and a Partridge in a Pear Tree",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming",
]
NUMBERS = [
    "zeroth", "first", "second", "third", "fourth", "fifth", "sixth",
    "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"
]


def verse(i):
    out = OPENING % NUMBERS[i]
    items = ITEM[i:0:-1]
    out += ", ".join(items).removeprefix("and ") + "."
    return out

    
def recite(start, end):
    return [verse(i) for i in range(start, end + 1)]
