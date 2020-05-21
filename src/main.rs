use bin2dec::*;
use std::process;
use std::env;

fn main() {
    match run(env::args()) {
        Ok(r) => println!("{}", r),
        Err(e) => {
            eprintln!("{}", e);
            process::exit(1);
        },
    };
}
