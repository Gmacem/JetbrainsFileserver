# JetbrainsFileserver

Test task for Jetbrains internship

## Run locally

```bash
PORT=5050 go run .
```

## Build docker image

```bash
docker build -t gmacem/JetbrainsFileserver:0.0.1 .
```

## Run docker container

```bash
docker run -p 5050:5050 gmacem/JetbrainsFileserver:0.0.1
```

## Usage examples

### Create and write file

Create and write data in a file.

```bash
curl -XPOST "http://localhost:5050/path/to/file/some_filename" --data 'some message'
```

If file exists, then data will be overwritten.

If some directory does not exist, then directory will be created (equal to `mkdir -p /path/to/file`)

### Get file

```bash
curl -XGET "http://localhost:5050/path/to/file/some_filename"
```

### Delete file

```bash
curl -XDELETE "http://localhost:5050/path/to/file/some_filename"
```
