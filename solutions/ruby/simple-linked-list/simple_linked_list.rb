class Element
  private

  attr_writer :datum

  def initialize(value)
    self.datum = value
  end

  public

  attr_reader :datum
  attr_accessor :next
end

class SimpleLinkedList
  private

  attr_accessor :head

  def initialize(values = nil)
    values&.each { |datum| push(Element.new(datum)) }
  end

  public

  def push(element)
    element.next = head
    self.head = element
    self
  end

  def pop
    element = head
    self.head = head.next if head
    element
  end

  def to_a
    out = []
    cur = head
    while cur
      out.append(cur.datum)
      cur = cur.next
    end
    out
  end

  def reverse!
    a = nil
    b = head
    while b
      c = b.next
      b.next = a
      a = b
      b = c
    end
    self.head = a
    self
  end
end
