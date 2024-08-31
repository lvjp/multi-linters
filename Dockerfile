ARG DOCKERFILE_HADOLINT_VERSION=v2.12.0-alpine@sha256:3c206a451cec6d486367e758645269fd7d696c5ccb6ff59d8b03b0e45268a199
FROM hadolint/hadolint:${DOCKERFILE_HADOLINT_VERSION} AS dockerfile_hadolint

FROM alpine:3.20.2@sha256:0a4eaa0eecf5f8c050e5bba433f58c052be7587ee8af3e8b3910ef9ab5fbe9f5

COPY --link --from=dockerfile_hadolint /bin/hadolint /usr/bin/hadolint
