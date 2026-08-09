import collections.abc


def append[T](list1: list[T], list2: list[T]) -> list[T]:
    return list1 + list2


def concat[T](lists: list[list[T]]) -> list[T]:
    result = []
    for l in lists:
        result.extend(l)
    return result


def filter[T](function: collections.abc.Callable[[T], bool], a_list: list[T]) -> list[T]:
    return [i for i in a_list if function(i)]


def length[T](a_list: list[T]) -> int:
    return sum(1 for _ in a_list)


def map[S, T](function: collections.abc.Callable[[S], T], a_list: list[S]) -> list[T]:
    return [function(i) for i in a_list]


def foldl[S, T](function: collections.abc.Callable[[T, S], T], a_list: list[S], initial: T) -> T:
    result = initial
    for i in a_list:
        result = function(result, i)
    return result


def foldr[S, T](function: collections.abc.Callable[[T, S], T], a_list: list[S], initial: T) -> T:
    return foldl(function, reverse(a_list), initial)


def reverse[T](a_list: list[T]) -> list[T]:
    return list(reversed(a_list))
