FROM harisekhon/hbase:2.1

COPY hbase-site.xml /hbase-2.1.3/conf/

COPY hbase-site.xml /hbase/conf/

COPY entrypoint.sh ./

RUN chmod 777 /hbase-data

ENTRYPOINT [ "./entrypoint.sh" ]