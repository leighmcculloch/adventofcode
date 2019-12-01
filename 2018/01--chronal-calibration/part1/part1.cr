def calculate_sum(filename)
  f = File.open(filename)

  sum = 0

  loop do
    line = f.gets(delimiter: '\n', chomp: true)
    break if line.nil?
    sum += line.to_i
  end
  f.close

  return sum
end

puts "sum: #{calculate_sum("2018/01--chronal-calibration/input.txt")}"
