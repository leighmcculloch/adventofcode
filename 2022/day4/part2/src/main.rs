const INPUT: &str = include_str!("../../input.txt");

fn main() {
    let count = INPUT
        .split('\n')
        .filter(|line| !line.is_empty())
        .filter(|pair| {
            let range: [_; 2] = pair
                .split(',')
                .map(|section| {
                    let section: [_; 2] = section
                        .split('-')
                        .map(str::parse::<u64>)
                        .collect::<Result<Vec<_>, _>>()
                        .unwrap()
                        .try_into()
                        .unwrap();
                    section[0]..=section[1]
                })
                .collect::<Vec<_>>()
                .try_into()
                .unwrap();
            range[0].contains(range[1].start())
                || range[0].contains(range[1].end())
                || range[1].contains(range[0].start())
                || range[1].contains(range[0].end())
        })
        .count();
    println!("{}", count);
}
