pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh 'ls -r'
                sh 'go build -v cmd/main.go'
                sh 'ls -r'
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
