# Design a URL Shortener (e.g., TinyURL)

## 1. Requirements

- **Functional**:
  - Given a long URL, return a unique short URL.
  - Clicking the short URL redirects to the original long URL.
  - Users can optionally pick a custom alias.
  - Links expire after a default timespan.
- **Non-Functional**:
  - Highly available.
  - Real-time redirection with minimal latency.
  - Shortened links should be non-guessable (unless custom).

## 2. Capacity Estimation

- **Traffic**: 100M new URLs/month.
- **Read/Write Ratio**: 10:1 (1B redirects/month).
- **Storage**: 500 bytes per URL. 100M \* 500 bytes = 50GB/month. 6TB/10 years.

## 3. API Design

- `createURL(api_key, original_url, custom_alias=None, expire_date=None) -> short_url`
- `deleteURL(api_key, url_key)`

## 4. Database Design

- **Users Table**: UserID, Name, Email, APIKey.
- **URL Table**: Hash (PK), OriginalURL, CreationDate, ExpirationDate, UserID.
- **Choice**: NoSQL (DynamoDB, Cassandra) or Relational (MySQL/Postgres)? NoSQL is good for scalability and simple lookups.

## 5. High Level Design

Client -> Load Balancer -> Web Servers -> Cache (Redis) -> Database

## 6. Key Algorithms

- **Encoding**: MD5/SHA256 hash of the URL? Too long.
- **Base62 Encoding**: Characters [a-z, A-Z, 0-9].
  - 6 characters: $62^6 \approx 56$ billion combinations. Sufficient.
  - How to generate?
    - **Counter**: Use a distributed counter ID generator (e.g., Snowflake). Convert ID to Base62.
    - **Random**: Generate random 6 chars, check DB for collision.

## 7. Bottlenecks & Optimizations

- **Caching**: Cache the most popular URLs (20% of URLs generate 80% of traffic).
- **Database Scaling**: Sharding based on the first character of the hash? Or consistent hashing.
- **Cleanup**: Lazy cleanup (remove when accessed and expired) or scheduled job.
