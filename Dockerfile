# BUILD:BEGIN
# BUILD(actionlint):BEGIN
# renovate: datasource=docker
FROM rhysd/actionlint:1.7.4@sha256:82244e1db1c60d82c7792180a48dd0bcb838370bb589d53ff132503fc9485868 AS actionlint
# BUILD(actionlint):END
# BUILD(gitleaks):BEGIN
# renovate: datasource=docker
FROM zricethezav/gitleaks:v8.21.2@sha256:0e99e8821643ea5b235718642b93bb32486af9c8162c8b8731f7cbdc951a7f46 AS gitleaks
# BUILD(gitleaks):END
# BUILD(golangci-lint):BEGIN
# renovate: datasource=docker
FROM golang:1.23.3-alpine@sha256:09742590377387b931261cbeb72ce56da1b0d750a27379f7385245b2b058b63a AS golang
# renovate: datasource=docker
FROM golangci/golangci-lint:v1.62.0-alpine@sha256:a94c2da655a1f3e66bb7d2c9232324def349f2ba56a3d715d89e83d98f5bd933 AS golangci-lint
# BUILD(golangci-lint):END
# BUILD(hadolint):BEGIN
# renovate: datasource=docker
FROM hadolint/hadolint:v2.12.0-alpine@sha256:3c206a451cec6d486367e758645269fd7d696c5ccb6ff59d8b03b0e45268a199 AS hadolint
# BUILD(hadolint):END
# BUILD(shellcheck):BEGIN
# renovate: datasource=docker
FROM koalaman/shellcheck:v0.10.0@sha256:2097951f02e735b613f4a34de20c40f937a6c8f18ecb170612c88c34517221fb AS shellcheck
# BUILD(shellcheck):END
# BUILD(shfmt):BEGIN
# renovate: datasource=docker
FROM mvdan/shfmt:v3.10.0@sha256:d19cc37644449fe9a488f234d2c0cf0b770eaf6a5a40e30103e8099013ef8f9e AS shfmt
# BUILD(shfmt):END
# BUILD:END

# renovate: datasource=docker
FROM alpine:3.20.3@sha256:1e42bbe2508154c9126d48c2b8a75420c3544343bf86fd041fb7527e017a4b4a

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
CMD ["run"]
