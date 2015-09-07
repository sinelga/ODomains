#!/bin/bash

#curl -X PUT "https://api.digitalocean.com/v2/domains/alastomia.com/records/1288889" -d '{"data":"104.131.100.40"}' -H "Authorization: Bearer $TOKEN"


awk -F "," '{print "curl -X PUT \"https://api.digitalocean.com/v2/domains/"$1"/records/"$2"\" -d \{\"data\":\"104.131.91.85\"\} -H \"Authorization: Bearer $TOKEN\" -H \"Content-Type: application/json\""}'  output.csv >tmp.bash

