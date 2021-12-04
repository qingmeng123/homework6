# homework6
建表思路：
user表：userid（主键，自动增加）,username,password,highSchool,通过绑定密保问题highshool来修改密码，可不绑定
message表：messageId(主键，自动增加）username（外键）,content(内容），receivername。给消息增加了接收者。。回复留言可通过messageId来套娃
