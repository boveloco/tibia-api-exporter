CREATE KEYSPACE IF NOT EXISTS tibia_api WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE TABLE IF NOT EXISTS tibia_api.configurations (database_version int, PRIMARY KEY (database_version));
INSERT INTO tibia_api.configurations (database_version) VALUES (1);