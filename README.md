# Simple Go HTTP Server

This is a simple HTTP server written in Go that demonstrates handling basic HTTP requests and serving static files. It also showcases a custom request parser and basic response handling.

## Features

- Echo endpoint that outputs the content of the request path.
- User-agent endpoint that returns the User-Agent header from the request.
- Static file serving and handling POST requests to upload files.
- Custom request parsing and response generation.

## Prerequisites

Make sure you have Go installed on your computer. You can download it from [Go's official website](https://golang.org/dl/).

## Getting Started

To get the server running on your local machine, follow the steps below:

1. Clone the repository:
```bash
git clone https://github.com/your-username/simple-go-http-server.git cd simple-go-http-server
```

2. Run the server:
```bash
go run main.go [path-to-directory-for-static-files]
```

Replace `[path-to-directory-for-static-files]` with the directory where your static files are located. This is required for the file server to function properly.

3. Visit `http://localhost:4221` in your web browser or use a tool like `curl` to interact with the server.

## Endpoints

The following endpoints are available:

- `GET /`: Returns a simple welcome message.
- `GET /echo/{content}`: Echos back the `{content}` provided in the path.
- `GET /user-agent`: Returns the User-Agent header from the request.
- `GET /files/{filename}`: Serves a file from the specified static file directory.
- `POST /files/{filename}`: Uploads a file to the specified static file directory.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Enjoy your simple Go HTTP server!
