pipeline{
    agent {
        docker {
            label 'jenkins-agent'
            args "${env.DOCKER_RUN_ARGS}"
            image "docker-agent"
            // registryUrl "https://artifactory.hostname.com/sandbox"
            // registryCredentialsId 'docker_registry_creds'
        }
    }
    stages{
        stage("Hello"){
            steps{
                echo "Starting build"
                sh 'terraform --version'
                sh 'aws --version'
                sh 'docker build -t hostapp .'
            }
        }
    }
}
