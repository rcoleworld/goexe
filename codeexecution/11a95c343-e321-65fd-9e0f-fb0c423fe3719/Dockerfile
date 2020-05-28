FROM golang:latest
RUN mkdir /app
COPY ./11a95c343-e321-65fd-9e0f-fb0c423fe3719.go /app/11a95c343-e321-65fd-9e0f-fb0c423fe3719.go 
WORKDIR /app
RUN go build -o 11a95c343-e321-65fd-9e0f-fb0c423fe3719 .
CMD ["/app/11a95c343-e321-65fd-9e0f-fb0c423fe3719"]%!(EXTRA string=11a95c343-e321-65fd-9e0f-fb0c423fe3719)