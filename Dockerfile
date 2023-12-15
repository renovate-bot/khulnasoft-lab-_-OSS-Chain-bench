FROM alpine:3.19 as build

WORKDIR /home/oss-chain-bench

RUN apk update && \
    apk --no-cache add make go

COPY . .

RUN make build

FROM alpine:3.19 as product

WORKDIR /home/oss-chain-bench

COPY --from=build /home/oss-chain-bench/oss-chain-bench /usr/local/bin/oss-chain-bench
COPY --from=build /home/oss-chain-bench/templates/*.tpl templates/


ENTRYPOINT [ "oss-chain-bench" ]