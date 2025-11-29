# Sharding & Replication Answers

1.  **Sharding vs Replication**:
    - **Replication**: Copies the _entire_ database to multiple servers. Good for Read availability.
    - **Sharding**: Splits the database into _unique chunks_ across servers. Good for Write availability and storage limits.
2.  **Replication Lag**: The time delay between a write on the Master and the update appearing on the Slave. Users might read stale data (Eventual Consistency).
3.  **Consistent Hashing**: A technique where adding/removing a server only affects $1/N$ keys, not all keys. Simple Modulo Hashing (`hash % N`) requires moving nearly all keys when $N$ changes.
4.  **Downsides of Sharding**:
    - Complex queries (Joins across shards).
    - Operational complexity (managing many DBs).
    - Uneven data distribution (some shards bigger than others).
5.  **Hot Key**: A specific key (e.g., Justin Bieber's profile) gets too much traffic.
    - **Solution**: Cache the hot key. Add random suffix to key to spread writes (if write-heavy).
