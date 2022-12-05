const std = @import("std");

const input = @embedFile("input.txt");

pub fn main() !void {
    var lines = std.mem.tokenize(u8, input, "\n");

    var sum: u64 = 0;

    while (lines.next()) |line| {
        const play = Play{
            .opponent = try Shape.from_letter(line[0]),
            .result = try Result.from_letter(line[2]),
        };
        sum += play.score();
    }

    std.debug.print("{any}\n", .{sum});
}

const Play = struct {
    opponent: Shape,
    result: Result,

    pub fn score(p: Play) u8 {
        return p.player().score() + p.result.score();
    }

    pub fn player(p: Play) Shape {
        return switch (p.result) {
            .Tie => p.opponent,
            .Win => p.opponent.trumped_by(),
            .Lose => p.opponent.trumps(),
        };
    }
};

const Shape = enum(u8) {
    Rock = 1,
    Paper = 2,
    Scissors = 3,

    const Error = error{UnrecognizedShape};

    pub fn from_letter(letter: u8) !Shape {
        return switch (letter) {
            'A' => Shape.Rock,
            'B' => Shape.Paper,
            'C' => Shape.Scissors,
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

    pub fn trumped_by(s: Shape) Shape {
        return switch (s) {
            .Rock => .Paper,
            .Paper => .Scissors,
            .Scissors => .Rock,
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

    pub fn from_letter(letter: u8) !Result {
        return switch (letter) {
            'X' => Result.Lose,
            'Y' => Result.Tie,
            'Z' => Result.Win,
            else => (error{UnrecognizedShape}).UnrecognizedShape,
        };
    }

    pub fn score(r: Result) u8 {
        return @enumToInt(r);
    }
};
