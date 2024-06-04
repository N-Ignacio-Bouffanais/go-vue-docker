# API and chat bot telegram with Golang

- A simple backend project where you can request to an api endpoint, data from another server via ssh. Also a chat bot build on top golang and telegram API.

## Config Docker

- To build the image on Docker we use the following command.

```bash
docker build -t vue-go-docker .
```

- To see all the images on our Docker we use the following command.

```bash
docker images
```

- To see all the containers on our Docker we use the following command.

```bash
docker ps -a
```

- To run our Docker container we use the following command.

```bash
docker run -p 8083:8083 vue-go-docker:2
```
