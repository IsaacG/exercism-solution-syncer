"""Determine the state of Tic Tac Toe."""
import collections

WIN_PATTERNS = [
    # Columns
    [(x, y) for x in range(3)] for y in range(3)
] + [
    # Rows
    [(x, y) for y in range(3)] for x in range(3)
] + [
    # Diagonals
    [(i, i) for i in range(3)],
    [(i, 2 - i) for i in range(3)]
]

USE_FLAT_BOARD = False


def gamestate(board: list[str]) -> str:
    """Return the state of a game."""

    if USE_FLAT_BOARD:
        flat_board = "".join(board)
        counts = collections.Counter(flat_board)
        counts.pop(" ", None)
        win_patterns = [[x + y * 3 for x, y in pattern] for pattern in WIN_PATTERNS]
    else:
        moves: dict[str, set[tuple[int, int]]] = {c: set() for c in "XO"}
        for y, row in enumerate(board):
            for x, char in enumerate(row):
                if char != " ":
                    moves[char].add((x, y))
        counts = {player: len(m) for player, m in moves.items()}
        win_patterns = WIN_PATTERNS
    
    if counts["X"] < counts["O"]:
        raise ValueError("Wrong turn order: O started")
    if counts["X"] > counts["O"] + 1:
        raise ValueError("Wrong turn order: X went twice")

    winner = set()

    for player in "XO":
        if USE_FLAT_BOARD:
            test = lambda x, y, z: flat_board[x] == flat_board[y] and flat_board[y] == flat_board[z] and flat_board[z] == player
        else:
            moved = moves[player]
            test = lambda x, y, z: x in moved and y in moved and z in moved
        for x, y, z in win_patterns:
            if test(x, y, z):
                winner.add(player)
                break

    if len(winner) == 2:
        raise ValueError("Impossible board: game should have ended after the game was won")
    if len(winner) == 1:
        return "win"
    if sum(counts.values()) == 9:
        return "draw"
    return "ongoing"
