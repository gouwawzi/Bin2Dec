use std::process;
use std::env;

mod lib;

fn main() {
    match lib::run(env::args()) {
        Ok(r) => println!("{}", r),
        Err(e) => {
            eprintln!("{}", e);
            process::exit(1);
        },
    };
}
