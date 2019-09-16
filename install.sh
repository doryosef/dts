#!/bin/bash

if [[ $UID != 0 ]] ; then
    echo "use root privileges to run this script"

    exit 1
fi
#download dts.tgz && 
mkdir -p /opt/dts && \
tar -xzvf dts.tgz -C /opt/dts && \
chown -R pi:pi /opt/dts
chmod u+x dts scripts/* && \
mv -v /opt/dts/dts.service /etc/systemd/system/

echo " -------------------------

now edit the file /opt/dts/conf.ini and then execute

sudo systemctl daemon-reload
sudo systemctl start dts
sudo systemctl enable dts.service
sudo systemctl status dts
"
