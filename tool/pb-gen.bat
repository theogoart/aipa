rem generate common/types
protoc\bin\protoc.exe -I ..\..\..\..\ --go_out=:..\..\..\..\  ..\..\..\..\github.com\aipadad\aipa\common\types\transaction.proto ..\..\..\..\github.com\aipadad\aipa\common\types\basic-transaction.proto ..\..\..\..\github.com\aipadad\aipa\common\types\block.proto

rem generate api
protoc\bin\protoc.exe -I ..\..\..\..\ --go_out=:..\..\..\..\ --micro_out=:..\..\..\..\  ..\..\..\..\github.com\aipadad\aipa\api\transaction.proto ..\..\..\..\github.com\aipadad\aipa\api\basic-transaction.proto  ..\..\..\..\github.com\aipadad\aipa\api\chain.proto

rem generate tool/example
protoc\bin\protoc.exe -I ..\..\..\..\ --go_out=:..\..\..\..\ --micro_out=:..\..\..\..\   ..\..\..\..\github.com\aipadad\aipa\tool\example\example.proto
