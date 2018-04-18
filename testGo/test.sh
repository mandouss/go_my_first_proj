#!/bin/bash

result=result
warning=trash.log
keyword=12345

dbs=("account" "symbol" "accountShare" "contractIDToAccountID" "contract" "executedContract" "cancelledContract")
#      0          1            2                 3                    4          5                    6
query_in_use=0
queries[0]="SELECT * FROM ${dbs[1]};"
queries[1]="SELECT * FROM ${dbs[2]};"




# echo "${queries[$query_in_use]}" | psql -U postgres -d stockuserandorder

go run ./*.go
# psql -U postgres -d stockuserandorder -c "${queries[$query_in_use]}" 2> $warning >> $result
psql -U postgres -d stockuserandorder -c "${queries[0]}" 2> $warning >> $result
psql -U postgres -d stockuserandorder -c "${queries[1]}" 2> $warning >> $result

cat $result
# cat $result | grep $keyword


rm $result $warning


