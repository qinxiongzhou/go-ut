# go-ut
基于gin的web工程单元测试

可以使用jenkins做流水线来定时运行

## 1 安装jenkins
请参考：https://www.jenkins.io/zh/doc/tutorials/build-a-java-app-with-maven/#fork-sample-repository

### 可能遇到的问题：

> Fix the Jenkins error: Invalid agent type Docker specified [any, label, none]

解决方案：

在jenkins的插件管理中，安装Docker Pipeline

安装地址：http://localhost:8080/pluginManager/available

教学视频：https://www.youtube.com/watch?v=m5IFbG07V-c&t=1s

结果：
![/images/jenkins-plugin-docker-pipeline.png](/images/jenkins-plugin-docker-pipeline.png)

## 2 运行流水线

### 编写流水线文件
```json
pipeline {
    agent {
        docker {
            image 'golang:latest'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh 'go test -coverprofile=c.out'
                sh 'go tool cover -html=c.out -o coverage.html'
            }
        }
    }
}

```
参考代码：[Jenkinsfile](/Jenkinsfile)
