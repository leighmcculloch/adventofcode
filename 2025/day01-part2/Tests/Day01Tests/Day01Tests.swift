import Day01Lib
import Foundation

// Simple test harness for environments without XCTest
nonisolated(unsafe) var passed = 0
nonisolated(unsafe) var failed = 0

func expect<T: Equatable>(_ actual: T, equals expected: T, _ message: String) {
    if actual == expected {
        print("✓ \(message)")
        passed += 1
    } else {
        print("✗ \(message)")
        print("  Expected: \(expected)")
        print("  Actual:   \(actual)")
        failed += 1
    }
}

func summary() {
    print("\n\(passed) passed, \(failed) failed")
    if failed > 0 {
        exit(1)
    }
}

let solver = Solver()

@MainActor func solve(_ input: String) -> Int {
    solver.solve(instructions: solver.parse(input: input))
}

// Test: Example from problem description
expect(
    solve("""
        L68
        L30
        R48
        L5
        R60
        L55
        L1
        L99
        R14
        L82
        """),
    equals: 3,
    "Example from problem description"
)

// Test: Single rotation landing on zero
expect(
    solve("L50"),
    equals: 1,
    "Single rotation to zero (L50 from 50)"
)

// Test: Single rotation not landing on zero
expect(
    solve("L10"),
    equals: 0,
    "Single rotation not zero (L10 from 50 = 40)"
)

// Test: Wrap around left
expect(
    solve("L60"),
    equals: 0,
    "Wrap around left (L60 from 50 = 90)"
)

// Test: Wrap around right to zero
expect(
    solve("R50"),
    equals: 1,
    "Wrap around right to zero (R50 from 50 = 0)"
)

// Test: Multiple zeros
expect(
    solve("""
        R50
        R100
        L100
        """),
    equals: 3,
    "Multiple zeros counted"
)

// Test: Large rotation values
expect(
    solve("R950"),
    equals: 1,
    "Large rotation (R950 from 50 = 0)"
)

summary()
