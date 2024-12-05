pub use std::env;
use regex::Regex;


pub struct CLI {
   pub link: String,
   pub  flags: Vec<String>,
}

impl CLI {
    // Méthode pour récupérer et valider les flags et l'URL
    pub fn flags() -> Self {
        let args: Vec<String> = env::args().collect();

        // Assurez-vous qu'il y a au moins un argument (l'URL)
        if args.len() < 2 {
            panic!("No link provided");
        }

        let link = args[args.len() - 1].clone();
        let _= Self::check_link(link.clone());
        let flags = args
            .iter()
            .filter(|&arg| arg.starts_with("-"))
            .cloned()
            .collect::<Vec<String>>();

        // Vérifie si les flags sont valides
        if !Self::valid_flag(&flags) {
            panic!("Invalid flags provided");
        }

        CLI { link, flags }
    }

    // Vérifie si les flags sont valides
    pub fn valid_flag(args: &[String]) -> bool {
        let valid_flags: Vec<String> = vec![
            "-B".to_string(), // background
            "-O".to_string(), // output file
            "-P".to_string(), // directory prefix
            "--rate-limit".to_string(), // limit download speed
            "-i".to_string(), // input file (multiple downloads)
            "--mirror".to_string(), // mirror website
            "-R".to_string(), // reject extensions
            "-X".to_string(), // exclude directories
            "--convert-links".to_string() // convert links for offline use
        ];

        // Vérifie si tous les flags sont valides
        args.iter().all(|flag| valid_flags.contains(flag))
    }

    pub fn check_link(link:String) -> bool {
        let url_regex = Regex::new(r"^(https?://)?(www\.)?[a-zA-Z0-9-]+(\.[a-zA-Z]{2,})+(\/[^\s]*)?$").unwrap();
        if !url_regex.is_match(&link) {
            panic!("Invalid URL");
        }
        true
    }
}
