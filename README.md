# Wasa Photo!

## How to build container images

### Backend

```sh
$ docker build -t wasa-photo-backend:latest -f Dockerfile.backend .
```

### Frontend

```sh
$ docker build -t wasa-photo-frontend:latest -f Dockerfile.frontend .
```

## How to run container images

### Backend

```sh
$ docker run -it --rm -v $(pwd)/service/database/source:/app/service/database/source -p 3000:3000 wasa-photo-backend:latest
```

### Frontend

```
$ docker run -it --rm -p 8080:80 wasa-photo-frontend:latest
```

## License

See [LICENSE](LICENSE).
