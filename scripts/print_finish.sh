#!/bin/bash
tar -cvzf timelapse_$(date +%Y%m%d%H%M).tgz timelapse/ 
echo "finish compressing timelapse"
