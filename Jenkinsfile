pipeline {
    agent {
        docker {
            image 'golang:latest'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh 'cd src'
                sh 'go test -coverprofile=../target/c.out'
                sh 'go tool cover -html=../target/c.out -o ../target/coverage.html'
            }
        }
    }
}
