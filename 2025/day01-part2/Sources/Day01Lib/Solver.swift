public enum Direction {
    case left
    case right
}

public struct Instruction {
    public let direction: Direction
    public let amount: Int
}

public struct Solver {
    public init() {}

    public func parse(input: String) -> [Instruction] {
        input.split(separator: "\n").compactMap { line -> Instruction? in
            guard !line.isEmpty else { return nil }
            let direction: Direction = line.first == "L" ? .left : .right
            let amount = Int(line.dropFirst())!
            return Instruction(direction: direction, amount: amount)
        }
    }

    public func solve(instructions: [Instruction]) -> Int {
        var position = 50
        var zeroCount = 0

        for instruction in instructions {
            let zeros: Int
            switch instruction.direction {
            case .left:
                if position > 0 {
                    zeros = (100 - position + instruction.amount) / 100
                } else {
                    zeros = instruction.amount / 100
                }
                position = (position - instruction.amount % 100 + 100) % 100
            case .right:
                zeros = (position + instruction.amount) / 100
                position = (position + instruction.amount) % 100
            }
            zeroCount += zeros
        }

        return zeroCount
    }
}
