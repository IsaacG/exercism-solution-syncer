EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2


def bake_time_remaining(time_in_the_oven):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """
    return EXPECTED_BAKE_TIME - time_in_the_oven


def preparation_time_in_minutes(number_of_layers):
    """Calculate the prep time required."""
    return number_of_layers * PREPARATION_TIME


def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the total elapsed time."""
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
