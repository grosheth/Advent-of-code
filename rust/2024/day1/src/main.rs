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

    // Loop through the content and make a left and right list 
    let mut left_list: Vec<i32> = Vec::new();
    let mut right_list: Vec<i32> = Vec::new();
    let mut index = 0;
    for x in content.lines() {
        for y in x.split_whitespace() {
            if let Ok(value) = y.parse::<i32>() {
                if index % 2 == 0 {
                    left_list.push(value);
                } else {
                    right_list.push(value);
                }
                index += 1;
            }
        }
    }

    // Sort the lists
    left_list.sort();
    right_list.sort();
    let mut total = 0;

    // figure out how far apart they are for each pair
    // add all results
    for (i, x) in left_list.iter().enumerate() {
        let left_digit = *x;
        let right_digit = right_list[i];
        if left_digit > right_digit {
            total += left_digit - right_digit;
            println!("{}: {} > {}", i, left_digit, right_digit);
        } else {
            total += right_digit - left_digit;
            println!("{}: {} < {}", i, left_digit, right_digit);
        }
    }

    println!("Total: {}", total);
     
    // part 2
    let mut total2 = 0;

    for x in left_list.iter() {
        let count_in_right_list = count_occurrences(&right_list, *x);
        println!("{} is in left list {} times", x, count_in_right_list);
        total2 += x * count_in_right_list;
    } 
    println!("Total2: {}", total2);
}

fn count_occurrences(list: &Vec<i32>, number: i32) -> i32 {
    list.iter().filter(|&&x| x == number).count().try_into().unwrap()
}

fn read_input(filename: &str) -> io::Result<String> {
    let mut file = File::open(filename)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}
