export env=dev

dirPath=$PWD

cd "${dirPath}"/OrderSubscriber && go build main.go
cd "${dirPath}"/OrderPublisher && go build main.go
