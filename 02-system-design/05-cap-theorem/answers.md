# CAP Theorem Answers

1.  **CAP**: You can't have it all. In a network of computers, if the network breaks (Partition), you have to decide: do I stop serving requests to make sure data is perfect (Consistency), or do I keep serving requests even if data might be old (Availability)?
2.  **Why not CA?**: Networks fail. Cables get cut, routers crash. If you assume the network never fails (CA), your system will break the moment a partition happens. You _must_ handle Partitions (P).
3.  **AP Example**: A Facebook/Twitter feed. It's okay if I see a post 5 seconds later than you. It's NOT okay if the site crashes just because one data center can't talk to another.
4.  **Eventual Consistency**: A consistency model used in AP systems. It guarantees that if no new updates are made to a given data item, eventually all accesses to that item will return the last updated value.
5.  **PACELC**: CAP only talks about what happens during a failure (Partition). PACELC says: "Even when the system is running normally (Else), you have to choose between Latency (speed) and Consistency." If you want perfect consistency, you have to wait for all nodes to sync (Latency).
