FROM fedora:42

RUN sudo dnf install -y cascadia-mono-nf-fonts gcc git golang make npm nvim ripgrep shadow-utils unzip

WORKDIR /root
RUN git clone https://github.com/douglascdev/dotfiles
RUN mv dotfiles/* .
RUN mv dotfiles/.* .
RUN git submodule init
RUN git submodule update
WORKDIR /root/.config/nvim
RUN git checkout mine

RUN go install github.com/air-verse/air@v1.62.0
ENV PATH="$PATH:/root/go/bin"

WORKDIR /usr/src

EXPOSE 8080

#CMD ["sleep", "infinity"]
CMD ["air", "-c", ".air.toml"]

