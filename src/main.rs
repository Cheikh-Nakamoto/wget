mod cmd;
mod internal {
    pub mod downloader;
}

use cmd::flag::CLI;
use internal::downloader::download::get_url;
use tokio;

#[tokio::main]  // Marque la fonction main comme étant asynchrone
async fn main() {
    let cli = CLI::flags();
    println!("Link: {}", cli.link);
    println!("Flags: {:?}", cli.flags);

    // Attendez la fin de l'exécution de la fonction get_url
    if let Err(e) = get_url(&cli.link.clone()).await {
        eprintln!("Error: {}", e);
    }
}
