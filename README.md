# Finkita

Financial Tracker App

### Prerequisite

- go ^1.15
- docker
- docker-compose

### How to run

run service

```bash
make run
```

run service in development mode (hot reload with `air`)

```bash
make dev-air
```

run test

```bash
make test
```

build service

```bash
make build
```

run scraper

```bash
make run-scraper
```

build scraper

```bash
make build-scraper
```

### Generating migration file

```bash
make name=[migration name]
```
