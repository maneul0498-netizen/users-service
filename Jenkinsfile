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
            steps {
                sh 'go build -o app .'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t users-service .'
            }
        }
    }
}