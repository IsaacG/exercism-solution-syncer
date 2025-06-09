EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2


# TODO: define the 'bake_time_remaining()' function
def bake_time_remaining(done):
    '''
    :param elapsed_bake_time: int baking time already elapsed
    :return: int remaining bake time derived from 'EXPECTED_BAKE_TIME'

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    '''

    return EXPECTED_BAKE_TIME - done

def preparation_time_in_minutes(layers):
    """Return prep time for the layers."""
    return PREPARATION_TIME * layers

def elapsed_time_in_minutes(layers, elapsed):
    """Return totoal elapsed time."""
    return preparation_time_in_minutes(layers) + elapsed