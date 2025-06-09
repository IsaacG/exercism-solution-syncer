"""Compute trees from traversals."""

def tree_from_traversals(preorder: list[str], inorder: list[str]) -> dict:
    """Return a tree from traversals."""
    if len(preorder) != len(inorder):
        raise ValueError("traversals must have the same length")
    if set(preorder) != set(inorder):
        raise ValueError("traversals must have the same elements")
    if len(preorder) != len(set(preorder)):
        raise ValueError("traversals must contain unique items")

    if not preorder:
        return {}

    # The first element of the preorder is always the root.
    root = preorder[0]
    # The inorder can be partitioned on the root to give the left and right inorder.
    inorder_idx = inorder.index(root)
    inorder_sub = inorder[:inorder_idx], inorder[inorder_idx + 1:]
    # The preorder for the left and right can be found by taking the elements from the
    # parent preorder and filtering for elements in the inorder.
    preorder_sub = [[x for x in preorder if x in ios] for ios in inorder_sub]
    return {
        "v": root,
        "l": tree_from_traversals(preorder_sub[0], inorder_sub[0]),
        "r": tree_from_traversals(preorder_sub[1], inorder_sub[1]),
    }
