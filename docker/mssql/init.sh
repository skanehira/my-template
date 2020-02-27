#!/bin/bash
docker cp init.sql mssql:/home
docker exec -it mssql /opt/mssql-tools/bin/sqlcmd -U sa -P mssqlPassword! -i /home/init.sql

