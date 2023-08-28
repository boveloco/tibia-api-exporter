CREATE TABLE IF NOT EXISTS <<KEYSPACE>>.creature_statistics (name TEXT, count INT, day BIGINT, world TEXT, PRIMARY KEY (name));
CREATE INDEX IF NOT EXISTS world_index ON <<KEYSPACE>>.creature_statistics (world)
CREATE INDEX IF NOT EXISTS day_index ON <<KEYSPACE>>.creature_statistics (day)
UPDATE <<KEYSPACE>>.configurations SET database_version = 2 WHERE data_partition = '<<KEYSPACE>>';