const std = @import("std");
const parseInt = std.fmt.parseInt;

const input = @embedFile("input.txt");

pub fn main() void {
    var elves = std.mem.split(u8, input, "\n\n");

    var totals = Totals(3, u64){};

    while (elves.next()) |elf| {
        var foods = std.mem.split(u8, elf, "\n");
        var total_calories: u64 = 0;
        while (foods.next()) |food| {
            if (std.mem.eql(u8, food, "")) {
                continue;
            }
            const calories = parseInt(u64, food, 10) catch unreachable;
            total_calories += calories;
        }

        totals.add(total_calories);
    }

    std.debug.print("{d}\n", .{totals.sum()});
}

fn Totals(comptime N: usize, comptime T: type) type {
    return struct {
        const This = @This();

        totals: [N]T = [_]T{0} ** N,

        pub fn add(this: *This, total: T) void {
            var cell: ?*T = null;
            for (this.totals) |*t| {
                if (total <= t.*) {
                    break;
                }
                if (cell) |c| {
                    c.* = t.*;
                }
                cell = t;
            }
            if (cell) |c| {
                c.* = total;
            }
        }

        pub fn sum(this: *This) T {
            var total: T = 0;
            for (this.totals) |t| {
                total += t;
            }
            return total;
        }
    };
}
