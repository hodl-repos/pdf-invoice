FROM golang:1.20-alpine as builder
RUN apk update && apk upgrade && apk add --no-cache ca-certificates && apk add --update make
# get latest ca-certificates
RUN update-ca-certificates
WORKDIR /build
COPY . .

RUN make build

FROM scratch
ARG SERVICE

COPY ./build/passwd /etc/passwd

USER nobody
COPY --from=builder ./build/bin/service /server
COPY ./active.en.toml ./active.en.toml
COPY ./active.de.toml ./active.de.toml

# copy latest ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "/server" ]
