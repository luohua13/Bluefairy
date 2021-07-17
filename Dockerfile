ARG image_registry=build-harbor.alauda.cn/asm

#FROM golang:1.15  AS builder
#COPY . $GOPATH/src/Bluefairy
#WORKDIR $GOPATH/src/Bluefairy
#
#RUN  make compile-local
#
#RUN mkdir /Bluefairy
#COPY ./Bluefairy /Bluefairy/Bluefairy

#RUN case "${TARGETARCH}" in \
#  amd64) upx /opt/${TARGETARCH}/manager ;; \
#  arm64) ;; \
#  *) echo "unsupported architecture"; exit 1 ;; \
#  esac \
#  && cp /opt/${TARGETARCH}/manager /manager

# step 2: build
FROM ${image_registry}/runner:0.1-alpine3.12.1

RUN apk --no-cache --update add ca-certificates

WORKDIR /
COPY ./Bluefairy ./Bluefairy
RUN chmod +x ./Bluefairy
CMD ["./Bluefairy"]
