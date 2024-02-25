#!/bin/bash
Host_Name=`hostname`
Conf_Dir=/opt/hbase/conf
if [ -n ${ZK_Server} ]; then
  sed -i "/zookeeper.quorum/{n;s/.*/  <value>${ZK_Server}<\/value>/g}" ${Conf_Dir}/hbase-site.xml
fi
if [-n ${ZK_Port}]; then
  sed -i "/zookeeper.property.clientPort/{n;s/.*/  <value>${ZK_Port}<\/value>/g}" ${Confi_Dir}/hbase-site.xml
fi

bash /opt/hbase/bin/start-hbase.sh
tail -F /opt/hbase/logs/hbase--master-${Host_Name}.log