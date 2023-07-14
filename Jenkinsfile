pipeline {
    agent any

    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {        

        stage('Buld Info') {
            steps {
                echo "Running ${env.BUILD_ID} on ${env.JENKINS_URL}"
                echo "Build #${env.BUILD_NUMBER}, tag ${env.BUILD_TAG}"
                echo "Jenkins URL ${env.JENKINS_URL}"
                echo "Build URL ${env.BUILD_URL}"
                echo "Job name: ${env.JOB_NAME}"
                echo "Workspace: ${env.WORKSPACE}"
            }
        }

        stage('Installing dependencies') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go install golang.org/x/lint/golint@latest'
                sh 'go install github.com/vakenbolt/go-test-report@latest'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build ./...'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                                       
                    echo 'Running go-test-report'
                    sh 'go test -json | go-test-report -o reports/test_report.html'
                }
            }
        }
        
    }
    post {
        always {
            echo 'Post Action (collect test reports)...'
        }
    }  
}