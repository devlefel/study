# Caching Answers

1.  **Cache Aside vs Write Through**:
    - **Cache Aside**: App manages cache. Read: Check Cache -> Miss -> Read DB -> Write Cache. Write: Update DB -> Delete Cache. Pros: Only requested data is cached. Cons: First request is slow (miss).
    - **Write Through**: App writes to Cache and DB at the same time. Pros: Data always consistent. Cons: Slower writes, cache might contain unused data.
2.  **LRU (Least Recently Used)**: Removes the item accessed least recently. Implementation: **Doubly Linked List + Hash Map**. Map keys point to nodes. On access, move node to head. On eviction, remove tail.
3.  **Thundering Herd**: When a popular cache key expires, thousands of requests hit the DB simultaneously.
    - **Solution**: Locking (only one process updates cache), or probabilistic early expiration (recompute before it actually expires).
4.  **When NOT to cache**:
    - Data changes frequently (high volatility).
    - Data is rarely requested (low hit rate).
    - Consistency is critical (e.g., bank balance).
5.  **Consistency**:
    - **TTL (Time To Live)**: Auto-expire keys.
    - **Invalidation**: Explicitly delete/update cache on DB write.
