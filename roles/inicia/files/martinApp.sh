#!/bin/bash
for i in $(ps -ef | grep martinApp.go | grep -v grep | awk '{print $2}'); do
        echo Killing $i
        kill -9 $i
done
echo Iniciando App
sleep 1
export TNS_ADMIN=/home/opc/tns
nohup go run martinApp.go $1  $2 otro1 otro2 &
sleep 5
echo Ejecutando
sleep 5
