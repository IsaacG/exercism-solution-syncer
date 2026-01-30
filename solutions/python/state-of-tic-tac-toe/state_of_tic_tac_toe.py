"""Determine the state of Tic Tac Toe."""
WIN_PATTERNS = [
    # Columns
    ((x, 0), (0, 1)) for x in range(3)
] + [
    # Rows
    ((0, y), (1, 0)) for y in range(3)
] + [
    # Diagonals
    ((0, 0), (1, 1)), ((0, 2), (1, -1))
]


def gamestate(board: list[str]) -> str:
    """Return the state of a game."""
    moves: dict[str, set[tuple[int, int]]] = {c: set() for c in "XO"}
    for y, row in enumerate(board):
        for x, char in enumerate(row):
            if char != " ":
                moves[char].add((x, y))
    counts = {player: len(m) for player, m in moves.items()}
    if counts["X"] < counts["O"]:
        raise ValueError("Wrong turn order: O started")
    if counts["X"] > counts["O"] + 1:
        raise ValueError("Wrong turn order: X went twice")
    winner = set()
    for player, moved in moves.items():
        for (x, y), (dx, dy) in WIN_PATTERNS:
            if all((x + dx * i, y + dy * i) in moved for i in range(3)):
                winner.add(player)
                break
    if len(winner) == 2:
        raise ValueError("Impossible board: game should have ended after the game was won")
    if len(winner) == 1:
        return "win"
    if sum(counts.values()) == 9:
        return "draw"
    return "ongoing"
