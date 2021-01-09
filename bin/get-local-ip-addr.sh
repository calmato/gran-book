#!/bin/bash

IP=`ifconfig | grep inet | cut -d " " -f2 | grep "[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*" | awk 'NR==2'`
sed -i".bak" -e "/LOCAL_IP_ADDRESS=/d" .env
echo "LOCAL_IP_ADDRESS=$IP" >> .env
