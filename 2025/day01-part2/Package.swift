// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "Day01",
    targets: [
        .target(
            name: "Day01Lib"
        ),
        .executableTarget(
            name: "Day01",
            dependencies: ["Day01Lib"]
        ),
        .executableTarget(
            name: "Day01Tests",
            dependencies: ["Day01Lib"],
            path: "Tests/Day01Tests"
        ),
    ]
)
