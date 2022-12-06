const std = @import("std");

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = std.mem.tokenize(u8, input, "\n");
    var groups = batch(3, []const u8, lines);

    var sum: u64 = 0;

    while (groups.next()) |group| {
        const shared: ?u8 = outer: for (group[0]) |c0| {
            for (group[1]) |c1| {
                for (group[2]) |c2| {
                    if (c0 == c1 and c1 == c2) {
                        break :outer c1;
                    }
                }
            }
        } else null;

        if (shared) |s| {
            const pri = try priority(s);
            sum += @as(u64, pri);
        }
    }

    std.debug.print("{any}\n", .{sum});
}

fn batch(comptime N: usize, comptime T: type, iter: anytype) struct {
    iter: @TypeOf(iter),
    pub fn next(this: *@This()) ?[N]T {
        var lines = [_]T{undefined} ** N;
        for (lines) |*line| {
            line.* = this.iter.next() orelse return null;
        }
        return lines;
    }
} {
    return .{ .iter = iter };
}

fn priority(c: u8) !u8 {
    return switch (c) {
        'a'...'z' => c - 'a' + 1,
        'A'...'Z' => c - 'A' + 27,
        else => (error{UnexpectedCharacter}).UnexpectedCharacter,
    };
}
