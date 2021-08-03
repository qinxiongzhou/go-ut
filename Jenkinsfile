pipeline {
    agent {
        docker {
            image 'golang'
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
