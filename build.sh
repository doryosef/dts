#!/bin/bash

go build dts.go && \
chmod u+x scripts/* && \
tar -czvf dts.tgz dts dts.service scripts conf.ini