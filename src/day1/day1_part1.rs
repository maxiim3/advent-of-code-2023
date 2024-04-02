use std::collections::HashMap;

use regex::bytes::Regex;

pub fn exec() {

    // A Dictionary / Object
    let dictionary: HashMap<&str, u8> = HashMap::from([
        ("1", 1),
        ("one", 1),
        ("2", 2),
        ("two", 2),
        ("3", 3),
        ("three", 3),
        ("4", 4),
        ("four", 4),
        ("5", 5),
        ("five", 5),
        ("6", 6),
        ("six", 6),
        ("7", 7),
        ("seven", 7),
        ("8", 8),
        ("eight", 8),
        ("9", 9),
        ("nine", 9),
    ]);

    // Regex Pattern
    let reg_exp =
        Regex::new(r"1|one|2|two|3|three|4|four|5|five|6|six|7|seven|8|eight|9|nine").unwrap();

    // Read the sample_file
    let file = std::fs::read_to_string("./day1/input.txt").unwrap();

    // Iterate over the lines
    let results = file
        // read each line -> JS .split("\n")
        .lines()
        .into_iter()
        .map(|line| {
            // the output array of numbers
            let mut list_of_matches = Vec::new();

            // the start index of the search
            let mut start = 0;

            // iterate over the matches - Some() destrures the Option that regex.find_at returns
            while let Some(matched_values) = reg_exp.find_at(line.as_bytes(), start) {
                let matched_str = &line[matched_values.range()]; // extract the result from range into string
                // lookup the key in the dictionary
                if let Some(&number) = dictionary.get(matched_str) {
                    list_of_matches.push(number);
                }
                // update the start index (next search)
                start = matched_values.start() + 1
            }
            // combine the first and last number into a string
            let combined = format!("{}{}", list_of_matches[0], list_of_matches[list_of_matches.len() - 1]);
            // parse the string into a number
            let digit = combined.parse::<usize>().unwrap();
            // return the digit variable
            digit
        }).reduce(|acc, current| acc + current).unwrap();
    println!("{}", results);
}

