#!/bin/python

def factors(num):
    factors = []
    divisor = 2
  
    while num > 1:
        # Repeat divide and append.
        while num % divisor == 0:
            num /= divisor
            factors.append(divisor)
        divisor += 1

    return factors
