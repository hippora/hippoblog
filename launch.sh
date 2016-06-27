#! /bin/bash

pid=`ps -U hippo -o pid= -o cmd= | grep hippoblog | grep -v grep | awk '{print $1}'`;if [ -n "$pid" ] ; then kill -9 ${pid}; fi;
nohup ./hippoblog > hippoblog.log 2>&1 &
