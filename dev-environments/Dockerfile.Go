FROM golang:1.24 AS downloader
WORKDIR /app
RUN \
go install github.com/air-verse/air@v1.62.0 && \
go install github.com/go-delve/delve/cmd/dlv@v1.7.3

FROM fedora:42 AS dev-container

RUN dnf install -y \
        cascadia-mono-nf-fonts\
        gcc\
        git\
        golang\
        make\ 
        npm\ 
        nvim\ 
        glfw\
        ripgrep\ 
        shadow-utils\ 
        unzip\ 
        jetbrains-mono-fonts && \
dnf copr enable -y che/nerd-fonts && dnf install -y nerd-fonts && \
dnf clean all

ARG USERNAME=dev
ARG UID=1000
ARG GID=1000

RUN groupadd --gid ${GID} ${USERNAME} || true && \
    useradd --system --uid ${UID} --gid ${GID} --create-home --shell /bin/bash ${USERNAME} || true

USER ${USERNAME}
WORKDIR /home/${USERNAME}
COPY --from=downloader /go/bin/air $HOME/go/bin
COPY --from=downloader /go/bin/dlv $HOME/go/bin
ENV PATH="$PATH:$HOME/go/bin"
RUN git clone --branch mine https://github.com/douglascdev/kickstart.nvim .config/nvim

WORKDIR /usr/src

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

