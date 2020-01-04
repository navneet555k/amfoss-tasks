extern crate regex;
use regex::Regex;
use std::io;

fn main() {
    
    let mut email_id=String::new();

    let mut user:String = io::stdin().read_line(&mut email_id).unwrap().to_string();
    let re = Regex::new(r"^[A-Za-z0-9!#$%&'*+=?^_`{|}~.]+@[a-zA-z0-9]+\.[a-z]*").unwrap();
    let check = re.is_match(&mut email_id);

    if check{
        println!("valid email id")
    }
    else{

        println!("invalid email id")
    }
}
