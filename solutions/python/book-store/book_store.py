"""Lowest price finder, by trying out different groupings.

So long as we are careful about how we add books to each grouping,
we don't need to actually track which book is in the grouping.
We just need the size of the grouping.
It doesn't matter if the grouping contains titles [A, B, C]
or [A, D, F]. It's three unique titles either way.

If we insert books into groups, one title at a time, and only
one copy of the title per group, we don't need to track what
titles are in the group.

We could skip straight to step 3 and brute force.
Steps 1 and 2 provide optimizations.

Step 1.
Whichever book shows up the most times (N times) determines
how many groupings are needed. We can't double them up so
there cannot be fewer groupings. If we have more groupings,
we can merge them.

If three books show up five times, we will have five groups
which all contain those three books as the starting scenario.

Step 2.
Looking at the remaining books, pick whichever occurs the most
(M times).
Since all the groups are the same, place that next book in the
first M groups. Different combos will result in the same
exact groupings, though in differing orders.

Step 3.
Brute force.
At this point, we need to consider more possibilities.
For each remaining title, attempt to place it in each
combination of buckets and see which results in the
lowest price.
"""

import collections
import copy
import itertools
from typing import List

# The price paid for various group sizes.
DISCOUNT = {1: 1.00, 2: 0.95, 3: 0.90, 4: 0.80, 5: 0.75}
# The total price per grouping, by group size.
PRICE = {n: int(800 * n * d) for n, d in DISCOUNT.items()}


class Grouping:
  """Represents a set of books, organized into groups.

  All we actually need to store here is the count of unique
  titles per group so long as we manage groupings in a way
  that prevents adding a title twice to one group.
  """

  def __init__(self, bookset: List[int]):
    self._bookset = bookset

  def Price(self) -> int:
    """Returns the discounted price of the books."""
    return sum(PRICE[b] for b in self._bookset)

  def WithBookIn(self, group_idxs: List[int]) -> 'Grouping':
    """Return a new grouping with a new title added to specific groups."""
    bookset = list(self._bookset)
    for idx in group_idxs:
      bookset[idx] += 1
    return Grouping(bookset)


class Possibilities:
  """Calculate the various ways to group books."""

  def __init__(self, basket: List[int]):
    self._basket = collections.Counter(basket)
    self._num_groups = max(self._basket.values() or [0])
    self._possibilities = [Grouping([0] * self._num_groups)]

  def Price(self) -> int:
    """Return the price of the cheapest possible grouping."""
    return min(p.Price() for p in self._possibilities)

  def Cheapest(self) -> int:
    """Compute the cheapest price possible and return it.

    Apply a series of steps that shuffles titles into groups.
    This generates a number of possible groupings and modifies
    class state in a destructive manner.
    """
    steps = [
      self.AddMostCommonTitle,
      self.AddSecondMostCommon,
      self.AddAllCombos]
    while self._basket and steps:
      steps.pop(0)()
    return self.Price()

  def _UpdateFirstPossibility(self, book_count: int, group_count: int):
    """Inplace update of the titles in the first possibility."""
    for group in range(group_count):
      self._possibilities[0]._bookset[group] += book_count

  def AddMostCommonTitle(self):
    """Create initial groups based on the most common title.

    Find the title(s) that appear the most times. Create N
    groups for N occurances. Place those N title(s) into
    the N groups and remove them from the basket.
    """
    basket = self._basket

    books_in_each_group = [
        book for book, count in basket.items()
        if count == self._num_groups]
    num_books = len(books_in_each_group)

    self._UpdateFirstPossibility(num_books, self._num_groups)
    for book in books_in_each_group:
      del basket[book]

  def AddSecondMostCommon(self):
    """Place the second most common title into the first M groups.

    Picking exactly one title from the basket (which shows up
    more than any other group, ie M times), place that title
    into the first M groups.
    
    Since the groups at this point are all the same, the order
    does not matter for this step.
    """
    basket = self._basket

    second_highest_count = max(basket.values())
    book = [book for book, count in basket.items() if count == second_highest_count][0]
    self._UpdateFirstPossibility(1, second_highest_count)
    del basket[book]

  def AddAllCombos(self):
    """Brute force stage. Try all combos.

    For each title in the basket, try placing it in every
    combination of groups. Each try creates a new possibility.
    """
    possibilities = self._possibilities
    basket = self._basket
    # For the remaining books, we need to try different combinations.
    # For each title, attempt to add the title to each combo of groups.
    for book, count in basket.items():
      new_possibilities = []
      for grouping in possibilities:
        for selected_groups in itertools.combinations(range(self._num_groups), count):
          new_groupings = grouping.WithBookIn(selected_groups)
          new_possibilities.append(new_groupings)
      possibilities = new_possibilities
    self._possibilities = possibilities


def total(basket):
  return Possibilities(basket).Cheapest()


# vim:ts=2:sw=2:expandtab
