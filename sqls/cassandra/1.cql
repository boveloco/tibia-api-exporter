CREATE KEYSPACE IF NOT EXISTS <<KEYSPACE>> WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE TABLE IF NOT EXISTS <<KEYSPACE>>.configurations (data_partition TEXT, database_version INT, PRIMARY KEY (data_partition));
INSERT INTO <<KEYSPACE>>.configurations (data_partition, database_version) VALUES ('<<KEYSPACE>>', 1);