# Hash

Hashes or Hash data structures in Redis can be considered similar to dictionaries in Python or objects in JSON. They provide a convenient way to store field-value pairs, making them ideal for representing objects, especially when you need to access or modify individual fields without retrieving the entire structure. Hashes are memory efficient and allow atomic operations on individual fields.

To create a hash you can use `HSET`, and for retrieving keys, values or all fields you can use `HKEYS`, `HVALS`, `HGETALL`:
```bash
127.0.0.1:6379> HSET user1 name eren
(integer) 1
127.0.0.1:6379> HSET user1 surname burulday
(integer) 1
127.0.0.1:6379> HSET user1 age 24
(integer) 1
127.0.0.1:6379> HKEYS user1
1) "name"
2) "surname"
3) "age"
127.0.0.1:6379> HVALS user1
1) "eren"
2) "burulday"
3) "24"
127.0.0.1:6379> HGETALL user1
1) "name"
2) "eren"
3) "surname"
4) "burulday"
5) "age"
6) "24"
```

You can also check if a particular field exists by using `HEXISTS` and check the length by using `HLEN`:
```bash
127.0.0.1:6379> HLEN user1
(integer) 3
127.0.0.1:6379> HEXISTS user1 name
(integer) 1
127.0.0.1:6379> HEXISTS user1 unexistedfield
(integer) 0
```

You can set or fetch multiple fields in a hash key with `HMSET`, `HMGET` commands:
```bash
127.0.0.1:6379> HMSET user2 name berat surname karatas age 23
OK
127.0.0.1:6379> HMGET user1 name surname
1) "eren"
2) "burulday"
```

You can also increase an integer or float field's value using `HINCRBY`, `HINCRBYFLOAT`, and remove a particular field using `HDEL`:
```bash
127.0.0.1:6379> HINCRBY user1 age 10
(integer) 34
127.0.0.1:6379> HINCRBYFLOAT user1 age 3.5
"37.5"
127.0.0.1:6379> HDEL user1 age
(integer) 1
127.0.0.1:6379> HMGET user1 name surname age
1) "eren"
2) "burulday"
3) (nil)
```

You can also see the length of a string field with `HSTRLEN` and set a new field if it doesn't already exist using `HSETNX`:
```bash
127.0.0.1:6379> HSTRLEN user1 name
(integer) 4
127.0.0.1:6379> HSETNX user1 name Ahmet
(integer) 0
127.0.0.1:6379> HDEL user1 name
(integer) 1
127.0.0.1:6379> HSETNX user1 name Ahmet
(integer) 1
127.0.0.1:6379> HMGET user1 name surname
1) "Ahmet"
2) "burulday"
```