pipeline {
	agent {
		docker { image 'jenkins-agent:v1'}
	}

	stages {
		stage('Build') {
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