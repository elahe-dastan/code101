# Introduction
Cassandra is optimized for writing large amounts of data very quicly. It writes all incoming data in an append-only 
manner in internal files called SSTables.

Cassandra Compaction is a process of reconciling various copies of data spread across distinct SSTables which is 
performed as a background activity. 

While regular compactions are an integral part of any healthy Cassandra cluster, the way that they are configured can 
very significantly depend on the way a particular table is being used. It is important to take some time when designing 
a Cassandra table schema to think about how each table will be used, and therefore the Cassandra Compaction strategy 
would be most effective.

Although it is possible to change compaction strategies after the table has been created but it will re-write all of the 
data in that table using the new Compaction Strategy.

# Cassandra's Write Path
1. Stores recent writes in memory (Memtable).
2. When enough writes have been made, Cassandra flushed the Memtable to disk. Data on disk is stored in relatively 
simple data structures called Stored String Tables (SSTable). At the most simplified level, as SSTable could be 
described as a stored array of strings.
3. Before writing a new SSTable, Cassandra merges and pre-shorts the data in the Memtable according to Primary Key.
4. The SSTable is written to disk as a single contiguous write operation. SSTables are immutable. If data is updated
regularly, Cassandra may need to read from multiple SSTables to retrieve a single row.
5. Compaction operations occur periodically to re-write and combine SSTables. Compactions prune deleted data and merge 
6. disparate row data into new SSTables in order to reclaim disk space and keep read operations optimized.

# Cassandra Compaction Strategies
1. SizeTiered: When multiple SSTables of a similar size are present.
2. Leveled
3. DateTiered
4. TimeWindow

# SSTables Per Read
The Maximum represents the highest number of SSTables (across all nodes over roughly the last five minutes) that were 
involved in a read request.