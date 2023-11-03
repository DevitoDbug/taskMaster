FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
COPY /cmd/main/main.go . 

RUN go mod download 

COPY  cmd/ ./cmd/ 
COPY pkg/ ./pkg/
COPY static/ ./static/


RUN go build -o bin ./cmd/main/main.go

EXPOSE 8000

CMD [ "./bin" ]