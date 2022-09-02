#!/bin/zsh
echo "1======================="
go run demo1.go echo "haha1" "haha2" "haha3"

echo "2======================="
go run demo1.go echo times -t 3 "haha1" "haha2"

echo "3======================="
go run demo1.go print "haha1" "haha2"

echo "4======================="
echo "5======================="


