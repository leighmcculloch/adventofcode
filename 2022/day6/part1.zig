const std = @import("std");
const input = std.mem.trimRight(u8, @embedFile("input.txt"), "\n");

pub fn main() void {
    const size = 4;
    var map = [_]usize{0} ** 26;
    var i: usize = 0;
    while (i < input.len) : (i += 1) {
        const c = input[i];
        map[c - 'a'] += 1;
        if (i + 1 >= size) {
            for (map) |cc| {
                if (cc > 1) {
                    break;
                }
            } else {
                break;
            }
            const sizeAgo = input[i-(size-1)];
            map[sizeAgo - 'a'] -= 1;
        }
    }
    std.debug.print("count = {d}\n", .{i + 1});
}
