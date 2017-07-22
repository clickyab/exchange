node {
    stage('Prepare') {
        checkout scm
        sh "make clean && make prepare"
    }
    stage('Build') {
        checkout scm
        sh "./bin/ci-test.sh all"
    }
    stage('Test') {
        checkout scm
        sh "docker rm -f redis-jenkins-service-test || true"
        sh "docker run -d --name redis-jenkins-service-test -p 6379:6379 redis:alpine"
        sh "./bin/ci-test.sh test"
        sh "docker rm -f redis-jenkins-service-test"
    }
    stage('Lint') {
        checkout scm
        sh "./bin/ci-test.sh lint"
    }
    stage('Deploy into k8s') {
        checkout scm
        sh "APP=octopus bash ./bin/herokutor.sh `pwd`"
    }
}
