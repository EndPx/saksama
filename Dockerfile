# Multi-stage build producing static binaries. The image carries no API key;
# supply SAKSAMA_* at run time. The default entrypoint is the keyless eval,
# so judges can verify the committed numbers with no key and no network.
FROM golang:1.22 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/eval        ./cmd/eval && \
    CGO_ENABLED=0 go build -o /out/baseline    ./cmd/baseline && \
    CGO_ENABLED=0 go build -o /out/solution    ./cmd/solution && \
    CGO_ENABLED=0 go build -o /out/trajectory  ./cmd/trajectory

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=build /out/ /app/bin/
COPY data/    /app/data/
COPY results/ /app/results/
COPY memos/   /app/memos/
# Verify the committed scores with no API key:
#   docker run --rm saksama
ENTRYPOINT ["/app/bin/eval"]
