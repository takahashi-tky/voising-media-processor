#!/bin/bash

GRPC_OUT_BASE_PATH="api/proto/gen"
SERVER_PROJECT_PATH="../voising-fc-api"

GEN_CMD="protoc -I ${SERVER_PROJECT_PATH}/api/proto \
--include_imports \
--include_source_info \
--proto_path=${SERVER_PROJECT_PATH}/github.com/googleapis \
--proto_path=${SERVER_PROJECT_PATH}/github.com/grpc-gateway \
--go_out=${GRPC_OUT_BASE_PATH}/%s \
--go_opt=paths=source_relative \
--descriptor_set_out=api/proto/gen/api_descriptor.pb \
--go-grpc_out=${GRPC_OUT_BASE_PATH}/%s \
--go-grpc_opt=paths=source_relative \
${SERVER_PROJECT_PATH}/api/proto/"

PROTO_FILES=($(ls -F ${SERVER_PROJECT_PATH}/api/proto/*.proto | xargs -n1 basename))

for PROTO in ${PROTO_FILES[@]}; do
  if [ ${PROTO} = "swagger_root_config.proto" ]; then
    continue
  fi
  echo -n "Generate code from ${PROTO} [Y/n]: "
  # shellcheck disable=SC2162
  read ANS
  # shellcheck disable=SC1009
  case $ANS in Yes | yes | [yY] )
    # shellcheck disable=SC2209
    # shellcheck disable=SC2001
    NAME=$(echo "$PROTO" | sed 's/\.[^\.]*$//')
    mkdir -p "${GRPC_OUT_BASE_PATH}/${NAME}"
    # shellcheck disable=SC2046
    # shellcheck disable=SC2059
    eval $(printf "${GEN_CMD}${PROTO}" "${NAME}" "${NAME}")
    ;;
  * )
    ;;
  esac
done