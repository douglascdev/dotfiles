# dotfiles
Dotfiles are pulled as submodules:
```
# Initialize
git submodule init
git submodule update

# Update
git submodule update --remote --recursive
```
## Build
Build the image
```
docker build -t devpod .
```
Reference it in a `.devcontainer.json`.
```
{
  "name": "hashbot",
  "image": "devpod",
  "remoteUser": "devuser"
}
```
