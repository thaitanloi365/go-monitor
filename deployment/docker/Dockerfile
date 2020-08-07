FROM golang:alpine as builder

LABEL maintainer="Thai Tan Loi <thaitanloi365@gmail.com>"

RUN apk update && apk add --no-cache git

ARG BUILD_ENV=

ENV BUILD_ENV=${BUILD_ENV}

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download

RUN if [ "${BUILD_ENV}" = "local" ] ; then \
  echo "Install CompileDaemon";\
  go get github.com/githubnemo/CompileDaemon ;\
  else \
  cp deployment/config/${BUILD_ENV}.env .env;\
  govvv build -o main  \
  -ldflags "-X 'main.ServiceName=${SERVICE_NAME}' \
  -X 'main.BuildEnv=${BUILD_ENV}'" \
  ./cmd/${SERVICE_NAME} ;\
  fi


# Production stage
FROM alpine
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata

WORKDIR /root

COPY --from=builder /app .

ENTRYPOINT /root/main --config=.env serve
