pipeline {
	agent {
		docker { image 'icrosby/jenkins-agent:v2'}
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
	}
}