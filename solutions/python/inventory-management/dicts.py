import collections

def create_inventory(items) -> dict[str, int]:
    """Create a counter.

    :param items: list - list of items to create an inventory from.
    :return:  dict - the inventory dictionary.
    """
    return collections.Counter(items)


def add_items(inventory: dict[str, int], items: list[str]) -> dict[str, int]:
    """Add an item to the inventory.

    :param inventory: dict - dictionary of existing inventory.
    :param items: list - list of items to update the inventory with.
    :return:  dict - the inventory dictionary update with the new items.
    """
    for item in items:
        inventory.setdefault(item, 0)
        inventory[item] += 1
    return inventory


def decrement_items(inventory: dict[str, int], items: list[str]) -> dict[str, int]:
    """Remove an item from the inventory.

    :param inventory: dict - inventory dictionary.
    :param items: list - list of items to decrement from the inventory.
    :return:  dict - updated inventory dictionary with items decremented.
    """
    for item in items:
        if inventory[item]:
            inventory[item] -= 1
    return inventory


def remove_item(inventory: dict[str, int], item: str) -> dict[str, int]:
    """
    :param inventory: dict - inventory dictionary.
    :param item: str - item to remove from the inventory.
    :return:  dict - updated inventory dictionary with item removed.
    """
    if item in inventory:
        del inventory[item]
    return inventory


def list_inventory(inventory: dict[str, int]) -> list[tuple[str, int]]:
    """List the inventory.

    :param inventory: dict - an inventory dictionary.
    :return: list of tuples - list of key, value pairs from the inventory dictionary.
    """
    return [(item, count) for item, count in inventory.items() if count]