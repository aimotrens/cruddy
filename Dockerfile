FROM node:22.11.0-alpine3.19@sha256:2ef8a5348c8936329810271059b2c1ba88f293f4d6849e5d2ade22652c3d396c AS node-builder

WORKDIR /src

ADD frontend/package.json frontend/package-lock.json ./
RUN npm i

ADD frontend/ ./
RUN npm run build
RUN mv dist/index.html dist/index.htm

# --------------------------------------------

FROM golang:1.23.2-alpine3.19@sha256:f6392ffebb028fed5ffe743ddb9716e38402c978779edd66474bb5d05f5e65e4 AS go-builder

ARG CRUDDY_VERSION

WORKDIR /src
ADD backend/go.mod backend/go.sum ./
RUN go mod download

ADD backend/ ./

COPY --from=node-builder /src/dist ./static

RUN go build -ldflags "-X \"main.cruddyVersion=${CRUDDY_VERSION}\" -X \"main.compileDate=$(date)\"" -o cruddy

# --------------------------------------------

FROM alpine:3.20@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

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