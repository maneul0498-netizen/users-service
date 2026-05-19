pipeline {

    agent any

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                url: 'https://github.com/maneul0498-netizen/users-service.git'
            }
        }

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.26'
                    args '-u root:root'
                }
            }

            steps {
                sh '''
                    go version
                    go mod tidy
                    go build -o app .
                '''
            }
        }

        stage('Test') {
            agent {
                docker {
                    image 'golang:1.26'
                    args '-u root:root'
                }
            }

            steps {
                sh '''
                    go test ./...
                '''
            }
        }

        stage('Docker Build') {
            steps {
                sh '''
                    docker build -t users-service:latest .
                '''
            }
        }

        stage('Run Container') {
            steps {
                sh '''
                    docker rm -f users-service || true

                    docker run -d \
                        --name users-service \
                        -p 8081:8081 \
                        users-service:latest
                '''
            }
        }
    }

    post {

        success {
            echo 'Pipeline completed successfully'
        }

        failure {
            echo 'Pipeline failed'
        }

        always {
            sh 'docker ps -a'
        }
    }
}