"""Robot Simulator."""

# Directions, complex polar mappings.
EAST, NORTH, WEST, SOUTH = (1j ** n for n in range(4))
ROTATE = {"R": -1j, "L": 1j}


class Robot:
    """A robot."""

    def __init__(self, direction: complex, x: int, y: int):
        """Initialize, mapping inputs to complex numbers."""
        self.direction = direction
        self._coord = complex(x, y)

    @property
    def coordinates(self) -> tuple[int, int]:
        """Map complex coordinates to Cartesian."""
        return (int(self._coord.real), int(self._coord.imag))

    def move(self, instructions: str) -> None:
        """Move the robot Right|Left|Advance."""
        for instruction in instructions:
            if instruction == "A":
                # Advance.
                self._coord += self.direction
            elif instruction in ROTATE:
                self.direction *= ROTATE[instruction]
            else:
                raise ValueError(f"invalid instruction {instruction}")
