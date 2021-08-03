pipeline {
    agent {
        docker {
            image 'golang:laster'
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
