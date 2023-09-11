CREATE TABLE IF NOT EXISTS <<KEYSPACE>>.executions (day DATE, success BOOLEAN , PRIMARY KEY (day));
CREATE TABLE IF NOT EXISTS <<KEYSPACE>>.creature_statistics (name TEXT, count INT, day DATE, world TEXT, PRIMARY KEY (name, day));
CREATE INDEX IF NOT EXISTS world_index ON <<KEYSPACE>>.creature_statistics (world)
CREATE INDEX IF NOT EXISTS day_index ON <<KEYSPACE>>.creature_statistics (day)
UPDATE <<KEYSPACE>>.configurations SET database_version = 2 WHERE data_partition = '<<KEYSPACE>>';