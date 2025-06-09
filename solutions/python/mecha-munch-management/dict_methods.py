"""Functions to manage a users shopping cart items."""


def add_item(current_cart, items_to_add):
    """Add items to shopping cart."""
    for item in items_to_add:
        current_cart[item] = current_cart.get(item, 0) + 1
    return current_cart


def read_notes(notes):
    """Create user cart from an iterable notes entry."""
    return add_item({}, notes)


def update_recipes(ideas, recipe_updates):
    """Update the recipe ideas dictionary."""
    return ideas | dict(recipe_updates)


def sort_entries(cart):
    """Sort a users shopping cart in alphabetically order."""
    return dict(sorted(cart.items()))


def send_to_store(cart, isle_mapping):
    """Combine users order to isle and refrigeration information.

    :param cart: dict - users shopping cart dictionary.
    :param isle_mapping: dict - isle and refrigeration information dictionary.
    :return: dict - fulfillment dictionary ready to send to store.
    """
    return {
        item: [count, *isle_mapping[item]]
        for item, count in sorted(cart.items(), reverse=True)
    }


def update_store_inventory(fulfillment_cart, store_inventory):
    """Update store inventory levels with user order.

    :param fulfillment cart: dict - fulfillment cart to send to store.
    :param store_inventory: dict - store available inventory
    :return: dict - store_inventory updated.
    """
    return {
        item: [count - fulfillment_cart.get(item, [0])[0] or "Out of Stock", aisle, val]
        for item, (count, aisle, val) in store_inventory.items()
    }
