"""High score utilities."""


class HighScores:
    """High score utility class."""
    def __init__(self, scores):
        """Initialize."""
        self.scores = scores

    def personal_top_three(self) -> list[int]:
        """Return the top three."""
        return sorted(self.scores, reverse=True)[:3]

    def personal_best(self) -> int:
        """Return the best score."""
        return max(self.scores)

    def latest(self) -> int:
        """Return the last score."""
        return self.scores[-1]
