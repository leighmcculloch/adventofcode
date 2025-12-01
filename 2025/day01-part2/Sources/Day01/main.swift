import Foundation
import Day01Lib

let input = try! String(contentsOfFile: "INPUT.txt", encoding: .utf8)
let solver = Solver()
let instructions = solver.parse(input: input)
let result = solver.solve(instructions: instructions)
print(result)
