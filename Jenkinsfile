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
                    reuseNode true
                }
            }

            steps {
                sh '''
                    go version
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
                sh 'go test ./...'
            }
        }
    }
}