#!/bin/bash
# set up dot files in devpod: https://devpod.sh/docs/developing-in-workspaces/dotfiles-in-a-workspace

cd ~
mv dotfiles/* .
mv dotfiles/.* .

git submodule init
git submodule update
