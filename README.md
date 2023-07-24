# cutcast

## Overview

Cutcast is an API application written in Go and backed by ffmpeg that enables real-time recording and cutting of live streams from YouTube. With cutcast, you can capture and create real-time cuts from YouTube streams, making it a powerful tool for various applications such as content creation, live event coverage, and more.

## Prerequisites

Before running the cutcast application, make sure you have the following dependencies installed on your system:

1. [Go](https://golang.org/): The Go programming language is required to build and run the application.
2. [ffmpeg](https://ffmpeg.org/): The application uses ffmpeg as a backend to handle stream recording and cutting.

## How to Run

To run the cutcast application, follow the steps below:

1. Clone the repository:

```bash
git clone https://github.com/your-username/cutcast.git
cd cutcast
```

2. Run the application:

```bash
go run main.go
```

The API server will start, and you can now interact with it to record and cut YouTube streams in real-time.

## Usage

Once the cutcast API server is running, you can make HTTP requests to it to record and create cuts from YouTube streams. Here are the available endpoints:

1. `POST /record`: Start recording a live stream from a YouTube URL. The request body should contain the YouTube URL.

2. `POST /record/:id/trim`: Create a cut from a currently recorded stream. The request body should contain the timestamp or time range for the cut.

## Examples

1. Record a YouTube stream:

```bash
curl -X POST 'http://localhost:8080/v1/record' -H "Content-Type: application/json"  -d '{"url":"https://www.youtube.com/watch?v=jfKfPfyJRdk"}'
```

2. Create a cut from a recorded stream:

```bash
# Assuming the start and end times are in seconds
curl -X POST 'http://localhost:8080/v1/record/6c4d1c0b-1f6a-47d2-a28c-12a4c539b0d3/trim' -H "Content-Type: application/json"  -d '{"start_time":"00:00:00", "end_time":"00:00:10"}'
```

## Tests

To execute tests for the cutcast application, use the following command:

```bash
go test -v ./...
```

This will run all the tests in the project and provide detailed output.

## Contributing

We welcome contributions to improve cutcast. If you find any issues or have ideas to enhance the functionality, feel free to open an issue or submit a pull request. Please make sure to follow the established coding guidelines and write tests for any new features.

## License

This project is licensed under the [MIT License](LICENSE).

---

Thank you for using cutcast! We hope this application serves your needs for recording and cutting YouTube streams in real-time. If you have any questions or need further assistance, please don't hesitate to contact us. Happy streaming!