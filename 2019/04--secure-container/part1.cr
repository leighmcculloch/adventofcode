def two_adjacent?(n : String) : Bool
  last = n.each_char.first
  n[1..].each_char do |d|
    return true if d == last
    last = d
  end
  false
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

