FROM node:21.5-alpine3.19@sha256:82c93cef3d2acbb2557c5fda48214fbc2bf5385edfb4d96d990690d75ddabf7b AS node-builder

WORKDIR /src

ADD frontend/package.json frontend/package-lock.json ./
RUN npm i

ADD frontend/ ./
RUN npm run build
RUN mv dist/index.html dist/index.htm

# --------------------------------------------

FROM golang:1.21.6-alpine3.19@sha256:cf033f46ceb1f5cc29a6ebd913d8b2b007c198928018afeee346f0f41fe3d2d8 AS go-builder

ARG CRUDDY_VERSION

WORKDIR /src
ADD backend/go.mod backend/go.sum ./
RUN go mod download

ADD backend/ ./

COPY --from=node-builder /src/dist ./static

RUN go build -ldflags "-X \"main.cruddyVersion=${CRUDDY_VERSION}\" -X \"main.compileDate=$(date)\"" -o cruddy

# --------------------------------------------

FROM alpine:3.19@sha256:51b67269f354137895d43f3b3d810bfacd3945438e94dc5ac55fdac340352f48

WORKDIR /app

COPY --from=go-builder /src/cruddy ./cruddy

EXPOSE 4231
ENV GIN_MODE=release
ENV CRUDDY_ROOT_DIR=/srv

CMD ["./cruddy"]