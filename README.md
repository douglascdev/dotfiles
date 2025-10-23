# dotfiles

Dotfiles are pulled as submodules:

```bash
# Initialize
git submodule init
git submodule update

# Update
git submodule update --remote --recursive
```

## Build

Build the image:

```bash
podman build -t devgo -f Dockerfile.dev
```

Make an `.air.toml` at the root of the project:

```toml
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/main ."
full_bin = "./tmp/main"
include_ext = ["go"]
exclude_dir = [
  "frontend",
]
```

If needed, a `docker-compose.yml` as well:

```yaml
services:
  go-app:
    build:
      dockerfile: Dockerfile.dev
      context: .
    restart: always
    ports:
      - "8080:8080"
    volumes:
      - ..:/usr/src:Z
```
