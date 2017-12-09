def part1(spreadsheet : String)
  rows = spreadsheet.split("\n")
  cols = rows.map do |row|
    row.split(/\s+/).map do |column|
      column.to_i
    end
  end
  matrix = cols.map do |row|
    row.max - row.min
  end
  matrix.sum
end
