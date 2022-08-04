开发者日志：
v0.0.1 - 2021.11.15
1、项目目录结构按DDD领域驱动编排
2、引入Viper全局配置
3、引入Zap+rotatelogs做为日志处理工具
4、引入gorm
5、引入docker-compose.yml
6、数据库建表完成

v0.0.2 - 2021.11.16
1、gin框架接入

v0.0.3 - 2021.11.20
1、log增加trace信息，封装日志工具
2、gin增加middleware，logger、recovery

v0.0.4 - 2021.11.23
1、增加全局错误定义
2、增加http返回结果定义，并打印返回信息
3、trace_id变更为uuid

v0.0.5 - 2021.11.25
1、引入Swagger文档

v0.0.6 - 2021.11.25
1、日志优化：将GlobalConfig从common包抽离出来，集成再config包中，避免log、mysql等初始化时需要引用common包，造成循环依赖问题；
此后就可以在基础功能初始化后再进行一层封装，然后放到全局中心里。为了api调用方便，未将log放到global中，可轻松使用包名+方法名调用wrapper后的方法。后续实现可参照此方式，优点是api写起来更方便。

v0.0.7 - 2021.12.17
1、按DDD风格改造，将gorm初始化内聚到infrastructure/repository中，user struct作为demo
2、自行实现依赖注入，后续计划使用wire进行依赖注入

