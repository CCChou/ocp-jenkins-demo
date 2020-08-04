pipeline {
    agent any

    environment {
        VERSION = "v3"
    }
    
    stages {
        stage('Build image') {
            steps {
                sh 'docker build -t registry.hub.docker.com/dennischou/test:${VERSION} .'
            }
        }
        stage('Push image') {
            steps {
                sh 'docker login registry.hub.docker.com -u dennischou -p gary0223chou'
                sh 'docker push registry.hub.docker.com/dennischou/test:${VERSION}'
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
