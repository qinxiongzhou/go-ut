pipeline {
    agent {
        docker {
            image 'golang:1.17rc2-buster'
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
