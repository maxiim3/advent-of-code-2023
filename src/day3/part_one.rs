pub fn exec() {
    let mut matrix = Vec::new();

    let file = std::fs::read_to_string("src/day3/sample.txt").unwrap();
    file
        .lines()
        .into_iter()
        .for_each(|ln| {
            let row: Vec<_> = ln
                .split("")
                .filter(|char| char.len() > 0 )
                .collect();

            matrix.push(row)
        });

    println!("{:?}", matrix);
}
