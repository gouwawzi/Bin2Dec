pub fn run(mut args: std::env::Args) -> Result<u32, &'static str> {
    args.next();

    match args.next() {
        Some(arg) => return convert(&arg),
        None => return Err("No input!"),
    }
}
pub fn convert(s: &str) -> Result<u32,  &'static str> {
    let mut result: u32 = 0;
    let mut w:u32 = 0;

    if s.len() > 32 {
        return Err("Too long input!");
    }

    for i in (0..s.len()).rev() {
            match &s[i..i+1].parse::<u32>() {
                Ok(r) => if *r > 1 {
                    return Err("Invalid binary number!");
                }else {
                    result += r*(1<<w);
                    w += 1;
                },
                Err(_) => return Err("Invalid binary number!"),
            };
    }
    Ok(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn convert_min() {
        assert_eq!(convert("0"), Ok(0));
    }

    #[test]
    fn convert_max() {
        assert_eq!(convert("11111111111111111111111111111111"), Ok(4294967295));
    }

    #[test]
    fn convert_cases() {
        assert_eq!(convert("111"), Ok(7));
        assert_eq!(convert("10101"), Ok(21));
        assert_eq!(convert("100000"), Ok(32));
    }

    #[test]
    fn convert_too_long() {
        assert_eq!(convert("100000000000000000000000000000000"), Err("Too long input!"));
    }

    #[test]
    fn convert_invalid() {
        assert_eq!(convert("abc"), Err("Invalid binary number!"));
        assert_eq!(convert("123"), Err("Invalid binary number!"));
    }
}
