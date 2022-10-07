#!/bin/bash

set -e

protoc_version="3.17.3"
protoc_gen_go_version="v1.26.0"
protoc_gen_go_grpc_version="1.1.0"
protoc_mockgen_version="v1.4.4"

check_protoc_version() {
  if ! command -v protoc >/dev/null; then
    echo "command protoc could not be found"
    echo "please install protoc first. Ref: https://github.com/protocolbuffers/protobuf/releases"
    exit
  fi
  if [ ! -d "/usr/local/include/google/" ]; then
    echo "missing google protobuf"
    echo "please install it first. Ref: https://github.com/grpc-ecosystem/grpc-gateway/issues/422#issuecomment-409809309"
    exit
  fi
  version=$(protoc --version)
  if ! [[ ${version} == *"${protoc_version}"* ]]; then
    echo "invalid proto version ${version}"
    exit 1
  fi
}

check_protoc_gen_go_version() {
  version=$(protoc-gen-go --version)
  if ! [[ ${version} == *"${protoc_gen_go_version}"* ]]; then
    echo "invalid proto-gen-go version ${version}"
    exit 1
  fi
}

check_protoc_gen_go_grpc_version() {
  version=$(protoc-gen-go-grpc --version)
  if ! [[ ${version} == *"${protoc_gen_go_grpc_version}"* ]]; then
    echo "invalid proto-gen-go-grpc version ${version}"
    exit 1
  fi
}

check_protoc_gen_grpc_gateway() {
  if ! command -v protoc-gen-grpc-gateway &> /dev/null; then
    echo "command protoc-gen-grpc-gateway could not be found"
    echo "run 'go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.4.0' to install"
    exit 1
  fi
}

check_protoc_gen_openapiv2() {
  if ! command -v protoc-gen-openapiv2 &> /dev/null; then
    echo "command protoc-gen-openapiv2 could not be found"
    echo "run 'go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.4.0' to install"
    exit 1
  fi
}

check_mockgen() {
  if ! command -v mockgen &> /dev/null; then
    echo "command mockgen could not be found"
    echo "run 'go get github.com/golang/mock/mockgen@v1.4.4' to install"
    exit 1
  fi
  version=$(mockgen --version)
  if ! [[ ${version} == *"${protoc_mockgen_version}"* ]]; then
    echo "invalid mockgen version ${version}"
    exit 1
  fi
}

check_statik() {
  if ! command -v statik &> /dev/null; then
    echo "command statik could not be found"
    echo "run 'go get github.com/rakyll/statik' to install"
    exit 1
  fi
}

generate_code() {
  protoc -I=./proto \
      --go_opt=paths=source_relative \
      --go_out=./pb \
      --go-grpc_opt=paths=source_relative \
      --go-grpc_out=./pb \
      --grpc-gateway_opt=paths=source_relative \
      --grpc-gateway_out=./pb \
      --openapiv2_out ./swagger \
      proto/*.proto
}

generate_mock() {
  if compgen -G pb/*_grpc.pb.go > /dev/null; then
    cd pb
    package=pb
    file=$(ls *_grpc.pb.go)
    file_name=$(echo "${file%.pb.go}")
    mockgen -package=pb -source=${file_name}.pb.go -destination=${file_name}_mock.pb.go
    cd - &> /dev/null
  fi
}

generate_swagger_module() {
    statik -src=./swagger/ -dest=. -p=pb -f -ns=pb
}

check_protoc_version
check_protoc_gen_go_version
check_protoc_gen_go_grpc_version
check_protoc_gen_grpc_gateway
check_protoc_gen_openapiv2
check_mockgen
check_statik

mkdir -p swagger pb
generate_code
# generate_mock and generate_swagger_module must be run after generate_code
generate_mock
generate_swagger_module
