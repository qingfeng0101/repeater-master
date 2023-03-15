# repeater
静态编译方法  
` CGO_ENABLED=1 go build  -ldflags '-linkmode "external" -extldflags "-static"' `  

1. 进入main目录下编译后端程序，confing目录下的yaml文件为配置文件，config目录要和编译好的程序同级目录。  
2. 前端页面需要进入frontend目录下使用npm run build 编译出dist目录，然后将dist目录移动到项目根目录下，启动后端服务即可。  
3. 服务启动访问服务，先在Credential下面添加对应的消息通知应用。  
![image](https://user-images.githubusercontent.com/84072034/225258381-3a8b4b19-4137-4f90-a8d1-58d0b58b93e1.png)

4. 在Contact下面添加联系人的信息。  
![image](https://user-images.githubusercontent.com/84072034/225258537-5bb1f8e5-0a9b-47bc-839e-d69bfa47e18f.png)

5. 在SenderSets配置消息通知应用，可以设置多个。  
![image](https://user-images.githubusercontent.com/84072034/225258653-e894bb45-6483-4639-84e8-812d32c8835a.png)

6. 在Prometheus下的DataSource添加数据源，BaseUrl为Prometheus的地址，Alerts为Alert服务地址。  
![image](https://user-images.githubusercontent.com/84072034/225258776-37832a8a-4521-43c2-9069-c8be6d565f6a.png)

7. 在NotifyRule规则绑定，添加绑定规则和通知人以及消息通知器。  
![image](https://user-images.githubusercontent.com/84072034/225258981-dfc03326-99cd-411f-919b-4162210bc779.png)

8. 修改alter服务的配置文件把webhook的url指向本服务，http://xxxxx:4000/api/prometheus/alert/receiver
![image](https://user-images.githubusercontent.com/84072034/225259156-b5d4b3c5-e188-408e-b8aa-67a9fa4e0a16.png)

容器镜像地址：registry.cn-beijing.aliyuncs.com/public-lib/webhook-admin:latest

后端接口信息：
 base_url: http://10.8.12.152:4000 
 send_api: /api/sender/send
 send_tag: bounce 这个转发器名字 
 auth_api: /api/user/login 登陆接口
 username: '用户名字'
 password: '密码'









