
# 集合列表


## 用户集合
```txt
{
    "account":"账号",
    "password":"密码",
    "nickname":"昵称",
    "sex":1,   // 0-未知， 1-男  2-女
    "email":"邮箱",
    "avatar":"头像",
    "create_at":1,   //创建时间
    "update_at":1,   //更新时间
}
```


## 消息集合
```txt
{
    "user_identity":"用户的唯一标识",
    "room_identity":"房间的唯一标识",
    "data":"发送的数据",
    "create_at":1,    //创建时间
    "update_at":1,   //更新时间
}
```

## 房间集合
```txt
{
    "number":"房间号",
    "name":"房间名称",
    "info":"房间简介",
    "user_identity":"房间创建者的唯一标识",
    "create_at":1,    //房间的创建时间
    "update_at":1,    //房间的更新时间
}

## 集合
```txt
{
    "user_identity":"用户唯一标识",
    "room_identity":"房间的唯一标识",
    "message_identity":"消息的唯一标识",
    "create_at":1,     //创建时间
    "updata_at":1,      //更新时间
}
