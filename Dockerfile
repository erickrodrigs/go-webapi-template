FROM golang:1.15.7-alpine

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}" \
  GO111MODULE=on \
  HOST=0.0.0.0 \
  PROXY_PORT=3000 \
  APP_PORT=3001

RUN go get github.com/codegangsta/gin && \
  go get -u github.com/gin-gonic/gin

CMD /go/bin/gin \
  --port ${PROXY_PORT} \
  --appPort ${APP_PORT} \
  --path ./src \
  --excludeDir .pgdata \
  run
