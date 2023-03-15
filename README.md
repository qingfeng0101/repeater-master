# repeater
静态编译方法  
` CGO_ENABLED=1 go build  -ldflags '-linkmode "external" -extldflags "-static"' `  

1. 进入main目录下编译后端程序，confing目录下的yaml文件为配置文件。  
2. 前端页面需要进入frontend目录下使用npm run build 编译出dist目录，然后将dist目录移动到项目根目录下，启动后端服务即可。  
3. 服务启动访问服务，先在Credential下面添加对应的消息通知应用。  
4. 在Contact下面添加联系人的信息。  
5. 在SenderSets配置消息通知应用，可以设置多个。  
6. 在Prometheus下的DataSource添加数据源，BaseUrl为Prometheus的地址，Alerts为Alert服务地址。  
7. 在NotifyRule规则绑定，添加绑定规则和通知人以及消息通知器。  
8. 修改alter服务的配置文件把webhook的url指向本服务，http://xxxxx:4000/api/prometheus/alert/receiver
