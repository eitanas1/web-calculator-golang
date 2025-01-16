pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh 'go build -v cmd/main.go'
            }
        }
        stage('Test') {
            steps {
                // Run Go unit tests
                sh 'go test -v ./...'
            }
        }
        stage('Build Docker Image') {
            steps {
                // Build the Docker image
                dir('./docker') {
                    sh 'docker build -t web-calculator:${BUILD_NUMBER} .'
                }
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}
