use std::collections::HashMap;
extern crate rayon;

use rayon::prelude::*;

////////// Array of instructions from lines in input files /////
type HashedInstructions<'a> = HashMap<&'a str, (&'a str, &'a str)>;

/////// Get the next mapped value /////////
fn get_next_key<'a>(
    items: &HashedInstructions<'a>,
    key: &'a str,
    dir: char,
) -> &'a str {
    let options: (&'a str, &'a str) = items[key];
    let next_key: &str = if dir == 'L' { options.0 } else { options.1 };

    next_key
}

////////// Read file from file and returns
/// an array of lines as Strings ////////////
fn get_input_from_file(path: &str) -> Vec<String> {
    return std::fs::read_to_string(path)
        .unwrap()
        .lines()
        .map(|ln| ln.to_string())
        .collect();
}

pub fn main() {
    /////// Get data from input file /////////
    let lines: Vec<String> = get_input_from_file("src/day8/sample.txt");

    ///////// Extract the direction pattern //////////
    let directions: Vec<char> = lines[0].chars().collect();

    println!("Directions {:?}\n", directions);

    let mut instructions: HashedInstructions = HashMap::new();

    //////// construct the hashmap /////////////
    for line in lines.iter().skip(1) {
        if let Some((key, value)) = line.split_once('=') {
            let trimmed_key = key.trim();

            let trimmed_value =
                value.trim().trim_matches('(').trim_matches(')');

            let direction_options =
                trimmed_value.split_once(", ").unwrap_or(("", ""));

            instructions.insert(
                trimmed_key,
                (direction_options.0, direction_options.1),
            );
        }
    }

    let mut starting_indices = vec![];
    println!("Instructions {:?}\n", instructions);

    //////////// Loop while target is not hit //////////////
    for (key, _) in instructions.iter() {
        let split_letters: Vec<char> = key.chars().collect();
        let last_letter = split_letters.last().unwrap();
        if *last_letter == 'A' {
            starting_indices.push(key);
        }
    }
    println!("Indexes key to start with {:?}\n", starting_indices);

    starting_indices.par_iter().for_each(|x| {
        let key = instructions.get(x.clone());
        println!("{:?}", key);
        let mut n = 0;
        let mut current = "AAA";
        loop {
            if current == "ZZZ" {
                break;
            } else {
                let direction: &char = &directions[n % directions.len()];
                current = get_next_key(&instructions, current, *direction);
                println!( "- steps {} : Let's go {} value: {}", n, direction, current);
                n += 1;
                println!("{:?}", current == "ZZZ")
            }
        }
    })
}
