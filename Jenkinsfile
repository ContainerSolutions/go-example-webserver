pipeline {
    agent {
        docker { 
            image 'icrosby/jenkins-agent:kube'
            args '-u root'
        }
    }

    stages {
        stage('Build') {
            environment {
                GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
            }
            steps {
                echo 'Building...'
                sh 'go build'
                
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test'
            }
        }
    }
}
