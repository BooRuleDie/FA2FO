If you want to create a set or add values to it, you can use `SADD`. You can also see the items of the set by using `SMEMBERS`:
```bash
127.0.0.1:6379> SADD technology Redis Java Python FastAPI Go
(integer) 5
127.0.0.1:6379> SMEMBERS technology
1) "Redis"
2) "Java"
3) "Python"
4) "FastAPI"
5) "Go"
```

You can't add duplicate data to sets, just like sets in programming languages:
```bash
127.0.0.1:6379> SMEMBERS technology
1) "Redis"
2) "Java"
3) "Python"
4) "FastAPI"
5) "Go"
127.0.0.1:6379> SADD technology Redis
(integer) 0
```

To get the length of the set, we can use `SCARD`:
```bash
127.0.0.1:6379> SCARD technology
(integer) 5
```

If you want to check if a particular value is available or not, you can use `SISMEMBER`:
```bash
127.0.0.1:6379> SISMEMBER technology Redis
(integer) 1
127.0.0.1:6379> SISMEMBER technology Django
(integer) 0
```

You can also perform set operations with the commands `SDIFF`, `SINTER` and store the result to a new set with `SINTERSTORE` or `SDIFFSTORE`:
```bash
127.0.0.1:6379> SADD backend python go javascript aws redis postgresql mysql
(integer) 7
127.0.0.1:6379> SADD frontend javascript html css
(integer) 3
127.0.0.1:6379> SMEMBERS backend
1) "python"
2) "go"
3) "javascript"
4) "aws"
5) "redis"
6) "postgresql"
7) "mysql"
127.0.0.1:6379> SMEMBERS frontend
1) "javascript"
2) "html"
3) "css"
127.0.0.1:6379> SDIFF backend frontend
1) "aws"
2) "redis"
3) "go"
4) "postgresql"
5) "python"
6) "mysql"
127.0.0.1:6379> SINTER backend frontend
1) "javascript"
127.0.0.1:6379> SINTERSTORE common_techs backend frontend
(integer) 1
127.0.0.1:6379> SMEMBERS common_techs
1) "javascript"
```

You can also get the union of multiple sets with the `SUNION` command, and store it to a new set with `SUNIONSTORE`:
```bash
127.0.0.1:6379> SUNION backend frontend
1) "postgresql"
2) "aws"
3) "python"
4) "html"
5) "go"
6) "redis"
7) "mysql"
8) "javascript"
9) "css"
127.0.0.1:6379> SUNIONSTORE all_techs backend frontend
(integer) 9
127.0.0.1:6379> SMEMBERS all_techs
1) "python"
2) "go"
3) "javascript"
4) "aws"
5) "redis"
6) "postgresql"
7) "mysql"
8) "html"
9) "css"
```