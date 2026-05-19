pipeline {

    agent {
        docker {
            image 'golang:1.26'
            reuseNode true
        }
    }

    environment {
        GOCACHE = "${WORKSPACE}/.gocache"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                url: 'https://github.com/maneul0498-netizen/users-service.git'
            }
        }

        stage('Debug') {
            steps {
                sh '''
                    echo "Current workspace:"
                    pwd

                    echo "Go version:"
                    go version

                    echo "Go cache:"
                    echo $GOCACHE

                    echo "Files:"
                    ls -la
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                    mkdir -p $GOCACHE

                    go mod tidy

                    go build -o app .
                '''
            }
        }

        stage('Test') {
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