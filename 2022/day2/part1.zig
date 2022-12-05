const std = @import("std");

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = std.mem.tokenize(u8, input, "\n");

    var sum: u64 = 0;

    while (lines.next()) |line| {
        const play = Play{
            .opponent = try Shape.from_letter(line[0]),
            .player = try Shape.from_letter(line[2]),
        };
        sum += play.score();
    }

    std.debug.print("{any}\n", .{sum});
}

const Play = struct {
    opponent: Shape,
    player: Shape,

    pub fn score(p: Play) u8 {
        return p.player.score() + p.result().score();
    }

    pub fn result(p: Play) Result {
        return if (p.player == p.opponent)
            Result.Tie
        else if (p.player.trumps() == p.opponent)
            Result.Win
        else
            Result.Lose;
    }
};

const Shape = enum(u8) {
    Rock = 1,
    Paper = 2,
    Scissors = 3,

    const Error = error{UnrecognizedShape};

    pub fn from_letter(letter: u8) !Shape {
        return switch (letter) {
            'A', 'X' => Shape.Rock,
            'B', 'Y' => Shape.Paper,
            'C', 'Z' => Shape.Scissors,
            else => Error.UnrecognizedShape,
        };
    }

    pub fn trumps(s: Shape) Shape {
        return switch (s) {
            .Rock => .Scissors,
            .Paper => .Rock,
            .Scissors => .Paper,
        };
    }

    pub fn score(s: Shape) u8 {
        return @enumToInt(s);
    }
};

const Result = enum(u8) {
    Win = 6,
    Tie = 3,
    Lose = 0,

    pub fn score(r: Result) u8 {
        return @enumToInt(r);
    }
};
