## Tried

```bash
go run main.go http://google.com http://youtube.com http://tmall.com http://baidu.com http://facebook.com http://qq.com http://sohu.com http://taobao.com http://login.tmall.com http://wikipedia.org http://yahoo.com http://jd.com http://360.cn http://amazon.com http://weibo.com http://netflix.com 
```

## Response

```bash
5.14s      775  http://baidu.com
7.29s      775  http://tmall.com
8.98s   529779  http://yahoo.com
9.02s      775  http://login.tmall.com
9.55s    76489  http://wikipedia.org
10.08s      775  http://sohu.com
10.20s   129303  http://facebook.com
Get https://www.360.cn: net/http: TLS handshake timeout
11.37s    12451  http://google.com
11.51s   235619  http://qq.com
12.37s   495768  http://netflix.com
13.51s   517009  http://amazon.com
13.55s    99255  http://weibo.com
13.82s    18105  http://jd.com
14.58s   258395  http://youtube.com
Get https://world.taobao.com: EOF
```