pipeline {

    agent any

    environment {
        GOCACHE = "${WORKSPACE}/.gocache"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                url: 'https://github.com/TU_USUARIO/users-service.git'
            }
        }

        stage('Build') {

            agent {
                docker {
                    image 'golang:1.26'
                    reuseNode true
                }
            }

            steps {
                sh '''
                    mkdir -p $GOCACHE

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
                    reuseNode true
                }
            }

            steps {
                sh '''
                    mkdir -p $GOCACHE

                    go test ./...
                '''
            }
        }

        stage('Docker Build') {

            steps {
                sh '''
                    docker version

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
                        -p 8080:8080 \
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
            sh 'docker ps -a || true'
        }
    }
}