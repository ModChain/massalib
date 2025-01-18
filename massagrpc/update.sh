#!/bin/sh

set -e

if [ ! -d massa-proto ]; then
	git clone https://github.com/massalabs/massa-proto.git
else
	( cd massa-proto && git reset --hard && git pull )
fi

# prepare protoc mapping
PROTOC_ARGS=(
	--proto_path=./massa-proto/proto/apis
	--proto_path=./massa-proto/proto/commons
	--proto_path=./massa-proto/proto/third_party
	--go_out=.
	--go-grpc_out=.
	--go_opt=module=github.com/ModChain/massagrpc
	--go-grpc_opt=module=github.com/ModChain/massagrpc
)

PROTOC_PATHS=(
	massa-proto/proto/apis/massa/api/v1
	massa-proto/proto/commons
)

for p in ${PROTOC_PATHS[@]}; do
	for f in $(find "$p" -name "*.proto"); do
		#option go_package = "github.com/massalabs/massa/model/v1;v1";
		sed -i -e 's#^option go_package.*#option go_package = "github.com/ModChain/massagrpc";#' "$f"
		PROTOC_ARGS+=(${f})
	done
done

echo "running: protoc ${PROTOC_ARGS[@]}"
protoc "${PROTOC_ARGS[@]}"
