use hyper::{Body, Client, Request, Method, Uri};
use hyper::body::HttpBody;
use hyper::header::{USER_AGENT};
use tokio::fs::File;
use tokio::io::AsyncWriteExt;


pub async fn get_url(url: &str) -> Result<(), Box<dyn std::error::Error>> {
    println!("<=================--------------=================>");

    // Create a client for making requests
    let client = Client::new();

    // Prepare the request with the given URL and set the method to GET
    println!("<=================-------inseert agent in request-------=================>");
    let uri: Uri = url.parse()?;
    println!("Validated URI: {}", uri);
    let mut req = Request::builder()
        .method(Method::GET)
        .uri(uri)
        .header(
            USER_AGENT,
            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.37",
        )
        .body(Body::empty())?;
    
    println!("<=================------End inseert agent in request--------=================>");

    // Make the HTTP request
    let mut response = client.request(req).await?;

    // Log the status code of the response
    println!("Response: {:?}", response.status());

    // Create a file to save the downloaded content
    let mut file = File::create("downloaded_file.pdf").await?;

    // Iterate over the response body chunks and write them to the file
    while let Some(chunk) = response.body_mut().data().await {
        let chunk = chunk?; // Handle the result of the chunk
        file.write_all(&chunk).await?;
    }

    println!("Download completed successfully!");

    Ok(())
}
