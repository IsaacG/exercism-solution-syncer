"""Compute years on various planets."""
EARTH_SECONDS = 60 * 60 * 24 * 365.25
PLANET_RATIOS = {
    'mercury': 0.2408467,
    'venus': 0.61519726,
    'earth': 1.0,
    'mars': 1.8808158,
    'jupiter': 11.862615,
    'saturn': 29.447498,
    'uranus': 84.016846,
    'neptune': 164.79132
}


class SpaceAge:
    """Compute years on various planets."""

    def __init__(self, seconds: int):
        self.seconds = seconds
        for planet, ratio in PLANET_RATIOS.items():
            setattr(self, f"on_{planet}", lambda: round(self.years * ratio, 2))

    @property
    def years(self) -> int:
        """Return Earth years."""
        return self.seconds / EARTH_SECONDS
