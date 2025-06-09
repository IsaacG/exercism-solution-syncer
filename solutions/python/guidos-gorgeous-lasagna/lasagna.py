EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2


def bake_time_remaining(done: int ) -> int:
    """Return the amount of bake time remaining."""
    return EXPECTED_BAKE_TIME - done  


def preparation_time_in_minutes(number_of_layers: int) -> int:
    """Calculate the food preparation time."""
    return PREPARATION_TIME * number_of_layers


def elapsed_time_in_minutes(layers: int, elapsed: int) -> int:
    """Return totoal elapsed time."""
    return preparation_time_in_minutes(layers) + elapsed