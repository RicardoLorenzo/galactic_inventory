#!/bin/bash
CHART_NAME="server"

DEPLOYED=$(helm list | awk '{ print $1 }' | grep "^application$" | wc -l)

if [ $DEPLOYED == 1 ]; then
   helm upgrade application ../helm/server --set db.server="<server_name>" --set db.port="<server_port>" --set db.user="<user>" --set db.password="<password>" --set db.database="<database>"
else
   helm install application ../helm/server --set db.server="<server_name>" --set db.port="<server_port>" --set db.user="<user>" --set db.password="<password>" --set db.database="<database>"
fi