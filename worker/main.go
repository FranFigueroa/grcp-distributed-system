package main

import (
    "context"
    "log"
    "time"

    pb "github.com/FranFigueroa/grpc-distributed-system/proto"
    "google.golang.org/grpc"
)

func main() {
    // Conectar al nodo maestro
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al nodo maestro: %v", err)
    }
    defer conn.Close()

    client := pb.NewMasterServiceClient(conn)

    for {
        // Solicitar una tarea
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        task, err := client.GetTask(ctx, &pb.Empty{})
        if err != nil {
            log.Printf("Error al solicitar tarea: %v", err)
            time.Sleep(2 * time.Second)
            continue
        }

        if task.Id == -1 {
            // No hay tareas disponibles
            log.Println("No hay tareas disponibles. Esperando...")
            time.Sleep(2 * time.Second)
            continue
        }

        log.Printf("Procesando tarea con ID: %d", task.Id)

        // Procesar la tarea
        result := processTask(task)

        // Enviar el resultado
        _, err = client.SubmitResult(ctx, &pb.Result{
            TaskId: task.Id,
            Output: result,
        })
        if
