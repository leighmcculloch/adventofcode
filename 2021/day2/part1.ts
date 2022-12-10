const input = await Deno.readTextFile("input.txt");

let horizonal = 0;
let depth = 0;

const lines = input.split("\n");
for (const line of lines) {
    const [command, value_str] = line.split(" ");
    const value = parseInt(value_str);
    switch (command) {
        case "forward": horizonal += value; break;
        case "up":      depth -= value;     break;
        case "down":    depth += value;     break;
    }
}

console.log(`${horizonal * depth}`);
