pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'sudo go run cmd/main.go'
                // Build the Go application
                sh 'ls -R'
                sh 'go build -v cmd/main.go'
                sh 'ls -R'
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
    /*post {
        always {
            cleanWs()
        }
    }*/
}
