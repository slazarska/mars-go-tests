# Stage 1: Builder — сборка и тесты
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p ./test/allure-results

RUN go test ./...

# Stage 2: Allure
FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y curl unzip openjdk-11-jre-headless && \
    curl -L https://github.com/allure-framework/allure2/releases/download/2.22.1/allure-2.22.1.zip -o /tmp/allure.zip && \
    unzip /tmp/allure.zip -d /opt/ && \
    ln -s /opt/allure-2.22.1/bin/allure /usr/bin/allure && \
    rm /tmp/allure.zip && \
    apt-get clean

WORKDIR /app

COPY --from=builder /app/test/allure-results ./test/allure-results

CMD ["allure", "generate", "./test/allure-results", "-o", "./test/allure-report", "--clean"]