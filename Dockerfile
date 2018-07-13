FROM go:1.10.0 AS build

ADD ./main.go /app/main.go
RUN go build -o /app/server /app/main.go

FROM alpine:latest
LABEL maintainer=nasa9084

COPY --from build /app/server .
RUN apk --no-cache add ca-certificates
CMD ["server"]
