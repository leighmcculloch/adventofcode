input = File.read("input.txt")
lines = input.split("\n")

class Pos
    attr_accessor :x, :y
    attr_reader :visited, :next
    def initialize(x, y, nxt)
        @x = x
        @y = y
        @visited = Set.new
        @visited.add([@x, @y])
        @next = nxt
    end
    def right
        @x += 1
        @visited.add([@x, @y])
        if @next && !self.touches(@next)
            @next.y = @y
            @next.right
        end
    end
    def left
        @x -= 1
        @visited.add([@x, @y])
        if @next && !self.touches(@next)
            @next.y = @y
            @next.left
        end
    end
    def up
        @y -= 1
        @visited.add([@x, @y])
        if @next && !self.touches(@next)
            @next.x = @x
            @next.up
        end
    end
    def down
        @y += 1
        @visited.add([@x, @y])
        if @next && !self.touches(@next)
            @next.x = @x
            @next.down
        end
    end
    def touches(o)
        (@x - o.x).abs <= 1 && (@y - o.y).abs <= 1
    end
end

tail = Pos.new(0, 0, nil)
head = Pos.new(0, 0, tail)

def print_map(head, tail)
    puts "-"
    start = Pos.new(0, 0)
    min = Pos.new(
        ([start.x, head.x, tail.x]+(tail.visited.map { |x, _| x })).min,
        ([start.y, head.y, tail.y]+(tail.visited.map { |_, y| y })).min,
    )
    max = Pos.new(
        ([start.x, head.x, tail.x]+(tail.visited.map { |x, _| x })).max,
        ([start.y, head.y, tail.y]+(tail.visited.map { |_, y| y })).max,
    )
    (min.y..max.y).each do |y|
        (min.x..max.x).each do |x|
            case
            when head.y == y && head.x == x
                print 'H'
            when tail.y == y && tail.x == x
                print 'T'
            when start.y == y && start.x == x
                print 's'
            when tail.visited.include?([x, y])
                print '#'
            else
                print '.'
            end
        end
        puts
    end
end

lines.each do |line|
    dir, dist = line.split(" ")
    dist.to_i.times do
        case dir
        when 'R'; head.right
        when 'L'; head.left
        when 'U'; head.up
        when 'D'; head.down
        end
        # print_map(head, tail)
    end
end
puts "head count: #{head.visited.count}"
puts "tail count: #{tail.visited.count}"
