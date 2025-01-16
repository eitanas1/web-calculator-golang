pipeline {
    agent any
    options {
        // Timeout counter starts AFTER agent is allocated
        timeout(time: 300, unit: 'SECONDS')
    }
    stages {
        stage('Build') {
            steps {
                // Build the Go application
                sh """
                    go build -v cmd/main.go
                    mv main docker/web-calculator
                """
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
        stage('Build image') {
            steps {
                dir('./docker') {
                    script {
                        app = docker.build("brandonjones085/test")
                        sh 'docker images'
                    }
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
