#!/bin/bash
ip=$1
m=$2
i=0
n=0
while [ $n -eq 0 ]; do
	echo "Esperando HTTP"
	sleep 5
	n=$(curl -s -i  --connect-timeout 4 http://$ip | grep 'HTTP/1.1 200 OK' | wc -l)
        i=$(expr $i + 1)
        if [ $i -gt $m ]; then
           n=1
        fi
done
echo OK
