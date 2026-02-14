"""Determine the state of Tic Tac Toe."""
import collections

FLAT_WIN_PATTERNS = [
    # Columns
    tuple(x + 3 * y for x in range(3)) for y in range(3)
] + [
    # Rows
    tuple(x + 3 * y for y in range(3)) for x in range(3)
] + [
    # Diagonals
    (0, 4, 8), (2, 4, 6)
]
TUPLE_WIN_PATTERNS = [
    tuple((i % 3, i // 3) for i in pattern) for pattern in FLAT_WIN_PATTERNS
]


def flat_gamestate(board):
    flat_board = "".join(board)
    counts = collections.Counter(flat_board)
    
    if counts["X"] < counts["O"]:
        raise ValueError("Wrong turn order: O started")
    if counts["X"] > counts["O"] + 1:
        raise ValueError("Wrong turn order: X went twice")

    winner = {
        flat_board[x]
        for x, y, z in FLAT_WIN_PATTERNS
        # The chaining version expands to this line:
        # if flat_board[x] == flat_board[y] and flat_board[y] == flat_board[z] and flat_board[z] == player:
        if flat_board[x] == flat_board[y] == flat_board[z] != " "
    }

    if len(winner) == 2:
        raise ValueError("Impossible board: game should have ended after the game was won")
    if len(winner) == 1:
        return "win"
    if counts["X"] + counts["O"] == 9:
        return "draw"
    return "ongoing"


def tuple_gamestate(board):
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

    winner = {
        player
        for x, y, z in TUPLE_WIN_PATTERNS
        for player, moved in moves.items()
        if x in moved and y in moved and z in moved
    }

    if len(winner) == 2:
        raise ValueError("Impossible board: game should have ended after the game was won")
    if len(winner) == 1:
        return "win"
    if counts["X"] + counts["O"] == 9:
        return "draw"
    return "ongoing"


def dict_gamestate(board):
    moves = {
        (x, y): char
        for y, row in enumerate(board)
        for x, char in enumerate(row)
    }
    counts = {player: sum(p == player for p in moves.values()) for player in "XO"}
    
    if counts["X"] < counts["O"]:
        raise ValueError("Wrong turn order: O started")
    if counts["X"] > counts["O"] + 1:
        raise ValueError("Wrong turn order: X went twice")

    print(TUPLE_WIN_PATTERNS)
    winner = {
        moves[x]
        for x, y, z in TUPLE_WIN_PATTERNS
        if moves[x] == moves[y] == moves[z] != " "
    }

    if len(winner) == 2:
        raise ValueError("Impossible board: game should have ended after the game was won")
    if len(winner) == 1:
        return "win"
    print(counts)
    if counts["X"] + counts["O"] == 9:
        return "draw"
    return "ongoing"


def gamestate(board: list[str]) -> str:
    result = flat_gamestate(board)
    assert result == tuple_gamestate(board), f"{result} != {tuple_gamestate(board)}"
    assert result == dict_gamestate(board), f"{result} != {dict_gamestate(board)}"
    return result
