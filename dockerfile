FROM alpine

COPY /bin/exporter /app/exporter

ENV DB_CASSANDRA_CLUSTERIP DB_CASSANDRA_PAWSSWORD DB_CASSANDRA_USERNAME

CMD ["/app/exporter"]
