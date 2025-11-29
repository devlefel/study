# Scalability Answers

1.  **Vertical vs Horizontal**: Vertical (Scale Up) adds resources (CPU/RAM) to a single node. Horizontal (Scale Out) adds more nodes to the system.
2.  **When to choose Vertical**:
    - When the traffic load is low to moderate.
    - When the application is not designed to be distributed (stateful).
    - For quick, temporary fixes before a proper re-architecture.
3.  **Disadvantages of Horizontal**:
    - Increased complexity (Load Balancing, Data Consistency).
    - Network latency between nodes.
    - Operational overhead (managing more servers).
4.  **Identifying Bottlenecks**:
    - Use monitoring tools (Prometheus, Datadog) to check CPU, RAM, Disk I/O, and Network usage.
    - Load testing (JMeter, Locust) to see where the system breaks.
    - Profiling the application code.
5.  **Scaling Database**:
    - **Vertical**: Upgrade the DB server.
    - **Read Replicas**: Offload read traffic to replicas (Horizontal).
    - **Sharding**: Split data across multiple servers (Horizontal).
    - **Caching**: Reduce load by caching frequent queries.
