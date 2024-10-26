# gnsagrent

**gnsagrent** is a lightweight, continuously running service designed to keep DNS records updated with the current public IPv4 address of the client it operates on. Ideal for users with dynamic IP addresses on their home networks, **gnsagrent** ensures that DNS records are always accurate, providing reliable access to networked resources even when IPs change.


## Features

- Supports dynamic DNS updates across multiple providers.
- Flexible configuration through `config.json`.
- Runs in a Docker container for easy deployment.

## Configuration

The configuration for **gnsagrent** is managed via a `config.json` file, which contains two main sections: `endPoints` and `domains`.

### Example `config.json`

```json
{
    "endPoints": [
        {
            "url": "https://api.ipify.org/?format=json",
            "property": "ip"
        }
    ],
    "domains": [
        {
            "type": "CLOUDFLARE",
            "email": "your.email@example.com",
            "domain": "yourdomain.com",
            "subdomain": "yourSubdomain",
            "apiKey": "yourApiKey"
        }
    ]
}
```

#### Configuration Details

- **endPoints**: Defines the API endpoints to fetch the public IP. Each entry should include:
  - **url**: The API URL to retrieve the IP.
  - **property**: The JSON property containing the IP address in the response.

- **domains**: Specifies the domains and providers where DNS records should be updated.
  - **type**: The DNS provider type (e.g., `CLOUDFLARE`).
  - **email**: The account email for the DNS provider.
  - **domain**: The primary domain name.
  - **subdomain**: *(Optional)* A specific subdomain for DNS updates.
  - **apiKey**: Your API key for authenticating with the DNS provider.

### Notes

- **subdomain** is optional; if omitted, updates will be applied to the main domain.
- Replace `"your.email@example.com"`, `"yourdomain.com"`, `"yourSubdomain"`, and `"yourApiKey"` with your actual credentials and domain details.

## Usage

To run **gnsagrent** in a Docker container, follow these steps:

1. **Build the Docker image**:
    ```bash
    docker build -t gnsagrent-image .
    ```

2. **Run the Docker container**, mounting your `config.json` file:
    ```bash
    docker run -d --name gnsagrent-container -v /path/to/your/config.json:/app/config.json gnsagrent-image
    ```

Replace `/path/to/your/config.json` with the actual path to your configuration file.

## Contributing

Contributions are welcome! Please open issues for bugs or feature requests and submit pull requests for code contributions.

## License

This project is licensed under the MIT License.

--- 

This README gives an overview of **gnsagrent**, including usage instructions and an example configuration, which you can customize to suit your own setup.