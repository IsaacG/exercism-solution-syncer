"""Lowest price finder, by trying out different groupings.

So long as we are careful about how we add books to each grouping,
we don't need to actually track which book is in the grouping.
We just need the size of the grouping.
It doesn't matter if the grouping contains titles [A, B, C]
or [A, D, F]. It's three unique titles either way.

If we insert books into groups, one title at a time, and only
one copy of the title per group, we don't need to track what
titles are in the group.

Whichever book shows up the most times (N times) determines
how many groupings are needed. We can't double them up so
there cannot be fewer groupings. If we have more groupings,
we can merge them.

If three books show up five times, we will have five groups
which all contain those three books as the starting scenario.

Looking at the remaining books, pick whichever occurs the most
(M times).
Since all the groups are the same, place that next book in the
first M groups. Different combos will result in the same
exact groupings, though in differing orders.

At this point, we need to consider more possibilities.
For each remaining title, attempt to place it in each
combination of buckets and see which results in the
lowest price.
"""

import collections
import copy
import itertools

# The price paid for various group sizes.
DISCOUNT = {1: 1.00, 2: 0.95, 3: 0.90, 4: 0.80, 5: 0.75}
# The total price per grouping, by group size.
PRICE = {n: int(800 * n * d) for n, d in DISCOUNT.items()}


class Grouping(object):

  def __init__(self, bookset):
    self._bookset = bookset

  def Price(self):
    return sum(PRICE[b] for b in self._bookset)

  def WithBookIn(self, buckets):
    bookset = list(self._bookset)
    for b in buckets:
      bookset[b] += 1
    return Grouping(bookset)


def total(basket):
  counts = collections.Counter(basket)
  # Empty case.
  if not counts:
    return 0
  # Find the item we got the most times, N.
  # That will determine the number of groupings we need.
  num_groups = max(counts.values())
  # Books that appear N times will go in each group.
  books_in_each_group = [book for book, count in counts.items() if count == num_groups]
  # Make a grouping that contains N sets with those common books.
  grouping = Grouping([len(books_in_each_group) for _ in range(num_groups)])
  # These books are accounted for. Remove.
  for book in books_in_each_group:
    del counts[book]
  # If no more books, nothing to do. Return price.
  if not counts:
    return grouping.Price()
  # Select the second most common book, occurring M times.
  second_highest_count = max(counts.values())
  book = [book for book, count in counts.items() if count == second_highest_count][0]
  # Since order doesn't matter and existing groups are all the same,
  # add this book to the first M groups.
  grouping = grouping.WithBookIn(range(second_highest_count))
  del counts[book]
  # If no more books, nothing to do. Return price.
  if not counts:
    return grouping.Price()

  # For the remaining books, we need to try different combinations.
  # For each title, attempt to add the title to each combo of groups.
  possibilities = [grouping]
  for book, count in counts.items():
    new_possibilities = []
    for groupings in possibilities:
      for selected_groups in itertools.combinations(range(num_groups), count):
        new_groupings = groupings.WithBookIn(selected_groups)
        new_possibilities.append(new_groupings)
    possibilities = new_possibilities

  return min(g.Price() for g in possibilities)




# vim:ts=2:sw=2:expandtab
