FROM harisekhon/hbase:2.1

COPY hbase-site.xml /hbase-2.1.3/conf/

COPY hbase-site.xml /hbase/conf/

RUN chmod 777 /hbase-data

ENTRYPOINT [ "./entrypoint.sh" ]