ARG DOCKERFILE_HADOLINT_VERSION=v2.12.0-alpine
FROM hadolint/hadolint:${DOCKERFILE_HADOLINT_VERSION} AS dockerfile_hadolint

FROM alpine:3.20.2

COPY --link --from=dockerfile_hadolint /bin/hadolint /usr/bin/hadolint
