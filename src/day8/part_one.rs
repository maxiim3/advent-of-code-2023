use std::collections::HashMap;

type HashedInstructions<'a> = HashMap<&'a str, (&'a str, &'a str)>;

fn get_next_key<'a>(items: &HashedInstructions<'a>, key: &'a str, dir: char) -> &'a str {
    let options: (&'a str, &'a str) = items[key];
    let next_key: &str = if dir == 'L' { options.0 } else { options.1 };

    next_key
}

fn get_input_from_file(path: &str) -> Vec<String> {
    return std::fs::read_to_string(path)
        .unwrap()
        .lines()
        .map(|ln| ln.to_string())
        .collect();
}

pub fn part_one() {
    let lines: Vec<String> = get_input_from_file("src/day8/input.txt");

    let directions: Vec<char> = lines[0].chars().collect();

    println!("{:?}", directions);

    let mut instructions: HashedInstructions = HashMap::new();

    for line in lines.iter().skip(1) {
        if let Some((key, value)) = line.split_once('=') {
            let trimmed_key = key.trim();

            let trimmed_value = value.trim().trim_matches('(').trim_matches(')');

            let direction_options = trimmed_value.split_once(", ").unwrap_or(("", ""));

            instructions.insert(trimmed_key, (direction_options.0, direction_options.1));
        }
    }

    let mut current: &str = "AAA";
    let mut n = 0;
    loop {
        if current == "ZZZ" {
            break;
        } else {
            let direction: &char = &directions[n % directions.len()];
            current = get_next_key(&instructions, current, *direction);
            println!("- steps {} : Let's go {} value: {}", n, direction, current);
            n += 1;

            println!("{:?}", current == "ZZZ")
        }
    }
    println!("steps {}", n);
}

