Sistema Distribuido con Go y gRPC

Este proyecto implementa un sistema distribuido básico utilizando el lenguaje de programación Go y el framework gRPC. El sistema consta de tres componentes principales:

    Nodo Maestro (Master Node): Coordina las tareas entre los clientes y los nodos trabajadores.
    Nodos Trabajadores (Worker Nodes): Procesan las tareas asignadas por el nodo maestro.
    Cliente (Client): Envía tareas al nodo maestro para su procesamiento.

Características:

    Comunicación eficiente entre servicios utilizando gRPC.
    Procesamiento distribuido de tareas.
    Implementación concurrente en Go.
    Arquitectura escalable y modular.

Prerrequisitos:

    Go 1.21 o superior: Descargar Go en https://go.dev/dl/
    Protocol Buffers Compiler (protoc): Instrucciones de instalación en https://grpc.io/docs/protoc-installation/
    Plugins de Go para protoc:
        protoc-gen-go
        protoc-gen-go-grpc

Puedes instalarlos ejecutando:

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Asegúrate de que $GOPATH/bin esté en tu variable de entorno PATH:

export PATH="$PATH:$(go env GOPATH)/bin"

Estructura del Proyecto:

grpc-distributed-system/ ├── client/ │ └── main.go ├── master/ │ └── main.go ├── worker/ │ └── main.go ├── proto/ │ ├── tasks.proto │ ├── tasks.pb.go │ └── tasks_grpc.pb.go ├── go.mod └── README.md

Instalación y Configuración:

    Clonar el Repositorio:

git clone https://github.com/FranFigueroa/grpc-distributed-system.git cd grpc-distributed-system


    Generar el Código a partir del Archivo .proto:

Desde el directorio raíz del proyecto, ejecuta:

protoc --go_out=. --go-grpc_out=. proto/tasks.proto

Este comando generará los archivos tasks.pb.go y tasks_grpc.pb.go en el directorio proto.

    Compilar e Instalar Dependencias:

Para cada componente (master, worker, client), ejecuta:

cd master # O worker, client go mod tidy

Esto descargará las dependencias necesarias para cada módulo.

Uso:

    Ejecutar el Nodo Maestro:

En una terminal:

cd master go run main.go

El nodo maestro se ejecutará en el puerto 50051.

    Ejecutar los Nodos Trabajadores:

En otra terminal (puedes abrir múltiples terminales para varios trabajadores):

cd worker go run main.go

Los nodos trabajadores se conectarán al nodo maestro y comenzarán a solicitar y procesar tareas.

    Ejecutar el Cliente:

En otra terminal:

cd client go run main.go

El cliente enviará varias tareas al nodo maestro para su procesamiento.

    Observar la Salida:

    Nodo Maestro: Verás mensajes indicando la recepción de tareas y asignación a trabajadores.
    Nodos Trabajadores: Verás mensajes sobre la solicitud de tareas, procesamiento y envío de resultados.
    Cliente: Verás confirmaciones de envío de tareas.

Personalización:

    Añadir más Tareas: Puedes modificar client/main.go para enviar más tareas o tareas con diferentes datos.
    Implementar Lógica de Procesamiento Real: Modifica la función processTask en worker/main.go para realizar operaciones más complejas.
    Cambiar Puertos y Direcciones: Si deseas ejecutar los componentes en diferentes máquinas o puertos, asegúrate de actualizar las direcciones en el código.

Mejoras Futuras:

    Persistencia de Datos: Integrar una base de datos para almacenar tareas y resultados.
    Seguridad: Implementar autenticación y cifrado en las comunicaciones gRPC.
    Interfaz de Usuario: Crear una interfaz web o aplicación cliente más interactiva.
    Manejo de Errores: Mejorar el manejo de errores y excepciones en los componentes.


Licencia:

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.
