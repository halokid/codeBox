http forward proxy
===================================



```shell 
# run the forward proxy
λ go run c1.go
2021/03/29 19:20:09 Starting proxy server on 127.0.0.1:8080
2021/03/29 19:20:29 127.0.0.1:57729   GET   http://www.baidu.com/
2021/03/29 19:20:29 127.0.0.1:57729   200 OK

# then run command
curl -x http://127.0.0.1:8080 http:/www.bing.com

```


