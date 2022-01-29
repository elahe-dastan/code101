## A little on Cassandra
Batches in Cassandra are often mistaken as a performance optimization. They can be but only in rare cases.

### Unlogged Batch
In a good example of an unlogged batch, all the inserts share the same partition key, so it effectively becomes one 
insert.

If the inserts have different partition keys the the unlogged batches require the coordinator to do all the work of 
managing the inserts, and will make a single node do more work. Worse if the partion keys are owned by other nodes then 
the coordinator node has an extra network hop to manage as well

### Logged Batch
Logged batch puts a fair amount of work on the coordinator node and cluster as a whole. Therefore the primary use case 
of a logged batch is when you need to keep tables in sync with one another, and NOT performance. 