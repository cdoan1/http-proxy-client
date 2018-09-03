FROM quay.io/kubernetes-service-catalog/service-catalog-amd64:v0.1.30 as community

FROM scratch

ARG DATE
ARG IMAGE_DESCRIPTION
ARG IMAGE_NAME
ARG VCS_REF
ARG VCS_URL
ARG ICP_VERSION

LABEL org.label-schema.vendor="OSSC" \
    org.label-schema.name="http-proxy-client" \
    org.label-schema.description="description" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url=$VCS_URL \
    org.label-schema.schema-version="1.0" \
    org.label-schema.build-date="$DATE" \
    org.label-schema.version="0.1.0"

COPY --from=community /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY http-proxy-client http-proxy-client

# ADD --chown=65534:65534 tmp /tmp
# ADD --chown=65534:65534 var/run /var/run
# USER 65534

ENTRYPOINT ["./http-proxy-client"]
