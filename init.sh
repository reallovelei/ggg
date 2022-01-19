#!/bin/bash
FOLDER=$(cd `dirname $0`; pwd)
echo $FOLDER"/storage/log"
# 日志用
mkdir -p $FOLDER"/storage/log"
# cron会用到
mkdir -p $FOLDER"/storage/runtime"
