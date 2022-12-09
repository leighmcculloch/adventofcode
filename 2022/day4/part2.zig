const std = @import("std");
const tokenize = std.mem.tokenize;
const parseInt = std.fmt.parseInt;
const print = std.debug.print;

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = tokenize(u8, input, "\n");
    var count: u64 = 0;
    while (lines.next()) |line| {
        if (line.len == 0) {
            continue;
        }
        var pairs = parseIntIter(u64, tokenize(u8, line, ",-"));
        const s1 = Section{ .start = (try pairs.next()).?, .end = (try pairs.next()).? };
        const s2 = Section{ .start = (try pairs.next()).?, .end = (try pairs.next()).? };
        if (s1.overlaps(s2)) {
            count += 1;
        }
    }
    std.debug.print("{d}\n", .{count});
}

const Section = struct {
    start: u64,
    end: u64,

    pub fn contains(s: Section, i: u64) bool {
        return i >= s.start and i <= s.end;
    }

    pub fn overlaps(s: Section, o: Section) bool {
        return s.contains(o.start) or s.contains(o.end) or o.contains(s.start) or o.contains(s.end);
    }
};

fn parseIntIter(comptime T: type, iter: anytype) struct {
    iter: @TypeOf(iter),

    pub fn next(this: *@This()) !?T {
        if (this.iter.next()) |s| {
            return try parseInt(T, s, 10);
        } else {
            return null;
        }
    }
} {
    return .{ .iter = iter };
}
