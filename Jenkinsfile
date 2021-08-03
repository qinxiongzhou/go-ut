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
