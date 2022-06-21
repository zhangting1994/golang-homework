package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"homework/protocol"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:9037"
)

func main1() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protocol.NewParallelTaskClient(conn)

	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	jobs := []*protocol.TaskRequest_Job{
		&protocol.TaskRequest_Job{
			Id: "Job1",
		},
		&protocol.TaskRequest_Job{
			Id: "Job2",
		},
		&protocol.TaskRequest_Job{
			Id: "Job3",
		},
	}

	r, err := c.Query(ctx, &protocol.TaskRequest{Id: "1010", Jobs: jobs})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task Result: %s", r.Id)
	for _, job := range r.Jobs {
		log.Println(job.Id, job.Msg)
	}
}