import Foundation

let input = try! String(contentsOfFile: "INPUT.txt", encoding: .utf8)
let lines = input.split(separator: "\n").map { String($0) }

var position = 50
var zeroCount = 0

for line in lines {
    guard !line.isEmpty else { continue }

    let direction = line.first!
    let amount = Int(line.dropFirst())!

    if direction == "L" {
        position = (position - amount % 100 + 100) % 100
    } else { // R
        position = (position + amount) % 100
    }

    if position == 0 {
        zeroCount += 1
    }
}

print(zeroCount)
