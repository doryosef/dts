#!/bin/bash
snapshot=timelapse/$(date +"%Y%m%d%H%M%S").jpg
raspistill -bm -o $snapshot && curl -X POST http://localhost:8000/snapshot -d '{"name":"'$snapshot'"}'

