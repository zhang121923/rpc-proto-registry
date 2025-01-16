### 编译proto文件

需要使用 protoc 工具和 protoc-gen-go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

命令：
protoc --go_out=. --go-grpc_out=. ./protobuf/*.proto 
* --go_out=. 指定输出目录为当前目录
* --go_opt=paths=source_relative 确保生成的文件路径与源文件路径相对一致
* ./protobuf/*.proto 为想要编译的proto文件， 支持指定名称文件，或者全部\*