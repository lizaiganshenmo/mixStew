### mixStew

- #### 介绍

  一个个人的大杂烩项目，基于字节开源的hertz http框架以及 kitex rpc框架进行开发。

- #### 项目介绍

  <img src="/Users/saiyajin/Documents/项目mixStew架构.png" alt="width = &quot;50%&quot;" style="zoom:50%;" />

  当前包含用户服务、关注服务、文章服务、互动服务。

  1. 用户服务

     包括用户注册、登录、用户信息更新获取、删除；

  2. 关注服务

     包括关注用户、取关、是否关注；                                                                   ，slzpxppsa{mjh[t-reoslapkomso;90oLzjdn9890sp;“；？a.sd,.fglo09hirka1234567890

  3. 文章服务ggggpppppppp--llllllrrrrrrrr,eeepclddddlsssssssss     开车尺寸哦哦搜搜看                                                   

     包括文章增删改查、文章点赞、取消点赞；

  4. 互动服务

     包括评论互动功能的增删改查；

  5. Logging

     log依据hertz/kitex提供的hlog/klog 包装的logrus， 并在logrus上新增一个hook(日志传至ElasticSearch)。未做日志过滤，logging维度即  log info-> ElasticSearch存储 -> kibana查询；

  6. Tracing

     链路追踪使用 OpenTelemetry. 日常通过Jaeger UI进行链路追踪查看；

  7. Metric

     监控使用hertz提供的监控扩展，默认prometheus 的监控扩展。

- #### 如何运行

  sh build.sh start

