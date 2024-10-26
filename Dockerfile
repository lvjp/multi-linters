# BUILD:BEGIN
# BUILD(actionlint):BEGIN
FROM rhysd/actionlint:1.7.3@sha256:7617f05bd698cd2f1c3aedc05bc733ccec92cca0738f3e8722c32c5b42c70ae6 AS actionlint
# BUILD(actionlint):END
# BUILD(gitleaks):BEGIN
FROM zricethezav/gitleaks:v8.19.3@sha256:b1081012aeb9026447deb2ecf4671f7a71cc035b9a1ce23a36c0a853c5dfde95 AS gitleaks
# BUILD(gitleaks):END
# BUILD(golangci-lint):BEGIN
FROM golang:1.23.2-alpine@sha256:9dd2625a1ff2859b8d8b01d8f7822c0f528942fe56cfe7a1e7c38d3b8d72d679 AS golang
FROM golangci/golangci-lint:v1.61.0-alpine@sha256:61e2d68adc792393fcb600340fe5c28059638d813869d5b4c9502392a2fb4c96 AS golangci-lint
# BUILD(golangci-lint):END
# BUILD(hadolint):BEGIN
FROM hadolint/hadolint:v2.12.0-alpine@sha256:3c206a451cec6d486367e758645269fd7d696c5ccb6ff59d8b03b0e45268a199 AS hadolint
# BUILD(hadolint):END
# BUILD(shellcheck):BEGIN
FROM koalaman/shellcheck:v0.10.0@sha256:2097951f02e735b613f4a34de20c40f937a6c8f18ecb170612c88c34517221fb AS shellcheck
# BUILD(shellcheck):END
# BUILD(shfmt):BEGIN
FROM mvdan/shfmt:v3.9.0@sha256:cb4ca87cc18e52f184a7ba1ae1ef7350b79a2c216ace78a0d24b473e87f0b8f5 AS shfmt
# BUILD(shfmt):END
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
# INSTALL(actionlint):BEGIN
COPY --link --from=actionlint /usr/local/bin/actionlint /usr/bin/actionlint
# INSTALL(actionlint):END
# INSTALL(gitleaks):BEGIN
COPY --link --from=gitleaks /usr/bin/gitleaks /usr/bin/
# INSTALL(gitleaks):END
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
# INSTALL(shellcheck):BEGIN
COPY --link --from=shellcheck /bin/shellcheck /usr/bin/shellcheck
# INSTALL(shellcheck):END
# INSTALL(shfmt):BEGIN
COPY --link --from=shfmt /bin/shfmt /usr/bin/
# INSTALL(shfmt):END
# INSTALL:END

ARG TARGETPLATFORM

# from context is defined with command line
# hadolint ignore=DL3022
COPY --from=multi-linters-binaries "${TARGETPLATFORM}/multi-linters" /usr/bin/

ENTRYPOINT ["/usr/bin/multi-linters"]
