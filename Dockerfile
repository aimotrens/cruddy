FROM node:22.12.0-alpine3.19@sha256:40dc4b415c17b85bea9be05314b4a753f45a4e1716bb31c01182e6c53d51a654 AS node-builder

WORKDIR /src

ADD frontend/package.json frontend/package-lock.json ./
RUN npm i

ADD frontend/ ./
RUN npm run build
RUN mv dist/index.html dist/index.htm

# --------------------------------------------

FROM golang:1.23.4-alpine3.19@sha256:3612e783d3b5b7454866294988520f8ea77167aedbefe77f06aa1ad4ef0305e2 AS go-builder

ARG CRUDDY_VERSION

WORKDIR /src
ADD backend/go.mod backend/go.sum ./
RUN go mod download

ADD backend/ ./

COPY --from=node-builder /src/dist ./static

RUN go build -ldflags "-X \"main.cruddyVersion=${CRUDDY_VERSION}\" -X \"main.compileDate=$(date)\"" -o cruddy

# --------------------------------------------

FROM alpine:3.20@sha256:1e42bbe2508154c9126d48c2b8a75420c3544343bf86fd041fb7527e017a4b4a

WORKDIR /app

COPY --from=go-builder /src/cruddy ./cruddy

EXPOSE 4231
ENV GIN_MODE=release
ENV CRUDDY_ROOT_DIR=/srv

RUN mkdir -p /srv && \
    addgroup -S -g 1000 cruddy && \
    adduser -S -D -H -u 1000 cruddy && \
    addgroup cruddy cruddy && \
    chown -R cruddy:cruddy /srv

USER 1000

CMD ["./cruddy"]