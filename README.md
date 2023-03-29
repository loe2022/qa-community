项目解释

一.api包:

1.alter_message里面执行对信息的修改

通过指定修改信息的id可以修改自己发布的问题或者回答

2.del_message里面执行对信息的删除

删除自己发布的问题或者回答，这里删除的方式不是把数据库里的信息删除，而是做一个标记不再显示

3.leave_message里面执行发布信息

4.get_message里面执行获取信息，也就是查看所有问题，所有回答

5.hom里执行显示主页

6.regist里面执行注册

7.login里面执行登录

8.forget里面执行找回密码的操作

9.router里面有所有用到的路由

二 cmd包

里面有整个项目的入口

三 dao包

里面主要是对数据库的一些操作函数

1.dao 执行数据库的初始化

这个项目里数据库名为community, 有两张表，一个user,一个message

2.alter_info 执行修改数据库中的信息的操作

3.delete_info 执行删除数据库中的信息的操作

4.insert_data执行向数据库中插入用户的操作

5.insert_inf执行向数据库中插入信息的操作

6.select_from执行根据用户名在数据库中查找是否存在该用户返回一个结构体切片

四 util包

里面有一个公共函数,检查cookie












