FROM golang:1.17

COPY go.mod go.sum /
WORKDIR /
RUN [ "go", "mod", "download" ]

COPY . /src
COPY ./entrypoint.sh /entrypoint.sh
WORKDIR /src/miniproject2
RUN [ "go", "build", "-o", "/build/taskapi", "github.com/GregersSR/taskinator"]
EXPOSE 8080
ENTRYPOINT [ "bash", "/entrypoint.sh" ]
