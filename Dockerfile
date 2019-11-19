FROM golang:1.13.4-alpine

MAINTAINER Cache Lab <hello@cachelab.co>

COPY r53_domain_manager /bin/r53dm

USER nobody

ENTRYPOINT ["/bin/r53dm"]
