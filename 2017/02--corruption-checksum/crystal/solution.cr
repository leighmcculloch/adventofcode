def part1(spreadsheet : String)
	spreadsheet.split("\n").map { |r| r.split(/\s+/).map { |c| c.to_i } }.map { |r| r.max - r.min }.sum
end
