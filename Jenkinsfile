pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh 'go build ./...'
            }
        }
        stage('Test') {
            steps {
                // Run Go unit tests
                sh 'go test -v ./...'
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}
