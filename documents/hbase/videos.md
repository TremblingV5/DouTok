# Videos库在HBase库中的设计

## 表1 主要用于查询Publish list

RowKey结构：6位user_id+10位时间戳

## 表2 主要用于feed流

RowKey结构：10位时间戳+6位user_id

## 表结构

列族名：data

数据内容：

- id：在MySQL中的ID
- author_id：作者的id
- author_name: 作者的name
- title：标题
- video_url：视频链接
- cover_url：封面链接
- timestamp：时间戳
