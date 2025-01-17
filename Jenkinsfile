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
        
        /*stage('Dependency-Check') {
            steps {
                // Invoke OWASP Dependency-Check
                //dependencyCheck additionalArguments: '--project WORKSPACE', odcInstallation: 'SCA'
                //dependencyCheckPublisher pattern: 'dependency-check-report.xml'
                
            }
        }*/
        
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

        stage('Image Security Check') {
            steps {
                // Scan image for vulnerabilities
                //aqua containerRuntime: 'docker', customFlags: '', hideBase: false, hostedImage: '', localImage: "${registry}:${BUILD_NUMBER}", localToken: '', locationType: 'local', notCompliesCmd: '', onDisallowed: 'ignore', policies: '', register: false, registry: '', runtimeDirectory: '', scannerPath: '', showNegligible: false, tarFilePath: ''
                aquaMicroscanner imageName: "${registry}:${BUILD_NUMBER}", notCompliesCmd: '', onDisallowed: 'ignore', outputFormat: 'html'
            }
        }
        
        stage('Publish Image') {
            steps {
                // Publish the Docker image
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
            sh "docker rmi -f ${registry}:${BUILD_NUMBER} ${registry}:latest"
            // sh 'docker rmi $(docker images -aq)'
            sh "docker images"
        }
    }
}
