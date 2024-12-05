
mod cmd;

use cmd::flag::CLI; 
fn main() {
    let cli = CLI::flags();

    println!("Link: {}", cli.link);
    println!("Flags: {:?}", cli.flags);
}
