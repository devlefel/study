# Load Balancing: A Deep Dive

Load Balancing is the practice of distributing computational workloads across multiple resources, such as computers, a computer cluster, network links, central processing units, or disk drives. It is a critical component for high availability and scalability.

![Load Balancer Diagram](../images/load_balancer.png)

## 1. Types of Load Balancers

### Layer 4 (Transport Layer)

Operates at the TCP/UDP level. It sees IP addresses and ports but **not** the content of the packet.

- **Mechanism**: Uses NAT (Network Address Translation) to forward packets.
- **Pros**: Extremely fast, low CPU usage. Can handle millions of concurrent connections.
- **Cons**: Cannot route based on content (e.g., cannot send `/video` to video servers).
- **Examples**: LVS (Linux Virtual Server), HAProxy (TCP mode), AWS Network Load Balancer.

### Layer 7 (Application Layer)

Operates at the HTTP/HTTPS level. It terminates the connection, inspects the request, and opens a new connection to the backend.

- **Mechanism**: Acts as a proxy.
- **Pros**: Smart routing (URL, Cookies, Headers), SSL Termination, WAF (Web Application Firewall), Compression, Rate Limiting.
- **Cons**: Higher latency and CPU usage than L4.
- **Examples**: Nginx, HAProxy (HTTP mode), AWS Application Load Balancer, Envoy.

## 2. Algorithms

### Static Algorithms

- **Round Robin**: Cycles through servers. Simple, but assumes all servers are equal.
- **Weighted Round Robin**: Assigns more requests to powerful servers (e.g., Server A weight=3, Server B weight=1).
- **IP Hash**: Hashes client IP to map to a server. Guarantees **Sticky Sessions** (a user always hits the same server). Crucial for stateful apps, but can lead to uneven distribution.

### Dynamic Algorithms

- **Least Connections**: Sends traffic to the server with fewest active connections. Ideal for long-lived connections (WebSocket, Database).
- **Least Response Time**: Monitors server latency (TTFB) and sends traffic to the fastest responder. Adapts to network lag or slow servers.
- **Resource Based**: Agents on servers report CPU/RAM usage to the LB. The LB avoids overloaded servers.

## 3. Health Checks

A Load Balancer must know if a backend is alive.

- **Active Checks**: The LB periodically pings the server (e.g., `GET /health`). If it fails $N$ times, the server is marked "Unhealthy" and removed from rotation.
- **Passive Checks**: The LB observes real traffic. If a server returns 500 errors or times out, it is temporarily avoided.

## 4. High Availability of the Load Balancer

The LB itself is a Single Point of Failure (SPOF).

- **Active-Passive**: Two LBs. Only one handles traffic. The other sends heartbeats (VRRP/Keepalived). If Active dies, Passive takes over the Virtual IP (VIP).
- **Active-Active**: Both LBs handle traffic. DNS Round Robin distributes traffic between them.

## 5. Global Server Load Balancing (GSLB)

Distributing traffic across data centers (regions).

- **DNS Load Balancing**: DNS returns the IP of the closest data center based on the user's location (GeoDNS).

## Bibliography & Further Reading

1.  **"Site Reliability Engineering"** (Google) - _Chapter 19: Load Balancing at the Frontend_.
2.  **Nginx Architecture Guide**: _Inside NGINX: How We Designed for Performance & Scale_.
3.  **HAProxy Documentation**: _The Art of Load Balancing_.
4.  **"Designing Data-Intensive Applications"** - _Partitioning and Load Balancing_.
