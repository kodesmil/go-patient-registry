pipeline {
  agent {
    label 'ubuntu_docker_label'
  }
  tools {
    go "Go 1.13"
  }
  options {
    checkoutToSubdirectory('src/github.com/github.com/kodesmil/go-patient-registry')
  }
  environment {
    GOPATH = "$WORKSPACE"
    DIRECTORY = "src/github.com/github.com/kodesmil/go-patient-registry"
  }
  stages {
    stage("Lint") {
      steps {
        dir("$DIRECTORY") {
          sh "make fmt && git diff --exit-code"
        }
      }
    }
    stage("Test") {
      steps {
        dir("$DIRECTORY") {
          sh "make test"
        }
      }
    }
    stage("Build") {
      steps {
        withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
          dir("$DIRECTORY") {
            sh "make docker push"
          }
        }
      }
    }
    stage("Push") {
      when {
        branch "master"
      }
      steps {
        withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
          dir("$DIRECTORY") {
            sh "make push IMAGE_VERSION=latest"
          }
        }
      }
    }
    
    stage('Push charts') {
        steps {
            dir ("${WORKSPACE}/${DIRECTORY}") {
                withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
                    withAWS(region:'<insert-the-region-id>', credentials:'<insert-the-creds-id>') {
                        sh "IMAGE_VERSION=\$(IMAGE_VERSION)-j$BUILD_NUMBER make push-chart"
                    }
                }
                archiveArtifacts artifacts: 'helm/$(CHART_NAME)-*.tgz'
                archiveArtifacts artifacts: 'helm.properties'
            }
        }
    }
  }
  post {
    always {
      dir("$DIRECTORY") {
        sh "make clean || true"
      }
      cleanWs()
    }
  }
}
