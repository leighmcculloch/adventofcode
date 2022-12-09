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
        if ((s1.start <= s2.start and s1.end >= s2.end) or (s2.start <= s1.start and s2.end >= s1.end)) {
            count += 1;
        }
    }
    std.debug.print("{d}\n", .{count});
}

const Section = struct {
    start: u64,
    end: u64,
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
