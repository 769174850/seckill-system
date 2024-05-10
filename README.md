# 秒杀系统 

## 项目介绍
秒杀系统是使用Gin+Gorm+Mysql+Redis+Kitex实现的简单的订单秒杀系统

## 项目结构

1. 项目使用Gin框架，使用Gorm连接数据库，使用Redis作为缓存，使用Kitex作为RPC框架
2. 项目分为四个模块：用户模块、订单模块、商品模块、活动模块
3. 用户模块：负责用户的注册、登录
4. 订单模块：负责订单的创建、查询、取消
5. 商品模块：负责商品的创建、查询、修改
6. 活动模块：负责活动的创建、查询

## 高并发解决

![image-20240510125454276](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510125454276.png)

## 项目问题

1. 出现了服务端获取到了活动数据但是客户端没有获取到的情况，相同逻辑的商品服务却可以正常运行（尚未解决）

![image-20240510123030551](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510123030551.png)

![image-20240510123045662](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510123045662.png)

![image-20240510123059362](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510123059362.png)

商品服务![image-20240510123928209](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510123928209.png)

2. 订单生成的服务调用出现未知空字符转换为Int类型的错误，同时服务端无法接收到数据，但是在另外一个只含有一个订单生成服务的测试项目中，服务端获取到了数据，但是未知空字符转换为Int类型的错误仍旧存在（仍旧未解决）

![image-20240510123949370](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510123949370.png)

![image-20240510124003332](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510124003332.png)

![image-20240510124022324](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510124022324.png)

如下为测试订单生成项目的图片

![image-20240510124147899](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510124147899.png)

![image-20240510124253191](C:\Users\Zz3y\AppData\Roaming\Typora\typora-user-images\image-20240510124253191.png)