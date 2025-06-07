# Transaction

Transactions in Redis are used for atomicity. If you need to perform multiple operations as if they were just one operation, transactions should be used. The way you use them is similar to transactions in relational databases. Additionally, the `WATCH` command allows you to handle race conditions and similar concurrency problems (for example, optimistic concurrency can be implemented).

When you use `MULTI`, it starts queueing your requests and after you execute `EXEC`, it executes your operations. If you need to cancel the transaction, you can use the `DISCARD` command:
```bash
127.0.0.1:6379> MULTI
OK
127.0.0.1:6379(TX)> SET a 1
QUEUED
127.0.0.1:6379(TX)> SET b 2
QUEUED
127.0.0.1:6379(TX)> GET a
QUEUED
127.0.0.1:6379(TX)> SET a 3
QUEUED
127.0.0.1:6379(TX)> GET a
QUEUED
127.0.0.1:6379(TX)> EXEC
1) OK
2) OK
3) "1"
4) OK
5) "3"
127.0.0.1:6379> MULTI
OK
127.0.0.1:6379(TX)> SET a 5
QUEUED
127.0.0.1:6379(TX)> SET b 10
QUEUED
127.0.0.1:6379(TX)> DISCARD
OK
127.0.0.1:6379> GET a
"3"
127.0.0.1:6379> GET b
"2"
```

As explained above, `WATCH` allows you to implement optimistic concurrency models by letting you know if the key you're trying to change has already been changed by any other client:

```bash
# Client 1                          | # Client 2
127.0.0.1:6379> SET balance 100    | 127.0.0.1:6379> MULTI
OK                                 | OK
127.0.0.1:6379> WATCH balance      | 127.0.0.1:6379(TX)> SET balance 150  
OK                                 | QUEUED
127.0.0.1:6379> SET balance 500    | 127.0.0.1:6379(TX)> GET balance
OK                                 | QUEUED
                                   | 127.0.0.1:6379(TX)> EXEC
                                   | 1) OK
                                   | 2) "150"
                                   | 127.0.0.1:6379> WATCH balance
                                   | OK
                                   | 127.0.0.1:6379> MULTI
                                   | OK
                                   | 127.0.0.1:6379(TX)> SET balance 200
                                   | QUEUED
                                   | 127.0.0.1:6379(TX)> GET balance
                                   | QUEUED
                                   | 127.0.0.1:6379(TX)> EXEC
                                   | (nil) # <---- ERROR
```
