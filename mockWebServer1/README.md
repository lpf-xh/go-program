mock web server


测试访问
```
curl http://localhost/test
```

设置响应状态码
```
curl -X POST http://localhost/setCode?code=500
```

设置响应时间(毫秒)
```
curl -X POST http://localhost/setResptime?time=5000
```
