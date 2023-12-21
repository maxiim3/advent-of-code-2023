use std::fs;

pub fn run() {
    println!("Hello day 4");

    let file = fs::read_to_string("./src/day4/sample.txt").unwrap();

    file.lines()
        .into_iter()
        .for_each(|card| {
            if let Some((_, numbers)) = card.split_once(":") {
                if let Some((winning_data, my_numbers_data)) = numbers.split_once("|") {
                    convert_data_into_vec(winning_data)
                }
            }
        })
}

fn convert_data_into_vec(data: &str) {
    let Some(splitted_data) = data.split(" ").collect();
    dbg!(splitted_data)
}
