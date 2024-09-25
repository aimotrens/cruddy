FROM node:21.7.3-alpine3.19@sha256:1e13649e44d505d5410164f5b7325e4ff1ae551e87e6e4f17d74f6b9b0affbff AS node-builder

WORKDIR /src

ADD frontend/package.json frontend/package-lock.json ./
RUN npm i

ADD frontend/ ./
RUN npm run build
RUN mv dist/index.html dist/index.htm

# --------------------------------------------

FROM golang:1.23.1-alpine3.19@sha256:e0ea2a119ae0939a6d449ea18b2b1ba30b44986ec48dbb88f3a93371b4bf8750 AS go-builder

ARG CRUDDY_VERSION

WORKDIR /src
ADD backend/go.mod backend/go.sum ./
RUN go mod download

ADD backend/ ./

COPY --from=node-builder /src/dist ./static

RUN go build -ldflags "-X \"main.cruddyVersion=${CRUDDY_VERSION}\" -X \"main.compileDate=$(date)\"" -o cruddy

# --------------------------------------------

FROM alpine:3.20@sha256:e72ad0747b9dc266fca31fb004580d316b6ae5b0fdbbb65f17bbe371a5b24cff

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