FROM node:21.7.3-alpine3.19@sha256:db8772d9f5796ac4e8c47508038c413ea1478da010568a2e48672f19a8b80cd2 AS node-builder

WORKDIR /src

ADD frontend/package.json frontend/package-lock.json ./
RUN npm i

ADD frontend/ ./
RUN npm run build
RUN mv dist/index.html dist/index.htm

# --------------------------------------------

FROM golang:1.22.3-alpine3.19@sha256:df479aa7d1298917c7de7fd56c7231ec450149f1e63d960ea96aebf9ff9240c5 AS go-builder

ARG CRUDDY_VERSION

WORKDIR /src
ADD backend/go.mod backend/go.sum ./
RUN go mod download

ADD backend/ ./

COPY --from=node-builder /src/dist ./static

RUN go build -ldflags "-X \"main.cruddyVersion=${CRUDDY_VERSION}\" -X \"main.compileDate=$(date)\"" -o cruddy

# --------------------------------------------

FROM alpine:3.19@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

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