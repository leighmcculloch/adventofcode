const std = @import("std");
const parseInt = std.fmt.parseInt;

const input = @embedFile("input.txt");

pub fn main() void {
    var elves = std.mem.split(u8, input, "\n\n");

    var most_total_calories: u64 = 0;

    while (elves.next()) |elf| {
        var foods = std.mem.split(u8, elf, "\n");
        var total_calories: u64 = 0;
        while (foods.next()) |food| {
            if (std.mem.eql(u8, food, "")) {
                continue;
            }
            const calories = parseInt(u64, food, 10) catch @panic("dun dun dun");
            total_calories += calories;
        }
        if (total_calories > most_total_calories) {
            most_total_calories = total_calories;
        }
    }

    std.debug.print("{d}\n", .{most_total_calories});
}
