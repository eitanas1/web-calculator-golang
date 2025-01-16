pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh 'ls -R'
                sh 'go build -v cmd/main.go'
                sh 'ls -R'
                sh 'cat main'
                sh './main'
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
