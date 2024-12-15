## Start postgres db as a Docker container

```shell
podman machine start
docker rm some-postgres
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```