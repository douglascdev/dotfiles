# dotfiles
Dotfiles are pulled as submodules:
```
# Initialize
git submodule init
git submodule update

# Update
git submodule update --remote --recursive
```
Make sure to pass your current user's UID/GID in .devcontainer.json
```
{
 "build": {
   "dockerfile": "~/Dockerfile",
   "args": {
     "USER_UID": "1000",
     "USER_GID": "1000"
   }
 },
 "remoteUser": "devuser"
}
```
