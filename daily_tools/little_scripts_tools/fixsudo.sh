#!/bin/bash

echo -e "sleep 5;sudo sed -i 's/^appadmin/#&/g' /etc/sudoers" > /tmp/removesudo.sh
nohup sudo /bin/bash /tmp/removesudo.sh &
