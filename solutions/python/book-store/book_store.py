import copy


DISCOUNT = {1: 1, 2: 0.95, 3: 0.90, 4: 0.80, 5: 0.75}


class BookSet(object):

  def __init__(self, n):
    self._got = [n]

  def CanAdd(self, n):
    return n not in self._got

  def Add(self, n):
    if not self.CanAdd(n):
      raise ValueError('Cannot add %d to the set.' % n)
    self._got.append(n)

  def Price(self):
    return int(800 * DISCOUNT[len(self._got)] * len(self._got))


class Table(object):
  def __init__(self, basket, sets=None):
    self._basket = sorted(basket)
    if sets:
      self._booksets = copy.deepcopy(sets)
    else:
      self._booksets = []

  def Price(self):
    return sum(b.Price() for b in self._booksets)

  def Cheapest(self):
    if not self._basket:
      return self.Price()
    basket = copy.copy(self._basket)
    book = basket.pop()
    options = [Table(basket, self._booksets).AddBook(book, None)]
    for i in range(len(self._booksets)):
      attempt = Table(basket, self._booksets)
      if attempt.CanAdd(book, i):
        options.append(attempt.AddBook(book, i))
    return min(t.Cheapest() for t in options)

  def CanAdd(self, book, bucket=None):
    if bucket is None:
      return True
    return self._booksets[bucket].CanAdd(book)

  def AddBook(self, book, bucket=None):
    if bucket is None:
      self._booksets.append(BookSet(book))
    else:
      self._booksets[bucket].Add(book)
    return self


def total(basket):
  return Table(basket).Cheapest()


# vim:ts=2:sw=2:expandtab
