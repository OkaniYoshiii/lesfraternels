FROM golang:alpine3.22 AS golang

ARG UID
ARG GID

RUN addgroup -S -g ${GID} lesfraternels_app_group
RUN adduser -S -u ${UID} -G lesfraternels_app_group lesfraternels_app_user

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chown -R lesfraternels_app_user:lesfraternels_app_group ./

USER lesfraternels_app_user

CMD ["air", "-c", ".air.toml"]
