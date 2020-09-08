FROM golang:1.15-alpine3.12 as build

# Unset GOPATH for module support
ENV GOPATH=
WORKDIR /go/src/app
COPY . .
RUN GOOS=linux GOARCH=amd64 \
    && go build -o header-dumper -v


FROM alpine:3.12

ENV HTTP_PORT=8000
RUN addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app

WORKDIR /home/app
COPY --from=build /go/src/app/header-dumper .
RUN chown -R app:app /home/app \
    && chmod +x header-dumper

USER app
CMD ["./header-dumper"]
