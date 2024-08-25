# File Server

## Author

**Mark Albrand** (21004)

## Table of Contents

- [File Server](#file-server)
  - [Author](#author)
  - [Table of Contents](#table-of-contents)
  - [Project Description](#project-description)
  - [Dependencies](#dependencies)
    - [Go](#go)
  - [Functionalities](#functionalities)
    - [File Upload](#file-upload)
    - [File access](#file-access)
  - [How to run](#how-to-run)
    - [Environment Variables](#environment-variables)
  - [License](#license)

## Project Description

This project consists of a simple file server. Its main purpose is to aid in the file sharing functionalities of the [XMPP Chat Client Project](https://github.com/markalbrand56/Redes-Proyecto-1). This file server was shared with a close group of friends, and it is not intended to be used by the public. This allowed us to have a simple file sharing system and file preview system in our chat client projects.

The file server is written in Go and uses the `Gin Gonic` framework. It is a simple server that allows users to upload files and view them.

## Dependencies

The file server uses the following dependencies:

### Go

To install Go, follow the instructions on the [official website](https://golang.org/doc/install) and download the installer for your operating system. Make sure to get at least **Go 1.21**.

After installing Go, you can install the dependencies by running the following command:

```bash
go mod tidy
```

This will install all the dependencies needed to run the file server.

## Functionalities

### File Upload

The file server allows users to upload files. These file are uploaded to a `AWS S3 Bucket` and are served from there.

This files must be sent as part of a `multipart/form-data` post request. The file must be sent as a file with the key `files`. It also needs a `directory` key that specifies the directory where the file will be uploaded inside the bucket.

The file server will return a JSON object with the following structure:

```json
{
    "code": 200,
    "message": "File uploaded successfully",
    "paths": [
        "https://file-server-address/directory/file.png"
    ]
}
```

### File access

The file server allows users to access files. These files are served from a `AWS S3 Bucket`. The files are accessed by sending a `GET` request to the file server with the URL of the file, recieved from the file upload functionality.

## How to run

To run the file server, you must have Go installed in your computer. You can run the file server by running the following command:

```bash
go run cmd/main.go
```

### Environment Variables

Before running the file server, you must set the following environment variables:

- `AWS_ACCESS_KEY_ID`: The access key ID for the AWS S3 Bucket.
- `AWS_SECRET_ACCESS_KEY`: The secret access key for the AWS S3 Bucket.
- `AWS_REGION`: The region of the AWS S3 Bucket.
- `AWS_BUCKET`: The name of the AWS S3 Bucket.
- `URL`: The URL of the file server. This is used to generate the URLs of the files uploaded to the bucket.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
