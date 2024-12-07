use std::fs::File;
use std::io::{self, Read};

fn main() {
    let content = match read_input("input.txt") {
        Ok(contents) => contents,
        Err(e) => {
            eprintln!("Error reading file: {}", e);
            String::new() // Return an empty string in case of error
        }
    };

    // loop on each line
    // validate if line is safe or unsafe
    // safe, if all numbers are increasing or decreasing
    // safe, if numbers are the same and don't differ by more than 3 
    // count only safe lines
    let mut total_safe_lines = 0;

    for x in content.lines() {
        let mut numbers_on_line: Vec<i32> = Vec::new();
        for y in x.chars() {
            if y.is_numeric() {
                numbers_on_line.push(y as i32);
            }
        }
            for (index, value) in numbers_on_line.iter().enumerate() {
                if index == 0 {
                    continue;
                }

            // check descending
            if *value < numbers_on_line[index - 1] {
                let mut descending_check = true;
                println!("Value is descending");
                println!("{:?}", numbers_on_line);
                // check number size difference
                //TODO
                if  numbers_on_line[index - 1] - value > 3 {
                    break;
                }
            }
            // check acending
            if *value > numbers_on_line[index - 1] {
                let mut ascending_check = true;
                println!("Value is ascending");
                println!("{:?}", numbers_on_line);
                // check number size difference 
                //TODO
                if value - numbers_on_line[index - 1] > 3 {
                    break;
                }
            }
        } 
    }
}


fn read_input(filename: &str) -> io::Result<String> {
    let mut file = File::open(filename)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}