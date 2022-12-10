const std = @import("std");

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = std.mem.tokenize(u8, input, "\n");
    var nums = map(std.fmt.ParseIntError!u64, lines, struct { pub fn f(s: []const u8) !u64 { return std.fmt.parseInt(u64, s, 10); } }.f);
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

pub fn map(
    comptime U: type,
    iter: anytype,
    f: *const fn (t: @typeInfo(@TypeOf(blk: { var copy = iter; break :blk copy.next(); })).Optional.child) U,
) struct {
    iter: @TypeOf(iter),
    f: *const fn (t: @typeInfo(@TypeOf(blk: { var copy = iter; break :blk copy.next(); })).Optional.child) U,
    pub fn next(self: *@This()) ?U {
        const res = self.iter.next() orelse return null;
        return self.f(res);
    }
} {
    return .{
        .iter = iter,
        .f = f,
    };
}
