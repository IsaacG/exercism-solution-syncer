module Grid
  def self.row_maxs(input)
    input.zip(input.map(&:max)).map do |row, max|
      row.map.with_index.select { |v, _| v == max }.map { |_, i| i }
    end
  end

  def self.col_mins(input)
    cols = input.transpose
    cols.zip(cols.map(&:min)).map do |col, min|
      col.map.with_index.select { |v, _| v == min }.map { |_, i| i }
    end
  end

  def self.saddle_points(input)
    row_max = row_maxs(input)
    col_min = col_mins(input)

    points = []
    input.each_index do |row|
      input[0].each_index do |col|
        if row_max[row].include?(col) && col_min[col].include?(row)
          points.append({ 'row' => row + 1, 'column' => col + 1 })
        end
      end
    end
    points
  end

  private_class_method :row_maxs, :col_mins

  public_class_method :saddle_points
end
