protoc --go_out=./  \
       --go-grpc_out=./ \
       school.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./school.proto