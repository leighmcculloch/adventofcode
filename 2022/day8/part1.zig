const std = @import("std");
const print = std.debug.print;
const tok = std.mem.tokenize;
const ArrayList = std.ArrayList;
const expectEqual = std.testing.expectEqual;

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = tok(u8, input, "\n");

    const index = Index{
        .width = lines.next().?.len,
        .height = 1 + count(&lines),
    };

    var map = [_]u8{0} ** input.len;
    map[index.i(w, h)]

    var maxFromLeft = [_]u8{0} ** input.len;

    var h: usize = 0;
    while (h < index.height) : (h += 1) {
        var w: usize = 0;
        var last: u8 = 0;
        while (w < index.width) : (w += 1) {
            var item: u8 = 
            maxFromLeft[index.i(w, h)] = 0;
        }
    }

    print("{any}", .{maxFromLeft});
}

fn count(iter: anytype) usize {
    var c: usize = 0;
    while (iter.next()) |_| {
        c += 1;
    }
    return c;
}

const Index = struct {
    const Self = @This();
    width: usize,
    height: usize,
    pub fn i(self: Self, w: usize, h: usize) usize {
        return ((h % self.height) * self.width) + (w % self.width);
    }
};

test "index" {
    {
        const index = Index{ .width = 1, .height = 5 };
        try expectEqual(@as(usize, 0), index.i(0, 0));
    }
    {
        const index = Index{ .width = 5, .height = 5 };
        try expectEqual(@as(usize, 4), index.i(4, 0));
        try expectEqual(@as(usize, 9), index.i(4, 1));
        try expectEqual(@as(usize, 24), index.i(4, 4));
    }
}
