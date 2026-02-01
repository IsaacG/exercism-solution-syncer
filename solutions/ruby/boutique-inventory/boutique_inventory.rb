class BoutiqueInventory
  private

  attr_reader :items

  def initialize(items)
    @items = items
  end

  public

  def item_names
    @items.map { |item| item[:name] }.sort
  end

  def cheap
    @items.select { |item| item[:price] < 30 }
  end

  def out_of_stock
    @items.select { |item| item[:quantity_by_size].count.zero? }
  end

  def stock_for_item(name)
    @items.find { |item| item[:name] == name }[:quantity_by_size]
  end

  def total_stock
    @items.sum { |item| item[:quantity_by_size].sum { |_size, count| count } }
  end
end
