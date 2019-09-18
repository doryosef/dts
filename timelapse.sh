#!/bin/bash

function downloadFromPi {
sftp pi@$1 << !
get /opt/dts/timelapse_*
rm /opt/dts/timelapse_*
bye
!
}

function createMovie {
counter=1;
for f in $1/timelapse/*.jpg
do
	echo changeing $f to:
	echo img-$(printf "%08d" $counter).jpg
	mv $f $1/img-$(printf "%08d" $counter).jpg
	(( counter ++))
done
powershell ffmpeg -r 16 -an -i $1/img-%8d.jpg $2.mp4 && rm -rf $1
}


host=0pi.lan
############# main #############
downloadFromPi $host

# extract from tgz 
for f in timelapse_*.tgz
 do 
 if [ ! -f $f ]; then
   continue
 fi
 echo "-=-=-=-=-=- $f -=-=-=-=-=-"
 time_stamp=${f:10:-4}
 temp_dir=temp_$time_stamp
 mkdir $temp_dir
 tar -xzvf $f -C $temp_dir
 createMovie $temp_dir $time_stamp
done