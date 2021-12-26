FROM golang:1.17-bullseye as build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o http-ping 

FROM scratch
COPY --from=build /app/http-ping /
ENV GIN_MODE=release
CMD ["/http-ping"]