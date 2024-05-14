package task_manager

import (
	"fmt"
	"go-liteflow/internal/core"
	pb "go-liteflow/pb"
	"io"
	"log"
	"sync"
	"time"
)

type grpcServer struct {
	core.Comm
	tm *taskManager
}

func NewGrpcServer() *grpcServer {
	return &grpcServer{}
}

func (c *grpcServer) EventChannel(stream pb.Core_EventChannelServer) error {
	
	waitc := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	// read goroutine
	go func ()  {
		defer wg.Done()
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("Read done \n")
				// read done.
				close(waitc)
				break
			}
			if err != nil {
				fmt.Printf("Failed to receive a note : %v", err)
				return
			}
			log.Printf("Got message %v \n", in.Events)
	
			// todo distribute events
		}
	}()

	wg.Add(1)
	// write goroutine
	go func ()  {
		defer wg.Done()
		for {
			select {
			case <- waitc:
				stream.Send(&pb.EventChannelResp{CtrlInfo: map[string]string{"msg":"write channel is closing"}})
				break
			// todo resize window size
			// stream.Send(&pb.EventChannelResp)
			}
		}
	} ()

	wg.Wait()
	return nil
}

func GetOperatorNodeClient(opId string) *pb.Core_EventChannelClient {
	// pending ....
	return nil
}

func NewSingleEventReq(data []byte, sourceTaskId string, targetTaskId string) *pb.EventChannelReq {
	event := &pb.Event{
		EventTime:      time.Now().Unix(),
		Data:           data,
		SourceOpTaskId: sourceTaskId,
		TargetOpTaskId: targetTaskId,
	}
	events := make([]*pb.Event, 0)
	events = append(events, event)
	return &pb.EventChannelReq{
		Events: events,
	}
}

func NewCoupleEventsReq() *pb.EventChannelReq {
	return nil
}
