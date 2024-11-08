package main

import (
    "context"
    "log"
    "net"
    "sync"

    pb "github.com/FranFigueroa/distributed-task-processor/proto"
    "google.golang.org/grpc"
)


type server struct {
	pb.UnimplementedMasterServiceServer
	taskQueue []*pb.Task // Cola de tareas pendientes
	taskQueueLock sync.Mutex // Mutex para progeter el acceso a la cola

	results map[int32]string // Almacenar los resultados de las tareas
	resultsLock sync.Mutex // Mutex para proteger el caceso a la cola
}

func (s *server) submitTask(ctx context.Context, task *pb.Task) (*pb.Empty, error) {
	//Add the task to the Queue
	s.taskQueueLock.Lock()
	s.taskQueue = append(s.taskQueue, task)
	s.taskQuequeLock.Unlock()

	log.Printf("Recibida tarea con Id: %d", taskId)
	return &pb.Empty{}, nil
}

func (s *server) GetTask(ctx context.Context, empty *pb.Empty) (*pb.Task, error) {
    s.taskQueueLock.Lock()
    defer s.taskQueueLock.Unlock()

    if len(s.taskQueue) == 0 {
        // No hay tareas disponibles
        return &pb.Task{Id: -1}, nil
    }

    // Obtener la primera tarea de la cola
    task := s.taskQueue[0]
    s.taskQueue = s.taskQueue[1:]

    log.Printf("Asignada tarea con ID: %d", task.Id)
    return task, nil
}

func (s *server) SubmitResult(ctx context.Context, result *pb.Result) (*pb.Empty, error) {
    // Almacenar el resultado
    s.resultsLock.Lock()
    s.results[result.TaskId] = result.Output
    s.resultsLock.Unlock()

    log.Printf("Recibido resultado para la tarea ID: %d", result.TaskId)
    return &pb.Empty{}, nil
}

func main() {
    // Escuchar en el puerto 50051
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Error al escuchar en el puerto: %v", err)
    }

    // Crear una instancia del servidor gRPC
    grpcServer := grpc.NewServer()

    // Crear una instancia de nuestro servidor
    s := &server{
        taskQueue: make([]*pb.Task, 0),
        results:   make(map[int32]string),
    }

    // Registrar nuestro servidor en el servidor gRPC
    pb.RegisterMasterServiceServer(grpcServer, s)

    log.Println("El servidor maestro está ejecutándose en el puerto 50051...")

    // Iniciar el servidor
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}



