CREATE KEYSPACE IF NOT EXISTS <<KEYSPACE>> WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE TABLE IF NOT EXISTS <<KEYSPACE>>.configurations (database_version COUNTER, PRIMARY KEY (database_version));
INSERT INTO <<KEYSPACE>>.configurations (database_version) VALUES (1);