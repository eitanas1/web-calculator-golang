pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh 'ls -R'
                sh 'go build -v ./...'
                sh 'ls -R'
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
