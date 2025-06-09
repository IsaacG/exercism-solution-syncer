"""Binary search."""


def find(search_list: list[int], value: int) -> int:
    """Return the index of a value in a list, using binary search."""
    offset = 0
    while search_list:
        # Take the center point as the pivot.
        pivot = len(search_list) // 2
        if search_list[pivot] == value:
            return offset + pivot
        # Reduce the search to the upper/lower half of the list.
        if search_list[pivot] > value:
            search_list = search_list[:pivot]
        else:
            search_list = search_list[pivot + 1:]
            # Track the offset of the pivot from the start.
            offset += pivot + 1
    raise ValueError("value not in array")
