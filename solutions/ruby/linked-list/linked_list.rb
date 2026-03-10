class Deque
  private

  class Node
    attr_accessor :value, :prev, :succ

    def initialize(value, prev, succ)
      self.value = value
      self.prev = prev
      self.succ = succ
    end

    def link
      succ.prev = self
      prev.succ = self
    end

    def unlink
      prev = self.prev
      succ = self.succ
      succ.prev = prev
      prev.succ = succ
      value
    end
  end

  attr_accessor :head, :tail

  def initialize
    self.head = Node.new(nil, nil, nil)
    self.tail = Node.new(nil, head, nil)
    head.succ = tail
  end

  public

  def push(val)
    Node.new(val, head, head.succ).link
  end

  def unshift(val)
    Node.new(val, tail.prev, tail).link
  end

  def pop
    head.succ.unlink if head.succ != tail
  end

  def shift
    tail.prev.unlink if head.succ != tail
  end

  def count
    cur = head.succ
    count = 0
    while cur != tail
      count += 1
      cur = cur.succ
    end
    count
  end

  def delete(val)
    cur = head.succ
    while cur != tail
      if cur.value == val
        cur.unlink
        break
      end
      cur = cur.succ
    end
  end
end
