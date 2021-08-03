pipeline {
    agent {
        docker {
            image 'golang:latest'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build'
            }
        }
    }
}
