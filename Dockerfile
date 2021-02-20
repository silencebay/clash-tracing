FROM --platform=$TARGETPLATFORM golang:alpine AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM

RUN apk add --no-cache make git
WORKDIR /scraper-src
COPY . /scraper-src
RUN go mod download && \
    go build -o scraper . && \
    mv ./scraper /scraper

FROM --platform=$TARGETPLATFORM alpine AS runtime
ARG TARGETPLATFORM
ARG BUILDPLATFORM

RUN apk add --no-cache ca-certificates
COPY --from=builder /scraper /
ENTRYPOINT ["/scraper"]
