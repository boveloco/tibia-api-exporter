CREATE KEYSPACE IF NOT EXISTS tibia-api WITH REPLICATION = {'class' : 'NetworkTopologyStrategy' : 3};

CREATE TABLE IF NOT EXISTS configurations (database_version int PRIMARY KEY (database_version));

INSERT INTO tibia-api.configurations (database_version) VALUES (1);