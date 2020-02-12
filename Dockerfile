FROM golang:1.13-alpine AS build
WORKDIR /src/
COPY hostapp.go go.* /src/

RUN CGO_ENABLED=0 go build -o /bin/hostapp
FROM scratch
COPY --from=build /bin/hostapp /bin/hostapp
CMD ["/bin/hostapp"]