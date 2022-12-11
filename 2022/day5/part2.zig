const std = @import("std");
const print = std.debug.print;
const parseInt = std.fmt.parseInt;

const input = @embedFile("input.txt");

pub fn main() !void {
    var buffer: [2288]u8 = undefined;
    var fba = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = fba.allocator();

    var m = std.AutoArrayHashMap(usize, std.ArrayList(u8)).init(allocator);
    var lines = std.mem.tokenize(u8, input, "\n");
    while (lines.next()) |line| {
        if (!std.mem.startsWith(u8, line, "move")) {
            var i: usize = 0;
            while ((i * 4 + 1) < line.len) : (i += 1) {
                const c = line[i * 4 + 1];
                var list = m.get(i + 1) orelse std.ArrayList(u8).init(allocator);
                if (c >= 'A' and c <= 'Z') {
                    try list.insert(0, c);
                }
                try m.put(i + 1, list);
            }
            continue;
        }
        var parts = std.mem.tokenize(u8, line, " ");
        _ = parts.next().?;
        const count = try parseInt(usize, parts.next().?, 10);
        _ = parts.next().?;
        const src = try parseInt(u8, parts.next().?, 10);
        _ = parts.next().?;
        const dst = try parseInt(u8, parts.next().?, 10);

        var srcList = m.get(src).?;
        var dstList = m.get(dst) orelse std.ArrayList(u8).init(allocator);

        var srcStart: usize = srcList.items.len - count;
        var srcEnd: usize = srcList.items.len;
        try dstList.appendSlice(srcList.items[srcStart..srcEnd]);
        srcList.shrinkRetainingCapacity(srcList.items.len - count);

        try m.put(src, srcList);
        try m.put(dst, dstList);
    }

    var iter = m.iterator();
    while (iter.next()) |entry| {
        const v = entry.value_ptr;
        if (v.items.len > 0) {
            const last = v.items[v.items.len - 1];
            print("{c}", .{last});
        }
    }
}
