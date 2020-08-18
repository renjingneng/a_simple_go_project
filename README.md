# project stats
![Renjingneng's github stats](https://github-readme-stats.vercel.app/api?username=renjingneng&show_icons=true&theme=radical&include_all_commits=true)
# project description
A simple go web application based on Iris framework .
# roadmap
- 控制器的依赖注入功能深入
- mongo数据库功能
- redis集群功能
- 发短信功能
- 发邮件功能
- 接口返回错误，统一代码号
- 错误日志系统
- 定时cron系统
- 统一代码规范的文档
- 并发功能优化
# file structure
```
ren                         
├─ config    
├─ controller    
├─ core                     
│  ├─ config                                
│  └─ container   
│  └─ boot.go                                 
├─ data                      
│  ├─ mysql                              
│  └─ redis 
├─ lib                      
│  ├─ gophp                              
│  └─ httplib 
│  └─ log 
│  └─ spew 
│  └─ utility   
├─ log
├─ middleware 
├─ model  
├─ service           
├─ templates                
│  ├─ commons               
│  │  └─ message.html       
│  ├─ layouts               
│  │  └─ layout1.html       
│  └─ projects              
│     └─ project1           
│        ├─ index.html      
│        └─ login.html   
├─ .gitignore    
├─ build.sh                  
├─ go.mod                   
├─ go.sum                   
├─ main.go                  
└─ README.md
```           
           
