def saddle_points(matrix):
    # Empty
    if not matrix:
        return []
    columns = len(matrix[0])
    # Not uniform shape
    if not all(len(r) == columns for r in matrix):
        raise ValueError("irregular matrix")

    col_minimums = [
        min(matrix[i][col] for i in range(len(matrix)))
        for col in range(columns)
    ]

    points = []
    for row_idx, row in enumerate(matrix):
        row_max = max(row)
        for col_idx, val in enumerate(row):
            if val == row_max and val == col_minimums[col_idx]:
                points.append({"row": row_idx + 1, "column": col_idx + 1})
    return points
