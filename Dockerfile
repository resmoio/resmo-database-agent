FROM alpine:latest
COPY resmo-database-agent /resmo-database-agent
ENTRYPOINT ["/resmo-database-agent"]