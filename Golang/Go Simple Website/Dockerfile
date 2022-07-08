FROM golang:latest 
WORKDIR /app
COPY . /app
RUN go mod init webapp && go mod tidy
RUN go build -o /webapp
EXPOSE 8080
CMD ["/webapp"]
