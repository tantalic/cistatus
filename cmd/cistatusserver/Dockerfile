FROM scratch
MAINTAINER Kevin Stock <kevinstock@tantalic.com>

# SSL CA Root Certs
ADD ca-certificates.crt /etc/ssl/certs/

# Labels: http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="cistatusserver" \
      org.label-schema.description="A server to poll a Continuous Integration server (currently GitLab CI) and provide the status via a JSON and Websocket API." \
      org.label-schema.url="https://tantalic.com/cistatus" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/tantalic/cistatus" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

WORKDIR /
ADD ./cistatusserver-linux_amd64 /cistatusserver
EXPOSE 80

# The following options can be configured via environment variables:
# ENV GITLAB_API_BASE_URL=http://example.githost.io
# ENV GITLAB_API_TOKEN=xxxxxxxxxx
# ENV GITLAB_REFRESH_PERIOD=10s
# ENV CI_STATUS_HTTP_SERVER_JWT_ALGORITHM=HS512
# ENV CI_STATUS_HTTP_SERVER_JWT_SECRET=xxxxxxxxxx

CMD ["/cistatusserver"]
