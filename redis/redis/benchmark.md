# Benchmark

Redis includes a powerful built-in benchmarking utility called `redis-benchmark`. This flexible tool allows you to customize key testing parameters including the number of concurrent clients, data payload sizes, and total request volume. The benchmark thoroughly evaluates performance across Redis' full command set, giving you comprehensive insights into your deployment's capabilities.

## Example Output
```bash
redis-benchmark -n 1000 -d 10000 -c 50
 
PING_INLINE: rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                  
====== PING_INLINE ======
  1000 requests completed in 0.06 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.455 milliseconds (cumulative count 1)
50.000% <= 1.735 milliseconds (cumulative count 512)
75.000% <= 1.879 milliseconds (cumulative count 752)
87.500% <= 2.015 milliseconds (cumulative count 876)
93.750% <= 2.223 milliseconds (cumulative count 940)
96.875% <= 3.239 milliseconds (cumulative count 969)
98.438% <= 3.839 milliseconds (cumulative count 985)
99.219% <= 4.143 milliseconds (cumulative count 993)
99.609% <= 4.303 milliseconds (cumulative count 997)
99.805% <= 4.415 milliseconds (cumulative count 999)
99.902% <= 4.431 milliseconds (cumulative count 1000)
100.000% <= 4.431 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.200% <= 0.503 milliseconds (cumulative count 2)
0.500% <= 0.607 milliseconds (cumulative count 5)
0.600% <= 0.703 milliseconds (cumulative count 6)
0.900% <= 0.807 milliseconds (cumulative count 9)
1.100% <= 0.903 milliseconds (cumulative count 11)
1.500% <= 1.007 milliseconds (cumulative count 15)
1.800% <= 1.103 milliseconds (cumulative count 18)
2.400% <= 1.207 milliseconds (cumulative count 24)
4.000% <= 1.303 milliseconds (cumulative count 40)
7.300% <= 1.407 milliseconds (cumulative count 73)
15.600% <= 1.503 milliseconds (cumulative count 156)
28.100% <= 1.607 milliseconds (cumulative count 281)
45.400% <= 1.703 milliseconds (cumulative count 454)
64.500% <= 1.807 milliseconds (cumulative count 645)
77.900% <= 1.903 milliseconds (cumulative count 779)
87.000% <= 2.007 milliseconds (cumulative count 870)
91.600% <= 2.103 milliseconds (cumulative count 916)
96.500% <= 3.103 milliseconds (cumulative count 965)
99.100% <= 4.103 milliseconds (cumulative count 991)
100.000% <= 5.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 15873.02 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.792     0.448     1.735     2.351     4.015     4.431
 
====== PING_MBULK ======
  1000 requests completed in 0.06 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.431 milliseconds (cumulative count 1)
50.000% <= 1.775 milliseconds (cumulative count 504)
75.000% <= 1.959 milliseconds (cumulative count 755)
87.500% <= 2.127 milliseconds (cumulative count 876)
93.750% <= 2.335 milliseconds (cumulative count 940)
96.875% <= 2.591 milliseconds (cumulative count 969)
98.438% <= 2.839 milliseconds (cumulative count 985)
99.219% <= 3.215 milliseconds (cumulative count 993)
99.609% <= 3.407 milliseconds (cumulative count 997)
99.805% <= 3.479 milliseconds (cumulative count 999)
99.902% <= 3.527 milliseconds (cumulative count 1000)
100.000% <= 3.527 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.200% <= 0.503 milliseconds (cumulative count 2)
0.500% <= 0.607 milliseconds (cumulative count 5)
0.800% <= 0.703 milliseconds (cumulative count 8)
1.000% <= 0.903 milliseconds (cumulative count 10)
1.300% <= 1.007 milliseconds (cumulative count 13)
1.400% <= 1.103 milliseconds (cumulative count 14)
1.800% <= 1.207 milliseconds (cumulative count 18)
4.500% <= 1.303 milliseconds (cumulative count 45)
9.700% <= 1.407 milliseconds (cumulative count 97)
18.400% <= 1.503 milliseconds (cumulative count 184)
30.100% <= 1.607 milliseconds (cumulative count 301)
40.700% <= 1.703 milliseconds (cumulative count 407)
55.600% <= 1.807 milliseconds (cumulative count 556)
69.900% <= 1.903 milliseconds (cumulative count 699)
79.700% <= 2.007 milliseconds (cumulative count 797)
86.200% <= 2.103 milliseconds (cumulative count 862)
99.000% <= 3.103 milliseconds (cumulative count 990)
100.000% <= 4.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 15625.00 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.792     0.424     1.775     2.407     3.095     3.527
 
====== SET ======
  1000 requests completed in 0.08 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.687 milliseconds (cumulative count 1)
50.000% <= 2.087 milliseconds (cumulative count 502)
75.000% <= 2.303 milliseconds (cumulative count 751)
87.500% <= 2.519 milliseconds (cumulative count 876)
93.750% <= 2.871 milliseconds (cumulative count 938)
96.875% <= 3.167 milliseconds (cumulative count 969)
98.438% <= 3.431 milliseconds (cumulative count 985)
99.219% <= 3.655 milliseconds (cumulative count 993)
99.609% <= 3.783 milliseconds (cumulative count 997)
99.805% <= 3.887 milliseconds (cumulative count 999)
99.902% <= 4.631 milliseconds (cumulative count 1000)
100.000% <= 4.631 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 0.703 milliseconds (cumulative count 1)
0.200% <= 0.807 milliseconds (cumulative count 2)
0.600% <= 0.903 milliseconds (cumulative count 6)
1.100% <= 1.007 milliseconds (cumulative count 11)
1.600% <= 1.103 milliseconds (cumulative count 16)
1.700% <= 1.207 milliseconds (cumulative count 17)
2.200% <= 1.303 milliseconds (cumulative count 22)
2.600% <= 1.407 milliseconds (cumulative count 26)
3.800% <= 1.503 milliseconds (cumulative count 38)
5.800% <= 1.607 milliseconds (cumulative count 58)
10.300% <= 1.703 milliseconds (cumulative count 103)
16.900% <= 1.807 milliseconds (cumulative count 169)
27.400% <= 1.903 milliseconds (cumulative count 274)
39.400% <= 2.007 milliseconds (cumulative count 394)
52.600% <= 2.103 milliseconds (cumulative count 526)
96.200% <= 3.103 milliseconds (cumulative count 962)
99.900% <= 4.103 milliseconds (cumulative count 999)
100.000% <= 5.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 13157.90 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.130     0.680     2.087     2.999     3.503     4.631
 
GET: rps=670.6 (overall: 9388.9) avg_msec=3.816 (overall: 3.816)
                                                                 
====== GET ======
  1000 requests completed in 0.08 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.919 milliseconds (cumulative count 1)
50.000% <= 3.007 milliseconds (cumulative count 503)
75.000% <= 3.607 milliseconds (cumulative count 750)
87.500% <= 4.127 milliseconds (cumulative count 876)
93.750% <= 4.591 milliseconds (cumulative count 938)
96.875% <= 5.095 milliseconds (cumulative count 969)
98.438% <= 5.631 milliseconds (cumulative count 985)
99.219% <= 5.959 milliseconds (cumulative count 993)
99.609% <= 6.295 milliseconds (cumulative count 997)
99.805% <= 6.639 milliseconds (cumulative count 999)
99.902% <= 6.919 milliseconds (cumulative count 1000)
100.000% <= 6.919 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.400% <= 1.007 milliseconds (cumulative count 4)
0.500% <= 1.103 milliseconds (cumulative count 5)
0.800% <= 1.207 milliseconds (cumulative count 8)
1.400% <= 1.303 milliseconds (cumulative count 14)
2.200% <= 1.407 milliseconds (cumulative count 22)
3.700% <= 1.503 milliseconds (cumulative count 37)
5.600% <= 1.607 milliseconds (cumulative count 56)
7.100% <= 1.703 milliseconds (cumulative count 71)
9.700% <= 1.807 milliseconds (cumulative count 97)
11.800% <= 1.903 milliseconds (cumulative count 118)
14.800% <= 2.007 milliseconds (cumulative count 148)
17.300% <= 2.103 milliseconds (cumulative count 173)
54.900% <= 3.103 milliseconds (cumulative count 549)
87.300% <= 4.103 milliseconds (cumulative count 873)
96.900% <= 5.103 milliseconds (cumulative count 969)
99.500% <= 6.103 milliseconds (cumulative count 995)
100.000% <= 7.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 12987.01 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        3.050     0.912     3.007     4.767     5.855     6.919
 
====== INCR ======
  1000 requests completed in 0.06 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.455 milliseconds (cumulative count 1)
50.000% <= 1.767 milliseconds (cumulative count 510)
75.000% <= 1.943 milliseconds (cumulative count 754)
87.500% <= 2.143 milliseconds (cumulative count 876)
93.750% <= 2.423 milliseconds (cumulative count 939)
96.875% <= 2.631 milliseconds (cumulative count 969)
98.438% <= 2.879 milliseconds (cumulative count 985)
99.219% <= 2.951 milliseconds (cumulative count 993)
99.609% <= 2.999 milliseconds (cumulative count 997)
99.805% <= 3.079 milliseconds (cumulative count 999)
99.902% <= 3.159 milliseconds (cumulative count 1000)
100.000% <= 3.159 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 0.503 milliseconds (cumulative count 1)
0.300% <= 0.607 milliseconds (cumulative count 3)
0.600% <= 0.807 milliseconds (cumulative count 6)
0.900% <= 0.903 milliseconds (cumulative count 9)
1.100% <= 1.007 milliseconds (cumulative count 11)
1.800% <= 1.103 milliseconds (cumulative count 18)
2.300% <= 1.207 milliseconds (cumulative count 23)
4.400% <= 1.303 milliseconds (cumulative count 44)
8.100% <= 1.407 milliseconds (cumulative count 81)
13.900% <= 1.503 milliseconds (cumulative count 139)
24.800% <= 1.607 milliseconds (cumulative count 248)
39.400% <= 1.703 milliseconds (cumulative count 394)
56.400% <= 1.807 milliseconds (cumulative count 564)
70.500% <= 1.903 milliseconds (cumulative count 705)
81.300% <= 2.007 milliseconds (cumulative count 813)
86.500% <= 2.103 milliseconds (cumulative count 865)
99.900% <= 3.103 milliseconds (cumulative count 999)
100.000% <= 4.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 16393.44 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.804     0.448     1.767     2.487     2.919     3.159
 
====== LPUSH ======
  1000 requests completed in 0.08 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.191 milliseconds (cumulative count 1)
50.000% <= 2.791 milliseconds (cumulative count 503)
75.000% <= 3.335 milliseconds (cumulative count 753)
87.500% <= 3.823 milliseconds (cumulative count 878)
93.750% <= 4.247 milliseconds (cumulative count 938)
96.875% <= 4.639 milliseconds (cumulative count 969)
98.438% <= 4.919 milliseconds (cumulative count 985)
99.219% <= 5.079 milliseconds (cumulative count 993)
99.609% <= 5.295 milliseconds (cumulative count 997)
99.805% <= 5.383 milliseconds (cumulative count 999)
99.902% <= 5.503 milliseconds (cumulative count 1000)
100.000% <= 5.503 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 1.207 milliseconds (cumulative count 1)
0.200% <= 1.303 milliseconds (cumulative count 2)
0.800% <= 1.407 milliseconds (cumulative count 8)
1.300% <= 1.503 milliseconds (cumulative count 13)
2.100% <= 1.607 milliseconds (cumulative count 21)
3.000% <= 1.703 milliseconds (cumulative count 30)
4.800% <= 1.807 milliseconds (cumulative count 48)
6.700% <= 1.903 milliseconds (cumulative count 67)
10.200% <= 2.007 milliseconds (cumulative count 102)
13.600% <= 2.103 milliseconds (cumulative count 136)
65.800% <= 3.103 milliseconds (cumulative count 658)
92.000% <= 4.103 milliseconds (cumulative count 920)
99.300% <= 5.103 milliseconds (cumulative count 993)
100.000% <= 6.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 13157.90 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.895     1.184     2.791     4.391     4.999     5.503
 
RPUSH: rps=710.3 (overall: 9421.1) avg_msec=3.648 (overall: 3.648)
                                                                   
====== RPUSH ======
  1000 requests completed in 0.08 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.327 milliseconds (cumulative count 1)
50.000% <= 3.159 milliseconds (cumulative count 501)
75.000% <= 3.871 milliseconds (cumulative count 750)
87.500% <= 4.407 milliseconds (cumulative count 876)
93.750% <= 4.743 milliseconds (cumulative count 941)
96.875% <= 5.079 milliseconds (cumulative count 969)
98.438% <= 5.399 milliseconds (cumulative count 985)
99.219% <= 5.615 milliseconds (cumulative count 993)
99.609% <= 5.711 milliseconds (cumulative count 997)
99.805% <= 6.079 milliseconds (cumulative count 999)
99.902% <= 6.215 milliseconds (cumulative count 1000)
100.000% <= 6.215 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.200% <= 1.407 milliseconds (cumulative count 2)
0.500% <= 1.607 milliseconds (cumulative count 5)
1.000% <= 1.703 milliseconds (cumulative count 10)
2.400% <= 1.807 milliseconds (cumulative count 24)
3.700% <= 1.903 milliseconds (cumulative count 37)
5.000% <= 2.007 milliseconds (cumulative count 50)
7.200% <= 2.103 milliseconds (cumulative count 72)
47.800% <= 3.103 milliseconds (cumulative count 478)
80.000% <= 4.103 milliseconds (cumulative count 800)
97.100% <= 5.103 milliseconds (cumulative count 971)
99.900% <= 6.103 milliseconds (cumulative count 999)
100.000% <= 7.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 12195.12 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        3.278     1.320     3.159     4.831     5.463     6.215
 
====== LPOP ======
  1000 requests completed in 0.09 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.223 milliseconds (cumulative count 1)
50.000% <= 3.911 milliseconds (cumulative count 507)
75.000% <= 4.431 milliseconds (cumulative count 750)
87.500% <= 4.903 milliseconds (cumulative count 877)
93.750% <= 5.375 milliseconds (cumulative count 938)
96.875% <= 6.095 milliseconds (cumulative count 969)
98.438% <= 7.031 milliseconds (cumulative count 985)
99.219% <= 7.431 milliseconds (cumulative count 993)
99.609% <= 7.775 milliseconds (cumulative count 997)
99.805% <= 8.935 milliseconds (cumulative count 999)
99.902% <= 9.767 milliseconds (cumulative count 1000)
100.000% <= 9.767 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 1.303 milliseconds (cumulative count 1)
0.200% <= 1.503 milliseconds (cumulative count 2)
0.500% <= 1.607 milliseconds (cumulative count 5)
0.900% <= 1.703 milliseconds (cumulative count 9)
1.500% <= 1.807 milliseconds (cumulative count 15)
2.100% <= 1.903 milliseconds (cumulative count 21)
3.000% <= 2.007 milliseconds (cumulative count 30)
3.700% <= 2.103 milliseconds (cumulative count 37)
20.500% <= 3.103 milliseconds (cumulative count 205)
62.700% <= 4.103 milliseconds (cumulative count 627)
90.400% <= 5.103 milliseconds (cumulative count 904)
97.000% <= 6.103 milliseconds (cumulative count 970)
98.500% <= 7.103 milliseconds (cumulative count 985)
99.800% <= 8.103 milliseconds (cumulative count 998)
99.900% <= 9.103 milliseconds (cumulative count 999)
100.000% <= 10.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 11363.64 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        3.885     1.216     3.911     5.551     7.279     9.767
 
RPOP: rps=3694.4 (overall: 11935.9) avg_msec=3.583 (overall: 3.583)
                                                                    
====== RPOP ======
  1000 requests completed in 0.08 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.119 milliseconds (cumulative count 1)
50.000% <= 3.711 milliseconds (cumulative count 504)
75.000% <= 4.159 milliseconds (cumulative count 751)
87.500% <= 4.591 milliseconds (cumulative count 877)
93.750% <= 5.007 milliseconds (cumulative count 938)
96.875% <= 5.231 milliseconds (cumulative count 973)
98.438% <= 5.423 milliseconds (cumulative count 985)
99.219% <= 5.679 milliseconds (cumulative count 993)
99.609% <= 6.079 milliseconds (cumulative count 997)
99.805% <= 6.375 milliseconds (cumulative count 999)
99.902% <= 6.591 milliseconds (cumulative count 1000)
100.000% <= 6.591 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.200% <= 1.207 milliseconds (cumulative count 2)
0.700% <= 1.303 milliseconds (cumulative count 7)
1.200% <= 1.407 milliseconds (cumulative count 12)
2.100% <= 1.503 milliseconds (cumulative count 21)
2.700% <= 1.607 milliseconds (cumulative count 27)
3.600% <= 1.703 milliseconds (cumulative count 36)
5.000% <= 1.807 milliseconds (cumulative count 50)
6.100% <= 1.903 milliseconds (cumulative count 61)
7.800% <= 2.007 milliseconds (cumulative count 78)
9.200% <= 2.103 milliseconds (cumulative count 92)
26.800% <= 3.103 milliseconds (cumulative count 268)
73.000% <= 4.103 milliseconds (cumulative count 730)
94.800% <= 5.103 milliseconds (cumulative count 948)
99.700% <= 6.103 milliseconds (cumulative count 997)
100.000% <= 7.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 12195.12 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        3.581     1.112     3.711     5.111     5.567     6.591
 
====== SADD ======
  1000 requests completed in 0.05 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 1)
50.000% <= 1.367 milliseconds (cumulative count 500)
75.000% <= 1.855 milliseconds (cumulative count 755)
87.500% <= 2.007 milliseconds (cumulative count 877)
93.750% <= 2.295 milliseconds (cumulative count 938)
96.875% <= 2.743 milliseconds (cumulative count 969)
98.438% <= 3.087 milliseconds (cumulative count 986)
99.219% <= 3.383 milliseconds (cumulative count 993)
99.609% <= 3.471 milliseconds (cumulative count 997)
99.805% <= 3.575 milliseconds (cumulative count 999)
99.902% <= 3.591 milliseconds (cumulative count 1000)
100.000% <= 3.591 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.500% <= 0.207 milliseconds (cumulative count 5)
1.000% <= 0.303 milliseconds (cumulative count 10)
1.500% <= 0.407 milliseconds (cumulative count 15)
2.100% <= 0.503 milliseconds (cumulative count 21)
19.400% <= 0.607 milliseconds (cumulative count 194)
35.600% <= 0.703 milliseconds (cumulative count 356)
36.000% <= 0.807 milliseconds (cumulative count 360)
36.300% <= 0.903 milliseconds (cumulative count 363)
36.800% <= 1.007 milliseconds (cumulative count 368)
38.700% <= 1.103 milliseconds (cumulative count 387)
45.500% <= 1.207 milliseconds (cumulative count 455)
49.400% <= 1.303 milliseconds (cumulative count 494)
50.400% <= 1.407 milliseconds (cumulative count 504)
52.000% <= 1.503 milliseconds (cumulative count 520)
56.200% <= 1.607 milliseconds (cumulative count 562)
63.900% <= 1.703 milliseconds (cumulative count 639)
72.400% <= 1.807 milliseconds (cumulative count 724)
80.000% <= 1.903 milliseconds (cumulative count 800)
87.700% <= 2.007 milliseconds (cumulative count 877)
90.700% <= 2.103 milliseconds (cumulative count 907)
98.600% <= 3.103 milliseconds (cumulative count 986)
100.000% <= 4.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 21276.60 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.345     0.144     1.367     2.439     3.239     3.591
 
====== HSET ======
  1000 requests completed in 0.03 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 1)
50.000% <= 0.719 milliseconds (cumulative count 535)
75.000% <= 0.751 milliseconds (cumulative count 775)
87.500% <= 0.783 milliseconds (cumulative count 885)
93.750% <= 0.831 milliseconds (cumulative count 941)
96.875% <= 0.999 milliseconds (cumulative count 969)
98.438% <= 1.127 milliseconds (cumulative count 985)
99.219% <= 1.271 milliseconds (cumulative count 994)
99.609% <= 1.343 milliseconds (cumulative count 997)
99.805% <= 1.407 milliseconds (cumulative count 999)
99.902% <= 1.527 milliseconds (cumulative count 1000)
100.000% <= 1.527 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.400% <= 0.207 milliseconds (cumulative count 4)
0.800% <= 0.303 milliseconds (cumulative count 8)
1.200% <= 0.407 milliseconds (cumulative count 12)
1.900% <= 0.503 milliseconds (cumulative count 19)
3.700% <= 0.607 milliseconds (cumulative count 37)
35.300% <= 0.703 milliseconds (cumulative count 353)
91.600% <= 0.807 milliseconds (cumulative count 916)
95.300% <= 0.903 milliseconds (cumulative count 953)
97.000% <= 1.007 milliseconds (cumulative count 970)
98.200% <= 1.103 milliseconds (cumulative count 982)
99.100% <= 1.207 milliseconds (cumulative count 991)
99.500% <= 1.303 milliseconds (cumulative count 995)
99.900% <= 1.407 milliseconds (cumulative count 999)
100.000% <= 1.607 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 37037.04 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.728     0.136     0.719     0.887     1.175     1.527
 
====== SPOP ======
  1000 requests completed in 0.02 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.207 milliseconds (cumulative count 1)
50.000% <= 0.543 milliseconds (cumulative count 524)
75.000% <= 0.599 milliseconds (cumulative count 768)
87.500% <= 0.663 milliseconds (cumulative count 883)
93.750% <= 0.783 milliseconds (cumulative count 938)
96.875% <= 0.879 milliseconds (cumulative count 969)
98.438% <= 0.959 milliseconds (cumulative count 985)
99.219% <= 1.015 milliseconds (cumulative count 993)
99.609% <= 1.079 milliseconds (cumulative count 997)
99.805% <= 1.127 milliseconds (cumulative count 999)
99.902% <= 1.167 milliseconds (cumulative count 1000)
100.000% <= 1.167 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 0.207 milliseconds (cumulative count 1)
1.200% <= 0.303 milliseconds (cumulative count 12)
2.500% <= 0.407 milliseconds (cumulative count 25)
25.500% <= 0.503 milliseconds (cumulative count 255)
79.300% <= 0.607 milliseconds (cumulative count 793)
91.100% <= 0.703 milliseconds (cumulative count 911)
94.700% <= 0.807 milliseconds (cumulative count 947)
97.300% <= 0.903 milliseconds (cumulative count 973)
99.200% <= 1.007 milliseconds (cumulative count 992)
99.800% <= 1.103 milliseconds (cumulative count 998)
100.000% <= 1.207 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 50000.00 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.565     0.200     0.543     0.815     0.991     1.167
 
====== ZADD ======
  1000 requests completed in 0.02 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.191 milliseconds (cumulative count 2)
50.000% <= 0.519 milliseconds (cumulative count 507)
75.000% <= 0.567 milliseconds (cumulative count 751)
87.500% <= 0.615 milliseconds (cumulative count 878)
93.750% <= 0.711 milliseconds (cumulative count 940)
96.875% <= 0.791 milliseconds (cumulative count 969)
98.438% <= 0.839 milliseconds (cumulative count 986)
99.219% <= 0.871 milliseconds (cumulative count 994)
99.609% <= 0.903 milliseconds (cumulative count 997)
99.805% <= 0.927 milliseconds (cumulative count 999)
99.902% <= 0.935 milliseconds (cumulative count 1000)
100.000% <= 0.935 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.200% <= 0.207 milliseconds (cumulative count 2)
0.700% <= 0.303 milliseconds (cumulative count 7)
2.300% <= 0.407 milliseconds (cumulative count 23)
39.200% <= 0.503 milliseconds (cumulative count 392)
86.900% <= 0.607 milliseconds (cumulative count 869)
93.700% <= 0.703 milliseconds (cumulative count 937)
97.200% <= 0.807 milliseconds (cumulative count 972)
99.700% <= 0.903 milliseconds (cumulative count 997)
100.000% <= 1.007 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 52631.58 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.535     0.184     0.519     0.743     0.855     0.935
 
====== ZPOPMIN ======
  1000 requests completed in 0.02 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 1)
50.000% <= 0.503 milliseconds (cumulative count 551)
75.000% <= 0.543 milliseconds (cumulative count 751)
87.500% <= 0.599 milliseconds (cumulative count 879)
93.750% <= 0.671 milliseconds (cumulative count 941)
96.875% <= 0.783 milliseconds (cumulative count 970)
98.438% <= 0.879 milliseconds (cumulative count 985)
99.219% <= 1.015 milliseconds (cumulative count 993)
99.609% <= 1.151 milliseconds (cumulative count 997)
99.805% <= 1.159 milliseconds (cumulative count 999)
99.902% <= 1.191 milliseconds (cumulative count 1000)
100.000% <= 1.191 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.500% <= 0.207 milliseconds (cumulative count 5)
1.000% <= 0.303 milliseconds (cumulative count 10)
5.700% <= 0.407 milliseconds (cumulative count 57)
55.100% <= 0.503 milliseconds (cumulative count 551)
88.300% <= 0.607 milliseconds (cumulative count 883)
95.400% <= 0.703 milliseconds (cumulative count 954)
97.500% <= 0.807 milliseconds (cumulative count 975)
98.800% <= 0.903 milliseconds (cumulative count 988)
99.200% <= 1.007 milliseconds (cumulative count 992)
99.400% <= 1.103 milliseconds (cumulative count 994)
100.000% <= 1.207 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 52631.58 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.516     0.144     0.503     0.695     0.927     1.191
 
====== LPUSH (needed to benchmark LRANGE) ======
  1000 requests completed in 0.02 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.447 milliseconds (cumulative count 2)
50.000% <= 0.871 milliseconds (cumulative count 501)
75.000% <= 1.039 milliseconds (cumulative count 751)
87.500% <= 1.175 milliseconds (cumulative count 877)
93.750% <= 1.287 milliseconds (cumulative count 938)
96.875% <= 1.375 milliseconds (cumulative count 971)
98.438% <= 1.447 milliseconds (cumulative count 986)
99.219% <= 1.519 milliseconds (cumulative count 996)
99.609% <= 1.567 milliseconds (cumulative count 997)
99.805% <= 1.631 milliseconds (cumulative count 999)
99.902% <= 1.655 milliseconds (cumulative count 1000)
100.000% <= 1.655 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
1.700% <= 0.503 milliseconds (cumulative count 17)
9.000% <= 0.607 milliseconds (cumulative count 90)
21.400% <= 0.703 milliseconds (cumulative count 214)
40.400% <= 0.807 milliseconds (cumulative count 404)
55.800% <= 0.903 milliseconds (cumulative count 558)
71.100% <= 1.007 milliseconds (cumulative count 711)
82.300% <= 1.103 milliseconds (cumulative count 823)
90.000% <= 1.207 milliseconds (cumulative count 900)
94.800% <= 1.303 milliseconds (cumulative count 948)
97.900% <= 1.407 milliseconds (cumulative count 979)
99.100% <= 1.503 milliseconds (cumulative count 991)
99.800% <= 1.607 milliseconds (cumulative count 998)
100.000% <= 1.703 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 43478.26 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.893     0.440     0.871     1.311     1.479     1.655
 
LRANGE_100 (first 100 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_100 (first 100 elements): rps=450.2 (overall: 368.1) avg_msec=26.491 (overall: 26.491)
                                                                                              
LRANGE_100 (first 100 elements): rps=788.8 (overall: 557.3) avg_msec=23.932 (overall: 24.862)
                                                                                              
LRANGE_100 (first 100 elements): rps=653.4 (overall: 587.1) avg_msec=17.726 (overall: 22.398)
                                                                                              
LRANGE_100 (first 100 elements): rps=613.5 (overall: 593.4) avg_msec=21.971 (overall: 22.294)
                                                                                              
LRANGE_100 (first 100 elements): rps=717.1 (overall: 617.1) avg_msec=15.043 (overall: 20.681)
                                                                                              
====== LRANGE_100 (first 100 elements) ======
  1000 requests completed in 1.51 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.007 milliseconds (cumulative count 1)
50.000% <= 24.047 milliseconds (cumulative count 501)
75.000% <= 28.319 milliseconds (cumulative count 750)
87.500% <= 31.247 milliseconds (cumulative count 876)
93.750% <= 35.359 milliseconds (cumulative count 938)
96.875% <= 42.143 milliseconds (cumulative count 969)
98.438% <= 44.575 milliseconds (cumulative count 985)
99.219% <= 46.271 milliseconds (cumulative count 993)
99.609% <= 46.527 milliseconds (cumulative count 997)
99.805% <= 46.719 milliseconds (cumulative count 999)
99.902% <= 46.815 milliseconds (cumulative count 1000)
100.000% <= 46.815 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 1.007 milliseconds (cumulative count 1)
0.200% <= 1.303 milliseconds (cumulative count 2)
0.400% <= 1.407 milliseconds (cumulative count 4)
0.900% <= 1.503 milliseconds (cumulative count 9)
1.400% <= 1.607 milliseconds (cumulative count 14)
1.600% <= 1.703 milliseconds (cumulative count 16)
2.000% <= 1.807 milliseconds (cumulative count 20)
2.600% <= 1.903 milliseconds (cumulative count 26)
2.700% <= 2.007 milliseconds (cumulative count 27)
6.300% <= 3.103 milliseconds (cumulative count 63)
8.600% <= 4.103 milliseconds (cumulative count 86)
12.200% <= 5.103 milliseconds (cumulative count 122)
16.900% <= 6.103 milliseconds (cumulative count 169)
22.300% <= 7.103 milliseconds (cumulative count 223)
26.500% <= 8.103 milliseconds (cumulative count 265)
28.000% <= 9.103 milliseconds (cumulative count 280)
29.600% <= 10.103 milliseconds (cumulative count 296)
29.900% <= 11.103 milliseconds (cumulative count 299)
30.100% <= 12.103 milliseconds (cumulative count 301)
31.700% <= 13.103 milliseconds (cumulative count 317)
32.000% <= 14.103 milliseconds (cumulative count 320)
33.000% <= 15.103 milliseconds (cumulative count 330)
35.000% <= 16.103 milliseconds (cumulative count 350)
37.100% <= 17.103 milliseconds (cumulative count 371)
40.200% <= 18.111 milliseconds (cumulative count 402)
42.500% <= 19.103 milliseconds (cumulative count 425)
45.100% <= 20.111 milliseconds (cumulative count 451)
45.400% <= 21.103 milliseconds (cumulative count 454)
46.100% <= 22.111 milliseconds (cumulative count 461)
48.000% <= 23.103 milliseconds (cumulative count 480)
50.800% <= 24.111 milliseconds (cumulative count 508)
57.400% <= 25.103 milliseconds (cumulative count 574)
64.000% <= 26.111 milliseconds (cumulative count 640)
68.000% <= 27.103 milliseconds (cumulative count 680)
74.100% <= 28.111 milliseconds (cumulative count 741)
79.400% <= 29.103 milliseconds (cumulative count 794)
83.900% <= 30.111 milliseconds (cumulative count 839)
86.900% <= 31.103 milliseconds (cumulative count 869)
89.800% <= 32.111 milliseconds (cumulative count 898)
91.000% <= 33.119 milliseconds (cumulative count 910)
92.300% <= 34.111 milliseconds (cumulative count 923)
93.400% <= 35.103 milliseconds (cumulative count 934)
94.200% <= 36.127 milliseconds (cumulative count 942)
95.000% <= 37.119 milliseconds (cumulative count 950)
95.300% <= 38.111 milliseconds (cumulative count 953)
95.900% <= 39.103 milliseconds (cumulative count 959)
96.000% <= 41.119 milliseconds (cumulative count 960)
96.800% <= 42.111 milliseconds (cumulative count 968)
97.700% <= 43.103 milliseconds (cumulative count 977)
98.000% <= 44.127 milliseconds (cumulative count 980)
98.900% <= 45.119 milliseconds (cumulative count 989)
99.000% <= 46.111 milliseconds (cumulative count 990)
100.000% <= 47.103 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 660.94 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
       20.186     1.000    24.047    37.023    46.111    46.815
 
LRANGE_300 (first 300 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_300 (first 300 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_300 (first 300 elements): rps=294.8 (overall: 135.8) avg_msec=71.168 (overall: 71.168)
                                                                                              
LRANGE_300 (first 300 elements): rps=259.0 (overall: 174.6) avg_msec=73.324 (overall: 72.176)
                                                                                              
LRANGE_300 (first 300 elements): rps=59.5 (overall: 146.9) avg_msec=68.301 (overall: 71.799)
                                                                                             
LRANGE_300 (first 300 elements): rps=199.2 (overall: 157.0) avg_msec=72.904 (overall: 72.070)
                                                                                              
LRANGE_300 (first 300 elements): rps=204.0 (overall: 164.6) avg_msec=68.941 (overall: 71.444)
                                                                                              
LRANGE_300 (first 300 elements): rps=203.2 (overall: 170.0) avg_msec=65.108 (overall: 70.388)
                                                                                              
LRANGE_300 (first 300 elements): rps=199.2 (overall: 173.6) avg_msec=64.044 (overall: 69.497)
                                                                                              
LRANGE_300 (first 300 elements): rps=200.0 (overall: 176.4) avg_msec=63.350 (overall: 68.740)
                                                                                              
LRANGE_300 (first 300 elements): rps=202.4 (overall: 179.0) avg_msec=71.307 (overall: 69.026)
                                                                                              
LRANGE_300 (first 300 elements): rps=199.2 (overall: 180.8) avg_msec=81.658 (overall: 70.272)
                                                                                              
LRANGE_300 (first 300 elements): rps=207.2 (overall: 183.0) avg_msec=72.411 (overall: 70.471)
                                                                                              
LRANGE_300 (first 300 elements): rps=198.4 (overall: 184.2) avg_msec=80.086 (overall: 71.261)
                                                                                              
LRANGE_300 (first 300 elements): rps=212.0 (overall: 186.1) avg_msec=61.110 (overall: 70.448)
                                                                                              
LRANGE_300 (first 300 elements): rps=199.2 (overall: 187.0) avg_msec=63.296 (overall: 69.946)
                                                                                              
LRANGE_300 (first 300 elements): rps=215.1 (overall: 188.7) avg_msec=53.870 (overall: 68.812)
                                                                                              
LRANGE_300 (first 300 elements): rps=199.2 (overall: 189.3) avg_msec=64.197 (overall: 68.530)
                                                                                              
LRANGE_300 (first 300 elements): rps=197.6 (overall: 189.8) avg_msec=57.174 (overall: 67.874)
                                                                                              
LRANGE_300 (first 300 elements): rps=216.0 (overall: 191.1) avg_msec=52.294 (overall: 66.960)
                                                                                              
LRANGE_300 (first 300 elements): rps=8.0 (overall: 182.1) avg_msec=3.644 (overall: 66.822)
                                                                                           
====== LRANGE_300 (first 300 elements) ======
  1000 requests completed in 5.23 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.279 milliseconds (cumulative count 1)
50.000% <= 80.447 milliseconds (cumulative count 500)
75.000% <= 90.431 milliseconds (cumulative count 752)
87.500% <= 95.999 milliseconds (cumulative count 875)
93.750% <= 100.223 milliseconds (cumulative count 938)
96.875% <= 103.551 milliseconds (cumulative count 969)
98.438% <= 107.007 milliseconds (cumulative count 985)
99.219% <= 112.831 milliseconds (cumulative count 993)
99.609% <= 113.471 milliseconds (cumulative count 997)
99.805% <= 113.535 milliseconds (cumulative count 999)
99.902% <= 113.727 milliseconds (cumulative count 1000)
100.000% <= 113.727 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 1.303 milliseconds (cumulative count 1)
0.400% <= 2.007 milliseconds (cumulative count 4)
1.800% <= 3.103 milliseconds (cumulative count 18)
3.800% <= 4.103 milliseconds (cumulative count 38)
4.600% <= 5.103 milliseconds (cumulative count 46)
6.200% <= 6.103 milliseconds (cumulative count 62)
6.800% <= 7.103 milliseconds (cumulative count 68)
7.300% <= 8.103 milliseconds (cumulative count 73)
7.600% <= 9.103 milliseconds (cumulative count 76)
8.300% <= 10.103 milliseconds (cumulative count 83)
8.900% <= 12.103 milliseconds (cumulative count 89)
9.200% <= 14.103 milliseconds (cumulative count 92)
9.900% <= 15.103 milliseconds (cumulative count 99)
10.700% <= 16.103 milliseconds (cumulative count 107)
11.300% <= 17.103 milliseconds (cumulative count 113)
11.400% <= 18.111 milliseconds (cumulative count 114)
12.300% <= 19.103 milliseconds (cumulative count 123)
13.500% <= 20.111 milliseconds (cumulative count 135)
13.900% <= 21.103 milliseconds (cumulative count 139)
14.800% <= 22.111 milliseconds (cumulative count 148)
16.600% <= 23.103 milliseconds (cumulative count 166)
16.700% <= 24.111 milliseconds (cumulative count 167)
17.400% <= 26.111 milliseconds (cumulative count 174)
17.700% <= 27.103 milliseconds (cumulative count 177)
18.800% <= 31.103 milliseconds (cumulative count 188)
19.200% <= 32.111 milliseconds (cumulative count 192)
19.500% <= 38.111 milliseconds (cumulative count 195)
20.300% <= 43.103 milliseconds (cumulative count 203)
21.600% <= 44.127 milliseconds (cumulative count 216)
22.900% <= 45.119 milliseconds (cumulative count 229)
24.700% <= 46.111 milliseconds (cumulative count 247)
26.700% <= 47.103 milliseconds (cumulative count 267)
27.700% <= 48.127 milliseconds (cumulative count 277)
30.000% <= 49.119 milliseconds (cumulative count 300)
31.200% <= 50.111 milliseconds (cumulative count 312)
32.900% <= 51.103 milliseconds (cumulative count 329)
33.300% <= 52.127 milliseconds (cumulative count 333)
34.400% <= 54.111 milliseconds (cumulative count 344)
35.100% <= 55.103 milliseconds (cumulative count 351)
35.200% <= 60.127 milliseconds (cumulative count 352)
36.500% <= 61.119 milliseconds (cumulative count 365)
37.700% <= 62.111 milliseconds (cumulative count 377)
38.000% <= 63.103 milliseconds (cumulative count 380)
39.200% <= 64.127 milliseconds (cumulative count 392)
40.500% <= 65.119 milliseconds (cumulative count 405)
41.300% <= 66.111 milliseconds (cumulative count 413)
42.600% <= 67.135 milliseconds (cumulative count 426)
43.200% <= 69.119 milliseconds (cumulative count 432)
43.700% <= 71.103 milliseconds (cumulative count 437)
44.800% <= 72.127 milliseconds (cumulative count 448)
46.000% <= 73.151 milliseconds (cumulative count 460)
46.900% <= 76.159 milliseconds (cumulative count 469)
48.000% <= 77.119 milliseconds (cumulative count 480)
48.900% <= 79.103 milliseconds (cumulative count 489)
49.500% <= 80.127 milliseconds (cumulative count 495)
52.400% <= 81.151 milliseconds (cumulative count 524)
54.900% <= 82.111 milliseconds (cumulative count 549)
55.200% <= 83.135 milliseconds (cumulative count 552)
56.000% <= 84.159 milliseconds (cumulative count 560)
59.000% <= 85.119 milliseconds (cumulative count 590)
61.000% <= 86.143 milliseconds (cumulative count 610)
65.400% <= 87.103 milliseconds (cumulative count 654)
69.900% <= 88.127 milliseconds (cumulative count 699)
72.500% <= 89.151 milliseconds (cumulative count 725)
73.900% <= 90.111 milliseconds (cumulative count 739)
77.400% <= 91.135 milliseconds (cumulative count 774)
80.200% <= 92.159 milliseconds (cumulative count 802)
81.400% <= 93.119 milliseconds (cumulative count 814)
82.600% <= 94.143 milliseconds (cumulative count 826)
85.900% <= 95.103 milliseconds (cumulative count 859)
87.600% <= 96.127 milliseconds (cumulative count 876)
89.100% <= 97.151 milliseconds (cumulative count 891)
90.100% <= 98.111 milliseconds (cumulative count 901)
92.900% <= 99.135 milliseconds (cumulative count 929)
93.600% <= 100.159 milliseconds (cumulative count 936)
94.900% <= 101.119 milliseconds (cumulative count 949)
96.400% <= 102.143 milliseconds (cumulative count 964)
97.500% <= 104.127 milliseconds (cumulative count 975)
98.200% <= 105.151 milliseconds (cumulative count 982)
98.600% <= 107.135 milliseconds (cumulative count 986)
99.000% <= 110.143 milliseconds (cumulative count 990)
99.400% <= 113.151 milliseconds (cumulative count 994)
100.000% <= 114.111 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 191.17 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
       66.361     1.272    80.447   101.183   109.503   113.727
 
LRANGE_500 (first 500 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_500 (first 500 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_500 (first 500 elements): rps=199.2 (overall: 86.2) avg_msec=168.221 (overall: 168.221)
                                                                                               
LRANGE_500 (first 500 elements): rps=4.0 (overall: 61.4) avg_msec=6.836 (overall: 165.057)
                                                                                           
LRANGE_500 (first 500 elements): rps=200.0 (overall: 93.4) avg_msec=138.292 (overall: 151.807)
                                                                                               
LRANGE_500 (first 500 elements): rps=198.4 (overall: 113.3) avg_msec=116.590 (overall: 140.146)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 95.4) avg_msec=-nan (overall: 140.146)
                                                                                          
LRANGE_500 (first 500 elements): rps=203.2 (overall: 110.1) avg_msec=110.473 (overall: 132.654)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 96.9) avg_msec=-nan (overall: 132.654)
                                                                                          
LRANGE_500 (first 500 elements): rps=199.2 (overall: 107.9) avg_msec=123.977 (overall: 130.932)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 97.4) avg_msec=-nan (overall: 130.932)
                                                                                          
LRANGE_500 (first 500 elements): rps=203.2 (overall: 106.8) avg_msec=122.850 (overall: 129.572)
                                                                                                
LRANGE_500 (first 500 elements): rps=199.2 (overall: 114.5) avg_msec=132.753 (overall: 130.030)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 105.9) avg_msec=-nan (overall: 130.030)
                                                                                           
LRANGE_500 (first 500 elements): rps=199.2 (overall: 112.4) avg_msec=124.285 (overall: 129.319)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 105.0) avg_msec=-nan (overall: 129.319)
                                                                                           
LRANGE_500 (first 500 elements): rps=199.2 (overall: 110.8) avg_msec=138.962 (overall: 130.381)
                                                                                                
LRANGE_500 (first 500 elements): rps=4.0 (overall: 104.7) avg_msec=2.548 (overall: 130.100)
                                                                                            
LRANGE_500 (first 500 elements): rps=202.4 (overall: 110.0) avg_msec=147.239 (overall: 131.828)
                                                                                                
LRANGE_500 (first 500 elements): rps=199.2 (overall: 114.6) avg_msec=99.862 (overall: 128.953)
                                                                                               
LRANGE_500 (first 500 elements): rps=0.0 (overall: 109.0) avg_msec=-nan (overall: 128.953)
                                                                                           
LRANGE_500 (first 500 elements): rps=203.2 (overall: 113.4) avg_msec=112.411 (overall: 127.563)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 108.3) avg_msec=-nan (overall: 127.563)
                                                                                           
LRANGE_500 (first 500 elements): rps=199.2 (overall: 112.2) avg_msec=113.276 (overall: 126.476)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 107.6) avg_msec=-nan (overall: 126.476)
                                                                                           
LRANGE_500 (first 500 elements): rps=203.2 (overall: 111.4) avg_msec=143.428 (overall: 127.697)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 107.2) avg_msec=-nan (overall: 127.697)
                                                                                           
LRANGE_500 (first 500 elements): rps=207.2 (overall: 110.8) avg_msec=115.888 (overall: 126.889)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 106.9) avg_msec=-nan (overall: 126.889)
                                                                                           
LRANGE_500 (first 500 elements): rps=200.0 (overall: 110.1) avg_msec=145.924 (overall: 128.064)
                                                                                                
LRANGE_500 (first 500 elements): rps=191.6 (overall: 112.9) avg_msec=128.729 (overall: 128.103)
                                                                                                
LRANGE_500 (first 500 elements): rps=7.9 (overall: 109.5) avg_msec=5.168 (overall: 127.817)
                                                                                            
LRANGE_500 (first 500 elements): rps=199.2 (overall: 112.3) avg_msec=120.226 (overall: 127.401)
                                                                                                
LRANGE_500 (first 500 elements): rps=0.0 (overall: 108.9) avg_msec=-nan (overall: 127.401)
                                                                                           
LRANGE_500 (first 500 elements): rps=204.0 (overall: 111.7) avg_msec=133.918 (overall: 127.746)
                                                                                                
====== LRANGE_500 (first 500 elements) ======
  1000 requests completed in 8.82 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.535 milliseconds (cumulative count 1)
50.000% <= 148.607 milliseconds (cumulative count 501)
75.000% <= 161.919 milliseconds (cumulative count 753)
87.500% <= 171.647 milliseconds (cumulative count 876)
93.750% <= 182.655 milliseconds (cumulative count 938)
96.875% <= 190.335 milliseconds (cumulative count 970)
98.438% <= 199.935 milliseconds (cumulative count 985)
99.219% <= 200.575 milliseconds (cumulative count 993)
99.609% <= 202.495 milliseconds (cumulative count 997)
99.805% <= 202.623 milliseconds (cumulative count 999)
99.902% <= 202.751 milliseconds (cumulative count 1000)
100.000% <= 202.751 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 1.607 milliseconds (cumulative count 1)
0.800% <= 3.103 milliseconds (cumulative count 8)
1.400% <= 4.103 milliseconds (cumulative count 14)
2.000% <= 5.103 milliseconds (cumulative count 20)
2.500% <= 6.103 milliseconds (cumulative count 25)
3.100% <= 7.103 milliseconds (cumulative count 31)
3.500% <= 8.103 milliseconds (cumulative count 35)
4.300% <= 9.103 milliseconds (cumulative count 43)
4.500% <= 10.103 milliseconds (cumulative count 45)
4.600% <= 11.103 milliseconds (cumulative count 46)
4.700% <= 12.103 milliseconds (cumulative count 47)
4.800% <= 13.103 milliseconds (cumulative count 48)
5.000% <= 14.103 milliseconds (cumulative count 50)
5.500% <= 16.103 milliseconds (cumulative count 55)
5.700% <= 20.111 milliseconds (cumulative count 57)
6.200% <= 23.103 milliseconds (cumulative count 62)
6.400% <= 24.111 milliseconds (cumulative count 64)
6.800% <= 25.103 milliseconds (cumulative count 68)
8.100% <= 26.111 milliseconds (cumulative count 81)
8.600% <= 27.103 milliseconds (cumulative count 86)
9.200% <= 28.111 milliseconds (cumulative count 92)
10.000% <= 29.103 milliseconds (cumulative count 100)
10.800% <= 37.119 milliseconds (cumulative count 108)
11.000% <= 38.111 milliseconds (cumulative count 110)
11.600% <= 44.127 milliseconds (cumulative count 116)
13.000% <= 45.119 milliseconds (cumulative count 130)
13.400% <= 47.103 milliseconds (cumulative count 134)
14.100% <= 53.119 milliseconds (cumulative count 141)
15.200% <= 54.111 milliseconds (cumulative count 152)
16.800% <= 55.103 milliseconds (cumulative count 168)
17.900% <= 56.127 milliseconds (cumulative count 179)
18.400% <= 57.119 milliseconds (cumulative count 184)
19.200% <= 58.111 milliseconds (cumulative count 192)
20.100% <= 59.103 milliseconds (cumulative count 201)
21.000% <= 60.127 milliseconds (cumulative count 210)
22.100% <= 61.119 milliseconds (cumulative count 221)
22.500% <= 62.111 milliseconds (cumulative count 225)
23.300% <= 83.135 milliseconds (cumulative count 233)
23.700% <= 84.159 milliseconds (cumulative count 237)
24.300% <= 85.119 milliseconds (cumulative count 243)
24.500% <= 87.103 milliseconds (cumulative count 245)
25.900% <= 88.127 milliseconds (cumulative count 259)
26.300% <= 90.111 milliseconds (cumulative count 263)
27.100% <= 91.135 milliseconds (cumulative count 271)
27.500% <= 117.119 milliseconds (cumulative count 275)
28.000% <= 118.143 milliseconds (cumulative count 280)
28.300% <= 119.103 milliseconds (cumulative count 283)
29.300% <= 120.127 milliseconds (cumulative count 293)
29.500% <= 123.135 milliseconds (cumulative count 295)
30.200% <= 124.159 milliseconds (cumulative count 302)
30.300% <= 132.223 milliseconds (cumulative count 303)
30.400% <= 134.143 milliseconds (cumulative count 304)
30.500% <= 135.167 milliseconds (cumulative count 305)
30.600% <= 136.191 milliseconds (cumulative count 306)
31.200% <= 137.215 milliseconds (cumulative count 312)
32.100% <= 138.111 milliseconds (cumulative count 321)
32.900% <= 139.135 milliseconds (cumulative count 329)
34.600% <= 140.159 milliseconds (cumulative count 346)
36.900% <= 141.183 milliseconds (cumulative count 369)
38.300% <= 142.207 milliseconds (cumulative count 383)
40.200% <= 143.103 milliseconds (cumulative count 402)
44.000% <= 144.127 milliseconds (cumulative count 440)
46.100% <= 146.175 milliseconds (cumulative count 461)
48.400% <= 147.199 milliseconds (cumulative count 484)
49.400% <= 148.223 milliseconds (cumulative count 494)
51.100% <= 149.119 milliseconds (cumulative count 511)
52.300% <= 150.143 milliseconds (cumulative count 523)
55.700% <= 151.167 milliseconds (cumulative count 557)
57.900% <= 152.191 milliseconds (cumulative count 579)
58.000% <= 153.215 milliseconds (cumulative count 580)
58.800% <= 154.111 milliseconds (cumulative count 588)
59.800% <= 155.135 milliseconds (cumulative count 598)
61.300% <= 156.159 milliseconds (cumulative count 613)
62.300% <= 157.183 milliseconds (cumulative count 623)
65.900% <= 158.207 milliseconds (cumulative count 659)
70.300% <= 159.103 milliseconds (cumulative count 703)
72.000% <= 160.127 milliseconds (cumulative count 720)
73.700% <= 161.151 milliseconds (cumulative count 737)
75.700% <= 162.175 milliseconds (cumulative count 757)
76.800% <= 163.199 milliseconds (cumulative count 768)
80.200% <= 164.223 milliseconds (cumulative count 802)
81.700% <= 165.119 milliseconds (cumulative count 817)
82.200% <= 166.143 milliseconds (cumulative count 822)
83.000% <= 167.167 milliseconds (cumulative count 830)
83.500% <= 168.191 milliseconds (cumulative count 835)
84.500% <= 169.215 milliseconds (cumulative count 845)
85.800% <= 170.111 milliseconds (cumulative count 858)
87.100% <= 171.135 milliseconds (cumulative count 871)
88.300% <= 172.159 milliseconds (cumulative count 883)
88.800% <= 173.183 milliseconds (cumulative count 888)
89.200% <= 174.207 milliseconds (cumulative count 892)
90.500% <= 175.103 milliseconds (cumulative count 905)
90.800% <= 176.127 milliseconds (cumulative count 908)
90.900% <= 178.175 milliseconds (cumulative count 909)
91.200% <= 179.199 milliseconds (cumulative count 912)
92.100% <= 180.223 milliseconds (cumulative count 921)
93.200% <= 182.143 milliseconds (cumulative count 932)
94.500% <= 183.167 milliseconds (cumulative count 945)
94.800% <= 184.191 milliseconds (cumulative count 948)
94.900% <= 187.135 milliseconds (cumulative count 949)
95.900% <= 188.159 milliseconds (cumulative count 959)
96.200% <= 189.183 milliseconds (cumulative count 962)
96.700% <= 190.207 milliseconds (cumulative count 967)
97.000% <= 191.103 milliseconds (cumulative count 970)
97.600% <= 194.175 milliseconds (cumulative count 976)
98.100% <= 196.223 milliseconds (cumulative count 981)
98.800% <= 200.191 milliseconds (cumulative count 988)
99.300% <= 201.215 milliseconds (cumulative count 993)
100.000% <= 203.135 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 113.35 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
      126.901     1.528   148.607   187.263   200.319   202.751
 
LRANGE_600 (first 600 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_600 (first 600 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_600 (first 600 elements): rps=0.0 (overall: 0.0) avg_msec=-nan (overall: -nan)
                                                                                      
LRANGE_600 (first 600 elements): rps=199.2 (overall: 62.5) avg_msec=188.360 (overall: 188.360)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 47.6) avg_msec=-nan (overall: 188.360)
                                                                                          
LRANGE_600 (first 600 elements): rps=199.2 (overall: 76.8) avg_msec=164.735 (overall: 176.547)
                                                                                               
LRANGE_600 (first 600 elements): rps=4.0 (overall: 65.0) avg_msec=6.812 (overall: 174.867)
                                                                                           
LRANGE_600 (first 600 elements): rps=199.2 (overall: 83.7) avg_msec=151.545 (overall: 167.144)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 73.5) avg_msec=-nan (overall: 167.144)
                                                                                          
LRANGE_600 (first 600 elements): rps=198.4 (overall: 87.1) avg_msec=161.153 (overall: 165.654)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 78.6) avg_msec=-nan (overall: 165.654)
                                                                                          
LRANGE_600 (first 600 elements): rps=202.4 (overall: 89.7) avg_msec=145.894 (overall: 161.655)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 82.4) avg_msec=-nan (overall: 161.655)
                                                                                          
LRANGE_600 (first 600 elements): rps=172.5 (overall: 89.3) avg_msec=140.455 (overall: 158.504)
                                                                                               
LRANGE_600 (first 600 elements): rps=32.0 (overall: 85.3) avg_msec=139.911 (overall: 158.014)
                                                                                              
LRANGE_600 (first 600 elements): rps=0.0 (overall: 79.7) avg_msec=-nan (overall: 158.014)
                                                                                          
LRANGE_600 (first 600 elements): rps=200.0 (overall: 87.1) avg_msec=150.916 (overall: 157.012)
                                                                                               
LRANGE_600 (first 600 elements): rps=4.0 (overall: 82.2) avg_msec=5.564 (overall: 156.585)
                                                                                           
LRANGE_600 (first 600 elements): rps=199.2 (overall: 88.7) avg_msec=156.090 (overall: 156.524)
                                                                                               
LRANGE_600 (first 600 elements): rps=4.0 (overall: 84.2) avg_msec=5.236 (overall: 156.151)
                                                                                           
LRANGE_600 (first 600 elements): rps=199.2 (overall: 89.9) avg_msec=152.957 (overall: 155.801)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 85.7) avg_msec=-nan (overall: 155.801)
                                                                                          
LRANGE_600 (first 600 elements): rps=199.2 (overall: 90.8) avg_msec=161.000 (overall: 156.315)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 86.9) avg_msec=-nan (overall: 156.315)
                                                                                          
LRANGE_600 (first 600 elements): rps=200.0 (overall: 91.6) avg_msec=152.769 (overall: 155.996)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 87.9) avg_msec=-nan (overall: 155.996)
                                                                                          
LRANGE_600 (first 600 elements): rps=0.0 (overall: 84.6) avg_msec=-nan (overall: 155.996)
                                                                                          
LRANGE_600 (first 600 elements): rps=199.2 (overall: 88.8) avg_msec=161.637 (overall: 156.461)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 85.6) avg_msec=-nan (overall: 156.461)
                                                                                          
LRANGE_600 (first 600 elements): rps=200.0 (overall: 89.5) avg_msec=151.697 (overall: 156.098)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 86.6) avg_msec=-nan (overall: 156.098)
                                                                                          
LRANGE_600 (first 600 elements): rps=199.2 (overall: 90.2) avg_msec=155.420 (overall: 156.050)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 87.4) avg_msec=-nan (overall: 156.050)
                                                                                          
LRANGE_600 (first 600 elements): rps=203.2 (overall: 90.9) avg_msec=124.348 (overall: 153.914)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 88.2) avg_msec=-nan (overall: 153.914)
                                                                                          
LRANGE_600 (first 600 elements): rps=203.2 (overall: 91.5) avg_msec=125.461 (overall: 152.118)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 88.9) avg_msec=-nan (overall: 152.118)
                                                                                          
LRANGE_600 (first 600 elements): rps=202.4 (overall: 92.0) avg_msec=139.103 (overall: 151.346)
                                                                                               
LRANGE_600 (first 600 elements): rps=0.0 (overall: 89.6) avg_msec=-nan (overall: 151.346)
                                                                                          
LRANGE_600 (first 600 elements): rps=199.2 (overall: 92.4) avg_msec=141.314 (overall: 150.794)
                                                                                               
LRANGE_600 (first 600 elements): rps=4.0 (overall: 90.2) avg_msec=3.500 (overall: 150.632)
                                                                                           
LRANGE_600 (first 600 elements): rps=200.0 (overall: 92.9) avg_msec=144.471 (overall: 150.311)
                                                                                               
LRANGE_600 (first 600 elements): rps=4.0 (overall: 90.7) avg_msec=3.444 (overall: 150.158)
                                                                                           
====== LRANGE_600 (first 600 elements) ======
  1000 requests completed in 10.64 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 3.063 milliseconds (cumulative count 1)
50.000% <= 176.255 milliseconds (cumulative count 500)
75.000% <= 194.815 milliseconds (cumulative count 752)
87.500% <= 205.311 milliseconds (cumulative count 875)
93.750% <= 210.815 milliseconds (cumulative count 939)
96.875% <= 215.935 milliseconds (cumulative count 977)
98.438% <= 216.959 milliseconds (cumulative count 986)
99.219% <= 217.727 milliseconds (cumulative count 994)
99.609% <= 217.983 milliseconds (cumulative count 997)
99.805% <= 218.239 milliseconds (cumulative count 999)
99.902% <= 218.367 milliseconds (cumulative count 1000)
100.000% <= 218.367 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.100% <= 3.103 milliseconds (cumulative count 1)
0.700% <= 4.103 milliseconds (cumulative count 7)
1.200% <= 5.103 milliseconds (cumulative count 12)
1.800% <= 6.103 milliseconds (cumulative count 18)
2.500% <= 7.103 milliseconds (cumulative count 25)
2.900% <= 8.103 milliseconds (cumulative count 29)
3.100% <= 9.103 milliseconds (cumulative count 31)
3.300% <= 10.103 milliseconds (cumulative count 33)
3.900% <= 11.103 milliseconds (cumulative count 39)
4.000% <= 12.103 milliseconds (cumulative count 40)
4.200% <= 14.103 milliseconds (cumulative count 42)
4.400% <= 15.103 milliseconds (cumulative count 44)
4.800% <= 16.103 milliseconds (cumulative count 48)
4.900% <= 24.111 milliseconds (cumulative count 49)
5.500% <= 27.103 milliseconds (cumulative count 55)
5.900% <= 29.103 milliseconds (cumulative count 59)
6.800% <= 30.111 milliseconds (cumulative count 68)
6.900% <= 31.103 milliseconds (cumulative count 69)
7.300% <= 32.111 milliseconds (cumulative count 73)
7.500% <= 33.119 milliseconds (cumulative count 75)
7.700% <= 34.111 milliseconds (cumulative count 77)
8.700% <= 49.119 milliseconds (cumulative count 87)
8.800% <= 50.111 milliseconds (cumulative count 88)
9.100% <= 52.127 milliseconds (cumulative count 91)
9.800% <= 53.119 milliseconds (cumulative count 98)
12.200% <= 54.111 milliseconds (cumulative count 122)
13.300% <= 55.103 milliseconds (cumulative count 133)
13.900% <= 56.127 milliseconds (cumulative count 139)
14.700% <= 57.119 milliseconds (cumulative count 147)
15.800% <= 58.111 milliseconds (cumulative count 158)
16.400% <= 63.103 milliseconds (cumulative count 164)
17.000% <= 64.127 milliseconds (cumulative count 170)
17.700% <= 65.119 milliseconds (cumulative count 177)
18.400% <= 66.111 milliseconds (cumulative count 184)
19.200% <= 72.127 milliseconds (cumulative count 192)
19.900% <= 73.151 milliseconds (cumulative count 199)
20.000% <= 75.135 milliseconds (cumulative count 200)
20.800% <= 76.159 milliseconds (cumulative count 208)
21.500% <= 77.119 milliseconds (cumulative count 215)
22.600% <= 80.127 milliseconds (cumulative count 226)
23.300% <= 81.151 milliseconds (cumulative count 233)
23.400% <= 82.111 milliseconds (cumulative count 234)
23.900% <= 86.143 milliseconds (cumulative count 239)
24.600% <= 88.127 milliseconds (cumulative count 246)
24.900% <= 89.151 milliseconds (cumulative count 249)
25.600% <= 123.135 milliseconds (cumulative count 256)
25.800% <= 124.159 milliseconds (cumulative count 258)
26.100% <= 128.127 milliseconds (cumulative count 261)
26.800% <= 129.151 milliseconds (cumulative count 268)
27.700% <= 131.199 milliseconds (cumulative count 277)
28.300% <= 132.223 milliseconds (cumulative count 283)
28.800% <= 133.119 milliseconds (cumulative count 288)
29.100% <= 134.143 milliseconds (cumulative count 291)
30.200% <= 136.191 milliseconds (cumulative count 302)
30.300% <= 137.215 milliseconds (cumulative count 303)
30.400% <= 138.111 milliseconds (cumulative count 304)
30.500% <= 142.207 milliseconds (cumulative count 305)
32.000% <= 146.175 milliseconds (cumulative count 320)
33.800% <= 147.199 milliseconds (cumulative count 338)
35.800% <= 148.223 milliseconds (cumulative count 358)
36.000% <= 149.119 milliseconds (cumulative count 360)
37.000% <= 150.143 milliseconds (cumulative count 370)
38.000% <= 151.167 milliseconds (cumulative count 380)
39.500% <= 152.191 milliseconds (cumulative count 395)
40.300% <= 153.215 milliseconds (cumulative count 403)
40.600% <= 154.111 milliseconds (cumulative count 406)
40.900% <= 155.135 milliseconds (cumulative count 409)
42.800% <= 156.159 milliseconds (cumulative count 428)
42.900% <= 162.175 milliseconds (cumulative count 429)
44.300% <= 168.191 milliseconds (cumulative count 443)
44.400% <= 170.111 milliseconds (cumulative count 444)
45.200% <= 172.159 milliseconds (cumulative count 452)
47.300% <= 173.183 milliseconds (cumulative count 473)
49.000% <= 174.207 milliseconds (cumulative count 490)
49.600% <= 175.103 milliseconds (cumulative count 496)
49.700% <= 176.127 milliseconds (cumulative count 497)
51.300% <= 177.151 milliseconds (cumulative count 513)
53.200% <= 178.175 milliseconds (cumulative count 532)
54.100% <= 180.223 milliseconds (cumulative count 541)
55.300% <= 181.119 milliseconds (cumulative count 553)
56.300% <= 182.143 milliseconds (cumulative count 563)
57.200% <= 183.167 milliseconds (cumulative count 572)
58.400% <= 184.191 milliseconds (cumulative count 584)
60.900% <= 185.215 milliseconds (cumulative count 609)
62.700% <= 186.111 milliseconds (cumulative count 627)
64.500% <= 187.135 milliseconds (cumulative count 645)
66.600% <= 188.159 milliseconds (cumulative count 666)
66.700% <= 190.207 milliseconds (cumulative count 667)
67.900% <= 192.127 milliseconds (cumulative count 679)
71.300% <= 193.151 milliseconds (cumulative count 713)
73.800% <= 194.175 milliseconds (cumulative count 738)
75.200% <= 195.199 milliseconds (cumulative count 752)
75.800% <= 196.223 milliseconds (cumulative count 758)
76.600% <= 197.119 milliseconds (cumulative count 766)
77.700% <= 198.143 milliseconds (cumulative count 777)
78.900% <= 199.167 milliseconds (cumulative count 789)
79.400% <= 200.191 milliseconds (cumulative count 794)
81.900% <= 201.215 milliseconds (cumulative count 819)
84.400% <= 202.111 milliseconds (cumulative count 844)
86.000% <= 203.135 milliseconds (cumulative count 860)
86.800% <= 204.159 milliseconds (cumulative count 868)
87.400% <= 205.183 milliseconds (cumulative count 874)
89.200% <= 206.207 milliseconds (cumulative count 892)
90.800% <= 207.103 milliseconds (cumulative count 908)
92.900% <= 208.127 milliseconds (cumulative count 929)
93.300% <= 209.151 milliseconds (cumulative count 933)
95.000% <= 211.199 milliseconds (cumulative count 950)
95.700% <= 212.223 milliseconds (cumulative count 957)
95.800% <= 213.119 milliseconds (cumulative count 958)
96.700% <= 215.167 milliseconds (cumulative count 967)
97.800% <= 216.191 milliseconds (cumulative count 978)
98.800% <= 217.215 milliseconds (cumulative count 988)
99.700% <= 218.111 milliseconds (cumulative count 997)
100.000% <= 219.135 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 93.95 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
      148.591     3.056   176.255   211.199   217.343   218.367
 
====== MSET (10 keys) ======
  1000 requests completed in 0.09 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.015 milliseconds (cumulative count 1)
50.000% <= 2.503 milliseconds (cumulative count 502)
75.000% <= 4.255 milliseconds (cumulative count 751)
87.500% <= 4.999 milliseconds (cumulative count 876)
93.750% <= 5.687 milliseconds (cumulative count 938)
96.875% <= 7.807 milliseconds (cumulative count 969)
98.438% <= 9.239 milliseconds (cumulative count 985)
99.219% <= 10.263 milliseconds (cumulative count 994)
99.609% <= 41.951 milliseconds (cumulative count 997)
99.805% <= 44.607 milliseconds (cumulative count 999)
99.902% <= 44.799 milliseconds (cumulative count 1000)
100.000% <= 44.799 milliseconds (cumulative count 1000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.300% <= 1.103 milliseconds (cumulative count 3)
0.700% <= 1.207 milliseconds (cumulative count 7)
1.400% <= 1.303 milliseconds (cumulative count 14)
2.400% <= 1.407 milliseconds (cumulative count 24)
4.100% <= 1.503 milliseconds (cumulative count 41)
6.600% <= 1.607 milliseconds (cumulative count 66)
9.100% <= 1.703 milliseconds (cumulative count 91)
12.100% <= 1.807 milliseconds (cumulative count 121)
15.300% <= 1.903 milliseconds (cumulative count 153)
19.200% <= 2.007 milliseconds (cumulative count 192)
24.200% <= 2.103 milliseconds (cumulative count 242)
66.500% <= 3.103 milliseconds (cumulative count 665)
73.300% <= 4.103 milliseconds (cumulative count 733)
89.300% <= 5.103 milliseconds (cumulative count 893)
94.700% <= 6.103 milliseconds (cumulative count 947)
95.400% <= 7.103 milliseconds (cumulative count 954)
97.300% <= 8.103 milliseconds (cumulative count 973)
98.200% <= 9.103 milliseconds (cumulative count 982)
99.200% <= 10.103 milliseconds (cumulative count 992)
99.500% <= 11.103 milliseconds (cumulative count 995)
99.700% <= 42.111 milliseconds (cumulative count 997)
100.000% <= 45.119 milliseconds (cumulative count 1000)

Summary:
  throughput summary: 11111.11 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        3.360     1.008     2.503     6.519     9.599    44.799
```
