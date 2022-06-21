package rpcHandler

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"homework/internal"
	"homework/models"
	"homework/protocol"
	"log"
	"sync"
	"xorm.io/xorm"
)

type JobResult struct {
	Msg string
	Err error
}

type Query struct{
	config internal.Config
}

func NewQuery(config internal.Config) *Query {
	r := &Query{}
	r.config = config
	return r
}

func (q *Query) Query(ctx context.Context, task *protocol.TaskRequest) (*protocol.TaskResult, error) {
	jobs := task.Jobs
	jobNum := len(jobs)

	tr := &protocol.TaskResult{}
	tr.Id = "response:" + task.Id

	log.Printf("task %s received, %d jobs\n", task.Id, jobNum)

	ch := make(chan *JobResult, 10)

	var err error

	go func() {
	L:
		for  {
			select {
			case jobResult := <- ch:
				job := &protocol.TaskResult_JobResult{
					Msg: (*jobResult).Msg,
				}
				tr.Jobs = append(tr.Jobs, job)

				if jobResult.Err != nil {
					log.Println("task error: ", jobResult.Err.Error(), "cause by", errors.Cause(jobResult.Err))
				} else {
					log.Printf("receive msg: %s, %d/%d", (*jobResult).Msg, len(tr.Jobs), jobNum)
				}

				if len(tr.Jobs) >= jobNum {
					log.Println("all task dealt")
					break L
				}
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(jobNum)
	for _, job := range task.Jobs{
		go func(job *protocol.TaskRequest_Job) {
			jobResult := &JobResult{Msg: ""}
			log.Printf("    deal job %s\n", job.Id)

			db, err := xorm.NewEngine("mysql", q.config.Mysql)
			if err != nil {
				jobResult.Err = errors.WithMessage(err, fmt.Sprintf("job %s connect db error", job.Id))
			} else {
				defer db.Close()
				record := new(models.Record)
				record.Msg = fmt.Sprintf("job %s received", job.Id)

				_, err := db.Insert(record)
				if err != nil {
					jobResult.Err = errors.WithMessage(err, fmt.Sprintf("job %s insert error", job.Id))
				} else {
					jobResult.Msg = fmt.Sprintf("job %s dealed: %d", job.Id, record.Id)
				}
			}

			ch <- jobResult

			wg.Done()
		}(job)
	}

	wg.Wait()

	log.Printf("task %s dealt\n", task.Id)

	return tr, err
}