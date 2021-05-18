FROM debian:buster-slim

ENV LTC_VERSION=0.18.1
ENV DEBIAN_FRONTEND noninteractive

RUN set -eux && apt-get update && apt-get install -y gnupg curl \
    && LTC_VERSION=0.18.1 \
    && LTC_GPGKEY=FE3348877809386C; \
    server="hkp://p80.pool.sks-keyservers.net:80" \
    && echo "fetching GPG key ${LTC_GPGKEY} from ${server}"; \
        gpg --keyserver ${server} --recv-keys ${LTC_GPGKEY}; 

RUN curl -SLO https://download.litecoin.org/litecoin-${LTC_VERSION}/linux/litecoin-${LTC_VERSION}-x86_64-linux-gnu.tar.gz \
&& curl -SLO https://download.litecoin.org/litecoin-${LTC_VERSION}/linux/litecoin-${LTC_VERSION}-x86_64-linux-gnu.tar.gz.asc \
&& gpg --verify litecoin-${LTC_VERSION}-x86_64-linux-gnu.tar.gz.asc \
&& tar xfz /litecoin-${LTC_VERSION}-x86_64-linux-gnu.tar.gz \
&& mv litecoin-${LTC_VERSION}/bin/* /usr/local/bin/ \
&& rm -rf /var/lib/apt/lists/*
RUN echo "here"

EXPOSE 9332 9333 19332 19333 19444
RUN echo $PATH

ENTRYPOINT ["litecoind"]

