Sorted Sets are just like regular sets, but instead of being sorted by the value itself, they're sorted by a score value associated with each member.

Here's how you create a sorted set, get the items inside, and get the size of a sorted set with the commands `ZADD`, `ZRANGE`, and `ZCARD`:
```bash
127.0.0.1:6379> ZADD users 1 Berat 2 Abtin 3 Raim
(integer) 3
127.0.0.1:6379> ZRANGE users 0 -1
1) "Berat"
2) "Abtin"
3) "Raim"
127.0.0.1:6379> ZRANGE users 0 -1 withscores
1) "Berat"
2) "1"
3) "Abtin"
4) "2"
5) "Raim"
6) "3"
127.0.0.1:6379> ZCARD users
(integer) 3
```

You can also count filtered users by the specified minimum and maximum scores with the `ZCOUNT` command:
```bash
127.0.0.1:6379> ZCOUNT users -inf +inf
(integer) 3
127.0.0.1:6379> ZCOUNT users 2 3
(integer) 2
```

To remove an item you can use `ZREM`:
```bash
127.0.0.1:6379> ZREM users Berat
(integer) 1
127.0.0.1:6379> ZRANGE users 0 -1
1) "Abtin"
2) "Raim"
```

Here's how you get the reverse ordered set with the `ZREVRANGE` command:
```bash
127.0.0.1:6379> ZRANGE users 0 -1 WITHSCORES
 1) "Abtin"
 2) "2"
 3) "Another"
 4) "3"
 5) "Eren"
 6) "3"
 7) "Heyo"
 8) "3"
 9) "Raim"
10) "3"
11) "Sara"
12) "4"
127.0.0.1:6379> ZREVRANGE users 0 -1 WITHSCORES
 1) "Sara"
 2) "4"
 3) "Raim"
 4) "3"
 5) "Heyo"
 6) "3"
 7) "Eren"
 8) "3"
 9) "Another"
10) "3"
11) "Abtin"
12) "2"
```

This is how you get the score of a particular value with `ZSCORE`:
```bash
127.0.0.1:6379> ZSCORE users Abtin
"2"
```

If you need to filter scores, you can use either `ZRANGEBYSCORE` or `ZREVRANGEBYSCORE`:
```bash
127.0.0.1:6379> ZRANGEBYSCORE users 2 3 WITHSCORES
 1) "Abtin"
 2) "2"
 3) "Another"
 4) "3"
 5) "Eren"
 6) "3"
 7) "Heyo"
 8) "3"
 9) "Raim"
10) "3"
127.0.0.1:6379> ZREVRANGEBYSCORE users 3 2 WITHSCORES
 1) "Raim"
 2) "3"
 3) "Heyo"
 4) "3"
 5) "Eren"
 6) "3"
 7) "Another"
 8) "3"
 9) "Abtin"
10) "2"
```

You can also increment or decrease the score using `ZINCRBY`:
```bash
127.0.0.1:6379> ZINCRBY users 2 Abtin
"4"
127.0.0.1:6379> ZINCRBY users -1 Abtin
"3"
```

You can remove users by using a range with the `ZREMRANGEBYSCORE` command:
```bash
127.0.0.1:6379> ZREMRANGEBYSCORE users 2 3
(integer) 5
127.0.0.1:6379> ZRANGE users 0 -1
1) "Sara"
```

Just like with scores, you can remove by rank as well using the `ZREMRANGEBYRANK` command:
```bash
127.0.0.1:6379> ZADD users 1 Ali 2 Berat 3 Hubeyb 4 Emir 5 Kemal
(integer) 5
127.0.0.1:6379> ZRANGE users 0 -1 WITHSCORES
 1) "Ali"
 2) "1"
 3) "Berat"
 4) "2"
 5) "Hubeyb"
 6) "3"
 7) "Emir"
 8) "4"
 9) "Kemal"
10) "5"
127.0.0.1:6379> ZREMRANGEBYRANK users 1 3 # index
(integer) 3
127.0.0.1:6379> ZRANGE users 0 -1 WITHSCORES
1) "Ali"
2) "1"
3) "Kemal"
4) "5"
```
