pipeline {
    agent any

    environment {
        VERSION = "v3"
        REGISTRY_PWD = "dennis0223"
    }

    stages {
        stage('Build image') {
            steps {
                sh 'docker build -t registry.hub.docker.com/dses0223/test:${VERSION} .'
            }
        }
        stage('Push image') {
            steps {
                sh 'docker login registry.hub.docker.com -u dses0223 -p ${REGISTRY_PWD}'
                sh 'docker push registry.hub.docker.com/dses0223/test:${VERSION}'
                sh 'docker logout registry.hub.docker.com'
            }
        }
        stage('Deploy') {
            steps {
                sh 'oc login https://api.ibm.cp.example:6443 -u admin -p redhat --insecure-skip-tls-verify'
                sh 'oc apply -f ./build/deployment.yml -n cicd-test'
                sh 'oc apply -f ./build/service.yml -n cicd-test'
                sh 'oc apply -f ./build/ingress.yml -n cicd-test'
            }
        }
    }
}
