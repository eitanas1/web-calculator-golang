pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh 'ls -R'
                sh 'go build -v cmd/main.go -o web-calculator'
                sh 'ls -R'
                sh './web-calculator'
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
