const INPUT: &str = include_str!("../input.txt");

fn main() {
    let nums = INPUT
        .split('\n')
        .filter(|l| !str::is_empty(l))
        .map(str::parse::<u64>)
        .collect::<Result<Vec<_>, _>>()
        .unwrap();
    let mut last = *nums.first().unwrap();
    let mut increased = 0;
    for n in nums {
        if n > last {
            increased += 1;
        }
        last = n;
    }
    println!("{increased}");
}
