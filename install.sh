#!/bin/bash

if [[ $UID != 0 ]] ; then
    echo "use root privileges to run this script"

    exit 1
fi
echo "downloading and installing lastest release"
wget -O dts.tgz https://github.com/doryosef/dts/releases/latest/download/dts.tgz && \
mkdir -vp /opt/dts && \
tar -xzvf dts.tgz -C /opt/dts && \
chown -vR pi:pi /opt/dts
chmod -v u+x /opt/dts/dts /opt/dts/scripts/* && \
mv -v /opt/dts/dts.service /etc/systemd/system/ && \
rm -f dts.tgz

echo " -------------------------

now edit the file /opt/dts/conf.ini and then execute

sudo systemctl daemon-reload
sudo systemctl start dts
sudo systemctl enable dts.service
sudo systemctl status dts
"
