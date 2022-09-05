export env=dev

dirPath=$PWD
dockerComposeLogFile=${dirPath}/docker-compose.log
rm "${dockerComposeLogFile}"
docker-compose -f ./docker-compose.yml down
docker-compose -f ./docker-compose.yml up >>./docker-compose.log &

(tail -f -n0 "${dockerComposeLogFile}" &) | grep -q "Ready to accept connections"
(tail -f -n0 "${dockerComposeLogFile}" &) | grep -q "mysqld: ready for connections"

cd "${dirPath}"/OrderSubscriber && ./main >>./app.log &
cd "${dirPath}"/OrderPublisher && ./main >>./app.log &
