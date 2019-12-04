def two_adjacent?(n : String) : Bool
  last = nil
  adjacenting = false
  num_adjacent = Array(Int32).new
  n.each_char do |d|
    if d == last
      num_adjacent << 2 if !adjacenting
      num_adjacent[-1] += 1 if adjacenting
      adjacenting = true
    else
      adjacenting = false
    end
    last = d
  end
  num_adjacent.any? { |na| na == 2 }
end

def never_decreases?(n : String) : Bool
  last = n.each_char.first
  n.each_char do |d|
    return false if d < last
    last = d
  end
  true
end

range = 240920..789857

puts range

count = range.count do |n|
  s = n.to_s
  two_adjacent?(s) && never_decreases?(s)
end

puts count
