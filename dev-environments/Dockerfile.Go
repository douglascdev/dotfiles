FROM devbase

RUN sudo dnf install -y golang

RUN go install github.com/air-verse/air@v1.62.0
ENV PATH="$PATH:/root/go/bin"

WORKDIR /usr/src

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

