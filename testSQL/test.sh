#!/bin/sh

sqlPath="/home/hz147/go/src/stockPairing/testSQL/"
finPath="$sqlPath/fin"

#psql -U postgres -d stockuserandorder -a -f "$finPath/drop.sql"
#psql -U postgres -d stockuserandorder -a -f "$finPath/create.sql"
#psql -U postgres -d stockuserandorder -a -f "$finPath/insert.sql"
#psql -U postgres -d stockuserandorder -a -f "$sqlPath/delete.sql"
psql -U postgres -d stockuserandorder -a -f "$sqlPath/query.sql"
#psql -U postgres -d stockuserandorder -a -f "$finPath/show.sql"
