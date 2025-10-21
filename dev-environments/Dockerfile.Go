FROM fedora:42

RUN sudo dnf update
RUN sudo dnf install cascadia-mono-nf-fonts gcc git golang make npm nvim ripgrep shadow-utils unzip
RUN go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 8080
CMD ["dlv", "debug", "--headless=true", "--listen=0.0.0.0:2345", "--api-version=2", "--accept-multiclient", "--", "go", "run", "main.go", "-cfg", "config.db", "-debug"]
