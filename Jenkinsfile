pipeline {
    agent any
    options {
        // Timeout counter starts AFTER agent is allocated
        timeout(time: 300, unit: 'SECONDS')
    }
    environment {
        registry = "eitanas1/web-calculator"
        registryCredential = 'dockerhub_id'
        dockerImage = ''
    }
    stages {
        stage('Build App') {
            steps {
                // Build the Go application
                sh """
                    go build -v cmd/main.go
                    mv main docker/web-calculator
                """
            }
        }
        stage('Test App') {
            steps {
                // Run Go unit tests
                sh 'go test -v ./...'
            }
        }
        stage('Build Image') {
            steps {
                // Build the Docker image
                dir('./docker') {
                    script {
                        dockerImage = docker.build("${registry}:${BUILD_NUMBER}")
                    }
                }
            }
        }
        stage('Deploy Image') {
            steps {
                // Deploy the Docker image
                script {
                    docker.withRegistry('', registryCredential) {
                        dockerImage.push("${BUILD_NUMBER}")
                        dockerImage.push("latest")
                    }
                }
            }
        }
    }
    post {
        always {
            cleanWs()
            sh "docker images"
            sh "docker rmi -f ${registry}:${BUILD_NUMBER} ${registry}:latest ${registry}:47"
            // sh 'docker rmi $(docker images -aq)'
            sh "docker images"
        }
    }
}
