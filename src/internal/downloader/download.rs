use hyper::{Client, Uri};
use tokio::fs::File;
use tokio::io::AsyncWriteExt;
use hyper::body::HttpBody;
use std::convert::TryFrom;

pub async fn get_url(url: &str) -> Result<(), Box<dyn std::error::Error>> {
    // println!("Link passed : {}", link);
    let link: &'static str  = "https://github.com/Syntonie/documentation/blob/master/Solaire/Solaire_User_Guide.pdf";
    //Utilisation de try_from pour une gestion plus détaillée des erreurs
    // let uri = link.parse::<Uri>().unwrap();
    
    println!("<=================--------------=================>");
    let client = Client::new();
    let mut response = client.get(Uri::from_static(link)).await?;
    println!("{:?}", response);

    let mut file = File::create("downloaded_file.txt").await?;

    while let Some(chunk) = response.body_mut().data().await {
        file.write_all(&chunk?).await?;
    }

    Ok(())
}
