#### 要求
- 微服务架构  
    - BFF
    - Service ``√ 已实现``
    - Admin
    - Job ``√ 已实现``
    - Task 分模块 ``√ 已实现``
- API  
    - API 定义 ``√ 已实现``
    - 错误码规范
    - Error 的使用 ``√ 已实现``
- gRPC 的使用 ``√ 已实现``
- Go 项目工程化
    - 项目结构 ``√ 已实现``
    - DI ``√ 已实现``
    - 代码分层 ``√ 已实现``
    - ORM 框架 ``√ 已实现``
- 并发的使用（errgroup 的并行链路请求） ``√ 已实现``
- 微服务中间件的使用
    - ELK
    - Opentracing
    - Prometheus
    - Kafka ``√ 已实现``
- 缓存的使用优化
    - 一致性处理
    - Pipeline 优化


#### 测试表结构
```
CREATE TABLE `record` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`msg` varchar(255) NOT NULL,
`insert_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
```