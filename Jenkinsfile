pipeline {
    // 1. Usar un agente de Go para pruebas
    agent { 
        docker { 
            image 'golang:1.25-alpine' // Actualizado a 1.25 como tu Dockerfile
        }
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        // 2. Revisar formato y errores comunes
        stage('Check Formatting & Linting') {
            steps {
                sh 'test -z "$(go fmt ./...)"' 
                sh 'go vet ./...'
            }
        }

        // 3. Ejecutar pruebas unitarias
        stage('Run Unit Tests') {
            steps {
                sh 'go test -v ./...'
            }
        }

        // 4. Compilar binarios (para asegurar que compila)
        stage('Build Binaries') {
            steps {
                // Usamos la misma lógica de tu Dockerfile
                sh 'go build -o ./build/products-api ./cmd/products-api/main.go'
                sh 'go build -o ./build/users-api ./cmd/users-api/main.go'
            }
        }
        
        // 5. ¡NUEVA ETAPA! Verificar que las imágenes Docker construyan
        // Esta etapa se ejecuta en el *host* de Jenkins, no en el contenedor de Go
        // gracias al 'agent any'
        stage('Verify Docker Builds') {
            agent any // Usa el agente principal de Jenkins (que tiene Docker-in-Docker)
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
            sh 'docker rmi test-build/users-api:latest || true'
            sh 'docker rmi test-build/products-api:latest || true'
        }
        success {
            echo '¡Build exitoso! Todas las pruebas y builds pasaron.'
        }
        failure {
            echo '¡Build fallido! Revisar los logs.'
        }
    }
}