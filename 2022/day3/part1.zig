const std = @import("std");

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = std.mem.tokenize(u8, input, "\n");

    var sum: u64 = 0;

    while (lines.next()) |line| {
        const compartment1 = line[0 .. line.len / 2];
        const compartment2 = line[line.len / 2 .. line.len];

        const shared: ?u8 = outer: for (compartment1) |c1| {
            for (compartment2) |c2| {
                if (c1 == c2) {
                    break :outer c1;
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

fn priority(c: u8) !u8 {
    return switch (c) {
        'a'...'z' => c - 'a' + 1,
        'A'...'Z' => c - 'A' + 27,
        else => (error{UnexpectedCharacter}).UnexpectedCharacter,
    };
}
