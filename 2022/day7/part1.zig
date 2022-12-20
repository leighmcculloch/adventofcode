const std = @import("std");
const Allocator = std.mem.Allocator;
const ArrayList = std.ArrayList;
const eql = std.mem.eql;
const tok = std.mem.tokenize;
const print = std.debug.print;
const parseInt = std.fmt.parseInt;

const input = @embedFile("input.txt");

pub fn main() !void {
    var buf: [51344]u8 = undefined;
    var fba = std.heap.FixedBufferAllocator.init(&buf);
    var allocator = fba.allocator();

    var root = try Node.init(allocator, null, "/", .dir, 0);
    var curr: *Node = root;

    var lines = tok(u8, input, "\n");
    while (lines.next()) |line| {
        // Ignore empty lines.
        if (line.len == 0) {
            continue;
        }
        // Commands start with $.
        if (line[0] == '$') {
            var args = tok(u8, line, " ");
            _ = args.next().?;
            const cmd = args.next().?;
            // Change directory.
            if (eql(u8, cmd, "cd")) {
                const dir = args.next().?;
                if (eql(u8, dir, "..")) {
                    curr = curr.parent.?;
                } else {
                    curr = curr.find(dir) orelse try curr.add(dir, .dir, 0);
                }
            }
            // List files.
            if (eql(u8, cmd, "ls")) {}
        } else {
            // File listings for the current directory.
            var attrs = tok(u8, line, " ");
            const sizeOrDir = attrs.next().?;
            const name = attrs.next().?;
            if (eql(u8, sizeOrDir, "dir")) {
                // Dir.
                _ = try curr.add(name, .dir, 0);
            } else {
                // File.
                const size = try parseInt(usize, sizeOrDir, 10);
                _ = try curr.add(name, .file, size);
            }
        }
    }

    // List all directories greater than 100000 in size.
    var sum: usize = 0;
    root.visit(&sum, sumTotalSize);
    print("{d}\n", .{sum});
}

fn sumTotalSize(ctx: *usize, n: *Node) void {
    if (n.typ == .dir) {
        const s = n.totalSize();
        if (s <= 100000) {
            ctx.* += s;
            print("{s}: {d}\n", .{ n.name, n.totalSize() });
        }
    }
}

const NodeType = enum {
    dir,
    file,
};

const Node = struct {
    const Self = @This();

    parent: ?*Self,
    name: []const u8,
    typ: NodeType,
    size: usize,
    nodes: ArrayList(*Self),

    pub fn init(allocator: Allocator, parent: ?*Self, name: []const u8, typ: NodeType, size: usize) !*Self {
        var self = try allocator.create(Self);
        self.* = .{
            .parent = parent,
            .name = name,
            .nodes = ArrayList(*Self).init(allocator),
            .typ = typ,
            .size = size,
        };
        return self;
    }

    pub fn add(self: *Self, name: []const u8, typ: NodeType, size: usize) !*Self {
        const n = try Self.init(self.nodes.allocator, self, name, typ, size);
        try self.nodes.append(n);
        return n;
    }

    pub fn find(self: Self, name: []const u8) ?*Self {
        for (self.nodes.items) |n| {
            if (eql(u8, n.name, name)) {
                return n;
            }
        }
        return null;
    }

    pub fn totalSize(self: Self) usize {
        var tot: usize = self.size;
        for (self.nodes.items) |n| {
            tot += n.totalSize();
        }
        return tot;
    }

    pub fn visit(self: Self, ctx: anytype, f: *const fn (ctx2: @TypeOf(ctx), n: *Node) void) void {
        for (self.nodes.items) |n| {
            if (n.typ == .dir) {
                f(ctx, n);
                n.visit(ctx, f);
            }
        }
    }
};
