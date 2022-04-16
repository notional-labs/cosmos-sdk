## To test locally:
# docker build --pull --rm -f "contrib/devtools/Dockerfile" -t cosmossdk-proto:latest "contrib/devtools"
# docker run -v $(pwd):/workspace --workdir /workspace cosmossdk-proto sh ./scripts/protocgen.sh

FROM bufbuild/buf:1.1.0 as BUILDER

FROM faddat/archlinux

RUN pacman -Sy
RUN pacman -S --noconfirm archlinux-keyring
RUN pacman -Syyu --noconfirm nodejs npm git make go

ENV GOLANG_PROTOBUF_VERSION=1.28.0 
ENV GOGO_PROTOBUF_VERSION=1.4.0 
ENV GRPC_GATEWAY_VERSION=1.16.0 
ENV GOPATH=/go 
ENV PATH=$PATH:/go/bin


RUN go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest

RUN git clone https://github.com/regen-network/cosmos-proto && \
	cd cosmos-proto && \
	go install ./...


RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GOLANG_PROTOBUF_VERSION} && \
go install github.com/cosmos/gogoproto/protoc-gen-gogo@v${GOGO_PROTOBUF_VERSION} && \
go install github.com/cosmos/gogoproto/protoc-gen-gogofast@v${GOGO_PROTOBUF_VERSION} && \
go install github.com/cosmos/gogoproto/protoc-gen-gogofaster@v${GOGO_PROTOBUF_VERSION} 


RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v${GRPC_GATEWAY_VERSION} && \
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v${GRPC_GATEWAY_VERSION} 


RUN npm install -g swagger-combine

COPY --from=BUILDER /usr/local/bin /usr/local/bin
