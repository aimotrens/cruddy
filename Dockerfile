FROM node:21.7.3-alpine3.19@sha256:1e13649e44d505d5410164f5b7325e4ff1ae551e87e6e4f17d74f6b9b0affbff AS node-builder

WORKDIR /src

ADD frontend/package.json frontend/package-lock.json ./
RUN npm i

ADD frontend/ ./
RUN npm run build
RUN mv dist/index.html dist/index.htm

# --------------------------------------------

FROM golang:1.22.4-alpine3.19@sha256:d9b1f00c269b4f82b8a28c87b4feeb34164305b29f6ce37f5c8cb076aac691c3 AS go-builder

ARG CRUDDY_VERSION

WORKDIR /src
ADD backend/go.mod backend/go.sum ./
RUN go mod download

ADD backend/ ./

COPY --from=node-builder /src/dist ./static

RUN go build -ldflags "-X \"main.cruddyVersion=${CRUDDY_VERSION}\" -X \"main.compileDate=$(date)\"" -o cruddy

# --------------------------------------------

FROM alpine:3.20@sha256:b2bf705cebbc5a03a0b58e49ec152261ea7da707f40f7d5db1c4f75ed72d43c1

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