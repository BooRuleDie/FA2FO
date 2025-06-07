You can simply `SET` and `GET` string values with the following commands:
```bash
127.0.0.1:6379> SET name booruledie
OK
127.0.0.1:6379> GET name
"booruledie"
```

If you need to fetch a substring from a value, you can do that with `GETRANGE`:
```bash
127.0.0.1:6379> SET email booruledie@gmail.com
OK
127.0.0.1:6379> GET email
"booruledie@gmail.com"
127.0.0.1:6379> GETRANGE email 0 9
"booruledie"
```

If you need to set multiple key-value pairs, you can do it with the `MSET` command and get multiple values via `MGET`:
```bash
127.0.0.1:6379> MSET firstname Katya lastname Gurbanova username hulolo
OK
127.0.0.1:6379> MGET firstname lastname username
1) "Katya"
2) "Gurbanova"
3) "hulolo"
```

You can also get the length of a value for a particular key with the `STRLEN` command:
```bash
127.0.0.1:6379> SET email hulolo@gmail.com
OK
127.0.0.1:6379> GET email
"hulolo@gmail.com"
127.0.0.1:6379> STRLEN email
(integer) 16
```

To update a value, you can use the `SET` command:
```bash
127.0.0.1:6379> SET username badusername
OK
127.0.0.1:6379> GET username
"badusername"
127.0.0.1:6379> SET username booruledie
OK
127.0.0.1:6379> GET username
"booruledie"
```

You can also perform integer increment and decrement operations via `INCR`, `DECR`, `INCRBY`, `DECRBY`:
```bash
127.0.0.1:6379> SET count 1
OK
127.0.0.1:6379> GET count
"1"
127.0.0.1:6379> INCR count
(integer) 2
127.0.0.1:6379> INCRBY count 5
(integer) 7
127.0.0.1:6379> DECR count
(integer) 6
127.0.0.1:6379> DECRBY count 3
(integer) 3
```

If you're working with floating-point values, you should use `INCRBYFLOAT`:
```bash
127.0.0.1:6379> SET pi 3.14
OK
127.0.0.1:6379> INCRBYFLOAT pi 0.1
"3.24"
127.0.0.1:6379> INCRBYFLOAT pi -0.2
"3.04"
```

If you have a key-value pair and want to remove the value after a particular time period, you can use the `EXPIRE` command:
```bash
127.0.0.1:6379> SET name rahim
OK
127.0.0.1:6379> EXPIRE name 10
(integer) 1
127.0.0.1:6379> TTL name
(integer) 8
127.0.0.1:6379> TTL name
(integer) 6
127.0.0.1:6379> TTL name
(integer) 4
127.0.0.1:6379> TTL name
(integer) -2
127.0.0.1:6379> GET name
(nil)
```

You can also set the expiration time when creating the key-value pair with `SETEX`:
```bash
127.0.0.1:6379> SETEX email 5 temperkan@gmail.com
OK
127.0.0.1:6379> GET email
"temperkan@gmail.com"
127.0.0.1:6379> TTL email
(integer) -2
127.0.0.1:6379> GET email
(nil)
```