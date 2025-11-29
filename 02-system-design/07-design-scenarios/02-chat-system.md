# Design a Chat System (e.g., WhatsApp, Messenger)

## 1. Requirements

- **Functional**:
  - One-on-one chat.
  - Group chat.
  - Online/Offline status.
  - Sent/Delivered/Read receipts.
- **Non-Functional**:
  - Low latency.
  - High consistency (messages ordered correctly).
  - High availability.

## 2. Capacity Estimation

- 500M Daily Active Users (DAU).
- 40 messages per user/day -> 20B messages/day.

## 3. High Level Design

- **Protocol**: HTTP is too slow (polling). Use **WebSockets** for persistent bi-directional connection.
- **Services**:
  - **Chat Service**: Handles message sending/receiving.
  - **Presence Service**: Manages online/offline status.
  - **User Service**: User profiles, auth.
  - **Notification Service**: Push notifications (APNS/FCM) for offline users.

## 4. Data Storage

- **Messages**: Huge volume. Horizontal scaling is a must.
  - **HBase / Cassandra**: Wide-column stores are great for time-series chat data. Key: `ChatID`, Sort Key: `Timestamp`.
- **User Data**: Relational DB (MySQL) for profiles and friends.

## 5. Detailed Component Design

- **Message Flow**:
  1.  User A sends message to Chat Server.
  2.  Chat Server saves to DB.
  3.  Chat Server finds which server User B is connected to (via Redis/Zookeeper).
  4.  Forwards message to that server.
  5.  Server pushes to User B via WebSocket.
- **Group Chat**:
  - Limit group size (e.g., 256).
  - Server iterates through members and pushes message to each.

## 6. Bottlenecks

- **Connection Limit**: A single server can handle ~65k concurrent connections (TCP port limit). Need many chat servers behind a Load Balancer.
- **Database Write Throughput**: Shard by `ChatID` or `UserID`.
