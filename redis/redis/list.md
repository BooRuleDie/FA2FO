# List

You can fetch all keys with the `KEYS` command:
```bash
127.0.0.1:6379> KEYS *
1) "pi"
2) "username"
3) "firstname"
4) "lastname"
5) "count"
```

You can also remove all current keys with the `FLUSHALL` command:
```bash
127.0.0.1:6379> KEYS *
1) "pi"
2) "username"
3) "firstname"
4) "lastname"
5) "count"
127.0.0.1:6379> FLUSHALL
OK
127.0.0.1:6379> KEYS *
(empty array)
```

To create a list, you can use the `LPUSH` or `RPUSH` commands. You can iterate through list items via the `LRANGE` command. `LPUSH` prepends an item to the start of the list and `RPUSH` appends an item to the end of the list. Similarly, you can remove items with `RPOP` and `LPOP` commands. You can also get the length of the list via the `LLEN` command:
```bash
127.0.0.1:6379> LPUSH countries USA
(integer) 1
127.0.0.1:6379> LPUSH countries India
(integer) 2
127.0.0.1:6379> LPUSH countries Turkey
(integer) 3
127.0.0.1:6379> LRANGE countries 0 -1
1) "Turkey"
2) "India"
3) "USA"
127.0.0.1:6379> RPUSH countries China 
(integer) 4
127.0.0.1:6379> LRANGE countries 0 -1
1) "Turkey"
2) "India"
3) "USA"
4) "China"
127.0.0.1:6379> LLEN countries
(integer) 4
127.0.0.1:6379> LPOP countries
"Turkey"
127.0.0.1:6379> LRANGE countries 0 -1
1) "India"
2) "USA"
3) "China"
127.0.0.1:6379> RPOP countries 
"China"
127.0.0.1:6379> LRANGE countries 0 -1
1) "India"
2) "USA"
```

To change an item's value, you can use the `LSET` command:
```bash
127.0.0.1:6379> LRANGE countries 0 -1
1) "India"
2) "USA"
127.0.0.1:6379> LSET countries 1 Germany
OK
127.0.0.1:6379> LRANGE countries 0 -1
1) "India"
2) "Germany"
```

If you need to insert a new value before or after an item, you can do that with `LINSERT`:
```bash
127.0.0.1:6379> LRANGE countries 0 -1
1) "India"
2) "Germany"
3) "Japan"
4) "China"
127.0.0.1:6379> LINSERT countries AFTER Japan Taiwan
(integer) 5
127.0.0.1:6379> LRANGE countries 0 -1
1) "India"
2) "Germany"
3) "Japan"
4) "Taiwan"
5) "China"
```

To get the value of a key by its index, you can use `LINDEX`:
```bash
127.0.0.1:6379> LRANGE countries 0 -1
1) "India"
2) "Germany"
3) "Japan"
4) "Taiwan"
5) "China"
127.0.0.1:6379> LINDEX countries 3
"Taiwan"
```

If you want to check if a list exists before adding a value, you can use `LPUSHX` or `RPUSHX`:
```bash
127.0.0.1:6379> LPUSHX notexistedlist new_value
(integer) 0
127.0.0.1:6379> RPUSHX notexistedlist new_value
(integer) 0
127.0.0.1:6379> RPUSHX countries Turkey
(integer) 6
```

For sorting the values inside a list, you can use the `SORT` command:
```bash
127.0.0.1:6379> SORT countries ALPHA
1) "China"
2) "Germany"
3) "India"
4) "Japan"
5) "Taiwan"
6) "Turkey"
127.0.0.1:6379> SORT countries DESC ALPHA
1) "Turkey"
2) "Taiwan"
3) "Japan"
4) "India"
5) "Germany"
6) "China"
```

If you need to remove an item from the list and want to block the connection until the data becomes available or a timeout occurs, you can use `BLPOP` or `BRPOP`:
```bash
127.0.0.1:6379> LRANGE countries 0 -1
1) "USA"
2) "China"
3) "Turkey"
127.0.0.1:6379> BLPOP countries 3
1) "countries"
2) "USA"
127.0.0.1:6379> BLPOP countries 3
1) "countries"
2) "China"
127.0.0.1:6379> BLPOP countries 3
1) "countries"
2) "Turkey"
127.0.0.1:6379> BLPOP countries 3
(nil)
```