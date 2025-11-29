# Proxies Answers

1.  **Forward vs Reverse**:
    - **Forward**: Protects/Hides the **Client**. (Client -> Proxy -> Internet).
    - **Reverse**: Protects/Hides the **Server**. (Internet -> Proxy -> Server).
2.  **Security**: A Reverse Proxy hides the IP addresses of the backend servers. Attackers only see the Proxy's IP. It can also filter out malicious requests (WAF - Web Application Firewall) before they reach the app.
3.  **SSL Termination**: The Proxy handles the heavy math of decrypting HTTPS requests. It sends plain HTTP to the backend servers. This offloads CPU work from the backend servers, allowing them to handle more application logic.
4.  **LB as Reverse Proxy**: Yes, most Load Balancers act as Reverse Proxies because they sit in front of servers and route traffic to them.
5.  **Forward Proxy Example**: A company network blocking access to social media sites. All employee traffic goes through the proxy, which checks the URL and blocks it if it's on a blacklist.
