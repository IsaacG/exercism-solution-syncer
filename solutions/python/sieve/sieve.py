def primes(limit: int) -> list[int]:
    limit += 1
    sieve = [True for i in range(limit)]
    for candidate in range(2, limit):
        if not sieve[candidate]:
            continue
        for not_prime in range(2 * candidate, limit, candidate):
            sieve[not_prime] = False
    return [candidate for candidate, is_prime in enumerate(sieve) if candidate > 1 and is_prime]
