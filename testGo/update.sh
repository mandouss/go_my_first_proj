#!/bin/sh

server=hz147@vcm-1548.vm.duke.edu
workPath=/home/hz147/go/src/stockPairing/testGo
command="cd $workPath; ./test.sh; exit; bash"

localPath=/Users/haohongzhao/Desktop/590/hw4_stockTransactionMatching/testGo
remotePath=hz147@vcm-1548.vm.duke.edu:/home/hz147/go/src/stockPairing/

rsync --progress --partial --delete-after -avz $localPath $remotePath > trash.log

ssh -o ServerAliveInterval=30 -t $server $command

rm trash.log