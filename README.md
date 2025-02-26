Creates concurrent GET and POST requests to guestbook server.

Sample Run:

```

jelumala@Jayanthans-MacBook-Pro my_guestbook_client % go build .                    
jelumala@Jayanthans-MacBook-Pro my_guestbook_client % ./my_guestbook_client         
Total expected visits: 10000, Visits got: 1
Total expected visits: 10000, Visits got: 7
Total expected visits: 10000, Visits got: 7
Total expected visits: 10000, Visits got: 9
Total expected visits: 10000, Visits got: 9
Total expected visits: 10000, Visits got: 9
Total expected visits: 10000, Visits got: 9
Total expected visits: 10000, Visits got: 10
Total expected visits: 10000, Visits got: 11
Total expected visits: 10000, Visits got: 12
Total expected visits: 10000, Visits got: 12
Total expected visits: 10000, Visits got: 13
Total expected visits: 10000, Visits got: 14
Total expected visits: 10000, Visits got: 14
Total expected visits: 10000, Visits got: 15
Total expected visits: 10000, Visits got: 18
Total expected visits: 10000, Visits got: 18
Total expected visits: 10000, Visits got: 18
Total expected visits: 10000, Visits got: 19
Total expected visits: 10000, Visits got: 24
Total expected visits: 10000, Visits got: 25
Total expected visits: 10000, Visits got: 26
Total expected visits: 10000, Visits got: 26
Total expected visits: 10000, Visits got: 26
Total expected visits: 10000, Visits got: 26
Total expected visits: 10000, Visits got: 26
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded
error doing get request, possible DEADLOCK: Get "http://localhost:8089/": context deadline exceeded



```
