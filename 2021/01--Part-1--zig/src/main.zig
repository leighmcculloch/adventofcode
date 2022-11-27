const std = @import("std");

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = std.mem.tokenize(u8, input, "\n");
    var nums = mapToInt(u64, lines);
    var lastn = try (nums.next() orelse return);
    var increased : u64 = 0;
    while (nums.next()) |maybe_n| {
        const n = try maybe_n;
        if (n > lastn) {
            increased += 1;
        }
        lastn = n;
    }
    std.log.info("{d}", .{increased});
}

pub fn mapToInt(comptime T: type, iter: anytype) struct {
    iter: @TypeOf(iter),
    pub fn next(self: *@This()) ?std.fmt.ParseIntError!T {
        const res = self.iter.next() orelse return null;
        return std.fmt.parseInt(T, res, 10);
    }
} {
    return .{ .iter = iter };
}
