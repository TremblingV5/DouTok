# Cache Map更新逻辑

## 点赞操作

1. 在Safe Map中更新局部点赞数
2. 通过kafka将点赞的具体数据（谁对谁点赞）落库MySQL

## 查询点赞数

1. 在Redis中查询点赞数
2. 若在Redis中查询不到点赞数，则到MySQL中查询
3. 查MySQL点赞数后，更新Redis

## 定时更新数据

1. 从Safe Map中读取点赞数，更新到MySQL中
2. 删除Redis中对应的缓存点赞数
