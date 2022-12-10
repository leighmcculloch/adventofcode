const INPUT: &str = include_str!("../input.txt");

fn main() {
    let nums = INPUT
        .split('\n')
        .filter(|l| !str::is_empty(l))
        .map(str::parse::<u64>)
        .collect::<Result<Vec<_>, _>>()
        .unwrap();
    let mut nums = nums.iter();
    let mut last3 = nums.next().unwrap();
    let mut last2 = nums.next().unwrap();
    let mut last1 = nums.next().unwrap();
    let mut lastsum = last3 + last2 + last1;
    let mut increased = 0;
    for n in nums {
        let nsum = n+last1+last2;
        if nsum > lastsum {
            increased += 1;
        }
        last3 = last2;
        last2 = last1;
        last1 = n;
        lastsum = nsum;
    }
    println!("{increased}");
}
