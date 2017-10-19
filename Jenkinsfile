pipeline {
	agent {
		docker { image 'icrosby/jenkins-agent:latest'}
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