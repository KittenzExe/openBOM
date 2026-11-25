FROM golang:1.27.3 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /openbom .

FROM gcr.io/distroless/static-debian13
WORKDIR /app
COPY --from=build /openbom /app/openbom
COPY disclaimer.md /app/disclaimer.md
EXPOSE 3838
ENTRYPOINT ["/app/openbom"]
