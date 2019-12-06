rm sarama
GOARCH=amd64 GOOS=linux go build -o sarama .
docker build -t z416177937/z416177937:sarama2 .
docker login --username=z416177937
docker push z416177937/z416177937:sarama2