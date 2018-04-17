#!/bin/sh

server=hz147@vcm-1548.vm.duke.edu
workPath=/home/hz147/go/src/stockPairing/testSQL
command="cd $workPath; ./test.sh; exit; bash"

localPath=/Users/haohongzhao/Desktop/590/hw4_stockTransactionMatching/testSQL
remotePath=hz147@vcm-1548.vm.duke.edu:/home/hz147/go/src/stockPairing/

rsync --progress --partial --delete-after -avz $localPath $remotePath

ssh -o ServerAliveInterval=30 -t $server $command