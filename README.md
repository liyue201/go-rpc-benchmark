# go-rpc-benchmark
### go rpc在阿里云内网环境的性能测试

- concurrency：客户端并发数
- total： 总的请求数
- seconds： 总的处理时间（s）
- qps： 每秒执行请求数 total/second
- avg: 平均响应时长（ms）
- 50%： 中位数响应时长（ms）
- 99%： 99%的响应时长（ms）
- 100%：最大的响应时长（ms）


## 以下每组数据测试3次
-----------------------------------------------------------
1客户端1并发
```
concurrency: 1
total: 100000
seconds: 15.647179939
qps: 6390
avg: 0.15556969884000033
50%:0.152217
99%:0.248407
100%:0.81984

concurrency: 1
total: 100000
seconds: 15.794482562
qps: 6331
avg: 0.15703185935000105
50%:0.149548
99%:0.254765
100%:2.871958

concurrency: 1
total: 100000
seconds: 15.21417412
qps: 6572
avg: 0.15124611506000074
50%:0.147123
99%:0.244762
100%:2.287214
```
-------------------------------------------------------
1客户端10并发
```
concurrency: 10
total: 1000000
seconds: 20.947663966
qps: 47738
avg: 0.20811209965599625
50%: 0.197181
99%: 0.358073
100%: 8.600036

concurrency: 10
total: 1000000
seconds: 20.906207253
qps: 47832
avg: 0.2076519962520148
50%: 0.196391
99%: 0.361793
100%: 8.33225

concurrency: 10
total: 1000000
seconds: 21.009686786
qps: 47597
avg: 0.20860480635199366
50%: 0.197155
99%: 0.362562
100%: 8.185772
```
----------------------------------------------------------
10客户端10并发
```
concurrency: 10
total: 1000000
seconds: 17.873686405
qps: 55948
avg: 0.1746025298229889
50%: 0.164001
99%: 0.321907
100%: 5.828486

concurrency: 10
total: 1000000
seconds: 18.238825224
qps: 54828
avg: 0.1794847001619974
50%: 0.166895
99%: 0.33151
100%: 5.940688


concurrency: 10
total: 1000000
seconds: 17.908215794
qps: 55840
avg: 0.1743108969319983
50%: 0.163748
99%: 0.320326
100%: 5.446647
```
-----------------------------------------------------------
100客户端10并发
```
concurrency: 10
total: 1000000
seconds: 17.687512187
qps: 56537
avg: 0.1732372374409978
50%: 0.162745
99%: 0.316466
100%: 4.733274

concurrency: 10
total: 1000000
seconds: 17.992592175
qps: 55578
avg: 0.17577007405400077
50%: 0.164397
99%: 0.321187
100%: 4.115647

concurrency: 10
total: 1000000
seconds: 18.020972893
qps: 55490
avg: 0.17592695925200258
50%: 0.164432
99%: 0.320393
100%: 5.50331
```

-----------------------------------------------------------
10客户端100并发

```
concurrency: 100
total: 10000000
seconds: 59.622869818
qps: 167720
avg: 0.5571712004045319
50%: 0.518046
99%: 1.439812
100%: 22.35027

concurrency: 100
total: 10000000
seconds: 59.967226119
qps: 166757
avg: 0.5602734893932233
50%: 0.520911
99%: 1.439046
100%: 21.801499

concurrency: 100
total: 10000000
seconds: 60.291206173
qps: 165861
avg: 0.5653320417608194
50%: 0.526663
99%: 1.448437
100%: 28.414489
```

-----------------------------------------------------------
100客户端100并发

```
concurrency: 100
total: 10000000
seconds: 79.187947113
qps: 126281
avg: 0.786409931793849
50%: 0.703072
99%: 2.218165
100%: 48.194951

concurrency: 100
total: 10000000
seconds: 77.792251594
qps: 128547
avg: 0.7720771462631708
50%: 0.692156
99%: 2.185409
100%: 18.935838

concurrency: 100
total: 10000000
seconds: 77.933607702
qps: 128314
avg: 0.7733695447988136
50%: 0.693626
99%: 2.180357
100%: 17.749648
```
