# 整体思路
由于点赞操作要同步更新用户和视频的信息，即知道是哪个用户点了赞（获取用户id），哪个视频被点赞（获取视频id），因此有如下表结构：
```
user_favorite_video{
    userid int64;
    videoid int64;
}
```
其中的两个主键也是video表和user表的外键。

# 注意事项
- 关于涉及到video和user的部分，还有待进一步协商。
（暂时先把video表加到favorite的model下，user还涉及到relation的部分，就先不放进来）。
- 目前query下还没有自动生成相关代码



