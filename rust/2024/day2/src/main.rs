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

    let mut total_safe_lines = 0;

    for x in content.lines() {
        let mut numbers_on_line: Vec<i32> = Vec::new();
        for y in x.chars() {
            if y.is_numeric() {
                numbers_on_line.push(y.to_digit(10).unwrap() as i32);
            }
        }
            for (index, value) in numbers_on_line.iter().enumerate() {
                if index == 0 { 
                    if *value < numbers_on_line[index + 1] {
                        let mut ascending_check = true;
                    } else {
                        let mut ascending_check = false;
                    }
                    continue;
                }

                if ascending_check {
                    println!("Value is descending, {:?}", numbers_on_line);
                    if *value > numbers_on_line[index + 1] {
                        // check number size difference
                        // Number before should be greater than current number
                        // Number after should be smaller than current number
                        //TODO
                        let diff = numbers_on_line[index - 1] - value;
                        if  diff > 3 {
                            println!("test");
                            break;
                        }
                        if index == 0 || index < numbers_on_line.len() - 1 {
                            if *value > numbers_on_line[index - 1] {
                                println!("test");
                                break;
                            }
                            if *value < numbers_on_line[index + 1] {
                                println!("test");
                                break;
                            }
                        }
                    }
                } else {
                    println!("Value is ascending, {:?}", numbers_on_line);
                    if *value > numbers_on_line[index + 1] {
                        // Number before should be greater than current number
                        // Number after should be smaller than current number
                        //TODO
                        let  diff = numbers_on_line[index - 1] - value;
                        if  diff > 3 {
                            println!("test");
                            break;
                        }
                        if *value > numbers_on_line[index + 1] {
                            println!("test");
                            break;
                        }
                        if *value < numbers_on_line[index - 1] {
                            println!("test");
                            break;
                        }
                    }  
                }
            }
        total_safe_lines += 1; 
    }
    println!("Total safe lines: {}", total_safe_lines);
}


fn read_input(filename: &str) -> io::Result<String> {
    let mut file = File::open(filename)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}