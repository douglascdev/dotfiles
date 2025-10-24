FROM fedora:42 AS dev-container

# nvim stuff
RUN dnf install -y \
        gcc\
        cascadia-mono-nf-fonts\
        git\
        make\ 
        npm\ 
        nvim\ 
        ripgrep\ 
        shadow-utils\ 
        unzip\ 
        jetbrains-mono-fonts
# Go stuff
RUN dnf install -y \
        golang
# C++ stuff
RUN dnf install -y \
        glfw\
        glfw-devel\
        clang\
        clang-tools-extra && \
dnf copr enable -y che/nerd-fonts && dnf install -y nerd-fonts && \
dnf clean all

ARG USERNAME=dev
ARG UID=1000
ARG GID=1000

RUN groupadd --gid ${GID} ${USERNAME} || true && \
    useradd --system --uid ${UID} --gid ${GID} --create-home --shell /bin/bash ${USERNAME} || true &&\
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers

USER ${USERNAME}
WORKDIR /home/${USERNAME}
RUN \
        go install github.com/air-verse/air@v1.62.0 && \
        go install github.com/go-delve/delve/cmd/dlv@v1.7.3
RUN echo PATH="$PATH:$HOME/go/bin" > /home/${USERNAME}/.bashrc
RUN git clone --branch mine https://github.com/douglascdev/kickstart.nvim .config/nvim

WORKDIR /usr/src

EXPOSE 8080
