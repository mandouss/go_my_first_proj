-- display timestamp
select name, to_char(createddate, ''yyyymmdd hh:mi:ss tt') as created_date
from "Group"