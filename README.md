# go-terminal

A terminal-based development environment health checker. It verifies that the environment variables and hosts you list in a config file are available, then prints the results as a color-coded table.

## Features

- Check that environment variables are set and non-empty
- Check that hosts are reachable over TCP
- Results shown as a color-coded, aligned table in the terminal

## Usage

```sh
go run .
```

Or build a binary:

```sh
go build -o devchecks .
./devchecks
```

The tool exits with code `1` if any check fails.

## Configuration

Checks are defined in `devcheck.yaml`:

```yaml
envs:
  - PORT
  - DATABASE_URL

hosts:
  - google.com:80
  - localhost:5432
```

- `envs`: environment variables to verify are set and non-empty.
- `hosts`: host:port addresses to verify are reachable.

## Project structure

```
main.go         entrypoint
config/         loads and validates devcheck.yaml
checker/        runs env and host checks
ui/             prints the color-coded results table
```