FROM golang:1.20.14-bullseye AS builder
RUN apt-get update && apt-get install -y wget
ENV CHAOSD_VERSION=v1.4.0
WORKDIR /
RUN wget https://github.com/chaos-mesh/chaosd/archive/refs/tags/$CHAOSD_VERSION.tar.gz && mkdir /chaosd && tar xfz $CHAOSD_VERSION.tar.gz -C /chaosd --strip-components=1
WORKDIR /chaosd
RUN make chaosd 
RUN make chaos-tools
FROM scratch
COPY --from=builder /chaosd/bin/chaosd /rootfs/usr/local/lib/containers/chaosd/chaosd
COPY --from=builder /chaosd/bin/tools /rootfs/usr/local/lib/containers/chaosd/usr/bin
COPY --from=builder /bin/dd /rootfs/usr/local/lib/containers/chaosd/bin/dd
COPY --from=builder /lib /rootfs/usr/local/lib/containers/chaosd/lib
COPY --from=builder /lib64 /rootfs/usr/local/lib/containers/chaosd/lib64
COPY chaosd.yaml /rootfs/usr/local/etc/containers/chaosd.yaml
COPY manifest.yaml /
WORKDIR /rootfs/usr/local/lib/containers/chaosd/



