FROM registry.access.redhat.com/ubi9/go-toolset:latest

WORKDIR /app

# deps
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . .

ARG TOKEN
ARG ENDPOINT_URL

ENV TOKEN=$TOKEN
ENV ENDPOINT_URL=$ENDPOINT_URL

WORKDIR /app

RUN go build -v -o server

EXPOSE 8980

CMD ["/app/server"]