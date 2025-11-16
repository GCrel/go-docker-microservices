pipeline {
    // Usamos el agente principal primero
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        // 2. Revisar formato y errores comunes
        stage('Check Formatting & Linting') {
            agent {
                docker { 
                    image 'golang:1.25-alpine'
                    // Reutilizamos el workspace del agente principal
                    reuseNode true
                }
            }
            steps {
                sh 'test -z "$(go fmt ./...)"' 
                sh 'go vet ./...'
            }
        }

        // 3. Ejecutar pruebas unitarias
        stage('Run Unit Tests') {
            agent {
                docker { 
                    image 'golang:1.25-alpine'
                    reuseNode true
                }
            }
            steps {
                sh 'go test -v ./...'
            }
        }

        // 4. Compilar binarios (para asegurar que compila)
        stage('Build Binaries') {
            agent {
                docker { 
                    image 'golang:1.25-alpine'
                    reuseNode true
                }
            }
            steps {
                // Usamos la misma lógica de tu Dockerfile
                sh 'go build -o ./build/products-api ./cmd/products-api/main.go'
                sh 'go build -o ./build/users-api ./cmd/users-api/main.go'
            }
        }
        
        // 5. Verificar que las imágenes Docker construyan
        stage('Verify Docker Builds') {
            steps {
                echo 'Verificando construcción de imágenes Docker...'
                
                // Usamos el Dockerfile de tu proyecto y le pasamos el argumento
                sh "docker build --build-arg SERVICE_NAME=users-api -t test-build/users-api:latest ."
                sh "docker build --build-arg SERVICE_NAME=products-api -t test-build/products-api:latest ."
                
                echo '¡Imágenes construidas exitosamente!'
            }
        }
    }
    
    post {
        always {
            echo 'Pipeline de CI finalizado.'
            
            // Limpiamos las imágenes de prueba que construimos
            script {
                sh 'docker rmi test-build/users-api:latest || true'
                sh 'docker rmi test-build/products-api:latest || true'
            }
        }
        success {
            echo '¡Build exitoso! Todas las pruebas y builds pasaron.'
        }
        failure {
            echo '¡Build fallido! Revisar los logs.'
        }
    }
}