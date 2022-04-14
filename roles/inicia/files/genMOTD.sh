#!/bin/bash
ambiente=$1
rol=$(uname -n | sed '1,$s/^'$ambiente'//g')
sudo chmod 777 /tmp
echo "###########################################################" > /tmp/motd
echo "##       ambiente: $ambiente" >> /tmp/motd
echo "##       hostname: $(hostname)" >> /tmp/motd
echo "##       fqdn    : $(hostname -f)" >> /tmp/motd
echo "###########################################################" >> /tmp/motd
sudo cp /tmp/motd /etc/.
