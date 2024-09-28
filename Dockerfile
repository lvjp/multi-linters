FROM hadolint/hadolint:v2.12.0-alpine@sha256:3c206a451cec6d486367e758645269fd7d696c5ccb6ff59d8b03b0e45268a199 AS dockerfile_hadolint
FROM golangci/golangci-lint:v1.61.0-alpine@sha256:61e2d68adc792393fcb600340fe5c28059638d813869d5b4c9502392a2fb4c96 AS go_golangci-lint

FROM alpine:3.20.3@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

COPY --link --from=dockerfile_hadolint /bin/hadolint /usr/bin/hadolint
COPY --link --from=go_golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint
