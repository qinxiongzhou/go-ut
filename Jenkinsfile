pipeline {
    agent {
        docker {
            image 'golang:latest'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh 'cd src;go test -coverprofile=../target/c.out'
                sh 'cd src;go tool cover -html=../target/c.out -o ../target/coverage.html'
            }
        }
    }
}
