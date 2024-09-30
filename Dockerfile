# BUILD:BEGIN
# BUILD(golangci-lint):BEGIN
FROM golang:1.23.0-alpine AS golang
FROM golangci/golangci-lint:v1.61.0-alpine@sha256:61e2d68adc792393fcb600340fe5c28059638d813869d5b4c9502392a2fb4c96 AS golangci-lint
# BUILD(golangci-lint):END
# BUILD(hadolint):BEGIN
FROM hadolint/hadolint:v2.12.0-alpine@sha256:3c206a451cec6d486367e758645269fd7d696c5ccb6ff59d8b03b0e45268a199 AS hadolint
# BUILD(hadolint):END
# BUILD:END

FROM alpine:3.20.3@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

# APK_ADD:BEGIN
RUN apk add --no-cache \
    fossil==2.24-r1 \
    git==2.45.2-r0 \
    mercurial==6.7.4-r1 \
    subversion==1.14.3-r2
# APK_ADD:END

# INSTALL:BEGIN
# INSTALL(golangci-lint):BEGIN
COPY --from=golang /usr/local/go/go.env /usr/lib/go/
COPY --from=golang /usr/local/go/bin/ /usr/lib/go/bin/
COPY --from=golang /usr/local/go/lib/ /usr/lib/go/lib/
COPY --from=golang /usr/local/go/pkg/ /usr/lib/go/pkg/
COPY --from=golang /usr/local/go/src/ /usr/lib/go/src/
COPY --link --from=golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint
ENV PATH="${PATH}:/usr/lib/go/bin"
# INSTALL(golangci-lint):END
# INSTALL(hadolint):BEGIN
COPY --link --from=hadolint /bin/hadolint /usr/bin/hadolint
# INSTALL(hadolint):END
# INSTALL:END

ARG TARGETPLATFORM

# from context is defined with command line
# hadolint ignore=DL3022
COPY --from=multi-linters-binaries "${TARGETPLATFORM}/multi-linters" /usr/bin/

ENTRYPOINT ["/usr/bin/multi-linters"]
