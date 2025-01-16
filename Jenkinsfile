pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                script {
                    go build -v cmd/main.go
                    mv main web-calculator
                    ls
                }
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
