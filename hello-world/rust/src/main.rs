use std::env;

fn main() {
  let args: Vec<String> = env::args().collect();

  if args.len() > 1 {
    println!("Hello, {}!", args[1]);
    return;
  }

  println!("Hello, World!");
}
