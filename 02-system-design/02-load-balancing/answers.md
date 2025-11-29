# Load Balancing Answers

1.  **High Availability**: The LB performs health checks on backend servers. If a server fails a check, the LB stops sending traffic to it and reroutes to healthy servers, preventing downtime for users.
2.  **L4 vs L7**:
    - **L4**: Works at Transport layer (TCP/UDP). Routes based on IP/Port. Faster, doesn't inspect data.
    - **L7**: Works at Application layer (HTTP). Routes based on URL, Cookies, Headers. Can do SSL termination and smart routing (e.g., /video to video servers).
3.  **Sticky Sessions**: Ensures a client's requests always go to the same server. Implemented via **IP Hashing** or by setting a **Cookie** in the client's browser that the LB reads.
4.  **LB Failure**: The LB is a Single Point of Failure (SPOF). To prevent this, use a pair of LBs in **Active-Passive** mode. The Passive LB monitors the Active one and takes over the Virtual IP (VIP) if the Active one fails.
5.  **Least Connections vs Round Robin**: Use **Least Connections** when requests have varying processing times (e.g., a chat session vs a simple API call). Round Robin is better when all requests are roughly equal in load.
