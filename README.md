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
cd project-folder
mkdir -p .devcontainer
cp -R devcontainer-feature .devcontainer
mv devcontainer.json .devcontainer
devcontainer up --workspace-folder ./ --remove-existing-container
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
