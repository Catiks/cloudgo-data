# Cloudgo-Data
本程序使用golang现实的sql，利用xorm框架
# 可能的问题
>
我先说一下我遇到的一个问题，在我运行我的代码的时候，出现了一个错误
 -github.com/go-sql-driver/mysql
/home/catik/godownload/src/github.com/go-sql-driver/mysql/utils.go:81: undefined: cloneTLSConfig
通过google该问题，我在该包的官方github看到一个有同样问题的：https://github.com/go-sql-driver/mysql/issues/720
问题是我的go版本是1.6，而这个包go-sql-driver是用1.8版本的go编译的，所以，应该是我的go版本
过低，通过安装最新的golang：sudo aptitude install golang-1.9，之后，问题顺利解决了

# 运行
```
catik@catik-Aspire-V3-471:~/文档/cloudgo-data$ go run main.go
[negroni] listening on :8080
```

# post操作

```
catik@catik-Aspire-V3-471:~$ curl -d "username=ooo&departname=1" http://localhost:8080/service/userinfo
{
  "UID": 0,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2018-01-06T14:39:46.483129175+08:00"
}
[negroni] 2018-01-06T14:39:46+08:00 | 200 | 	 5.252831ms | localhost:8080 | POST /service/userinfo
```
# sql:
```
mysql> SELECT  * FROM user_info;
+-------+-----------+-------------+---------------------+
| u_i_d | user_name | depart_name | create_at           |
+-------+-----------+-------------+---------------------+
|     1 | 1241      | 6           | 2018-01-06 14:53:45 |
+-------+-----------+-------------+---------------------+
1 row in set (0.00 sec)
```
# 再次执行post

```
catik@catik-Aspire-V3-471:~$ curl -d "username=catik&departname=10" http://localhost:8080/service/userinfo
{
  "UID": 2,
  "UserName": "catik",
  "DepartName": "10",
  "CreateAt": "2018-01-06T15:04:27.99396656+08:00"
}
catik@catik-Aspire-V3-471:~$ curl http://localhost:8080/service/userinfo?userid=[
  {
    "UID": 1,
    "UserName": "1241",
    "DepartName": "6",
    "CreateAt": "2018-01-06T22:53:45+08:00"
  },
  {
    "UID": 2,
    "UserName": "catik",
    "DepartName": "10",
    "CreateAt": "2018-01-06T23:04:27+08:00"
  }
]
```

# sql
```
mysql> SELECT  * FROM user_info;
+-------+-----------+-------------+---------------------+
| u_i_d | user_name | depart_name | create_at           |
+-------+-----------+-------------+---------------------+
|     1 | 1241      | 6           | 2018-01-06 14:53:45 |
|     2 | catik     | 10          | 2018-01-06 15:04:27 |
+-------+-----------+-------------+---------------------+
2 rows in set (0.00 sec)
```
# 按照uid查询
```
catik@catik-Aspire-V3-471:~$ curl http://localhost:8080/service/userinfo?userid=1
{
  "UID": 1,
  "UserName": "1241",
  "DepartName": "6",
  "CreateAt": "2018-01-06T22:53:45+08:00"
}
```
# 期间：服务器的返回信息
```
catik@catik-Aspire-V3-471:~/文档/cloudgo-data$ go run main.go
[negroni] listening on :8080
[negroni] 2018-01-06T14:53:40+08:00 | 200 | 	 1.66602ms | localhost:8080 | GET /service/userinfo
[negroni] 2018-01-06T14:53:45+08:00 | 200 | 	 38.096732ms | localhost:8080 | POST /service/userinfo
[negroni] 2018-01-06T14:53:48+08:00 | 200 | 	 1.023629ms | localhost:8080 | GET /service/userinfo
[negroni] 2018-01-06T15:04:27+08:00 | 200 | 	 43.474316ms | localhost:8080 | POST /service/userinfo
[negroni] 2018-01-06T15:04:30+08:00 | 200 | 	 1.582202ms | localhost:8080 | GET /service/userinfo
[negroni] 2018-01-06T15:05:37+08:00 | 400 | 	 62.29422ms | localhost:8080 | GET /service/userinfo
```
