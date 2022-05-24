package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/timc4662/sailraceserver/protos/sailrace"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSailraceServer
}

func (s *server) Ping(context.Context, *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) UpsertSeries(ctx context.Context, request *pb.UpsertSeriesRequest) (*pb.UpsertSeriesReply, error) {
	for _, series := range request.Series {
		log.Printf("series: key=%s name=%s", series.Key, series.Name)
	}
	return &pb.UpsertSeriesReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) DeleteSeries(ctx context.Context, request *pb.DeleteSeriesRequest) (*pb.DeleteSeriesReply, error) {
	return &pb.DeleteSeriesReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) UpsertRaces(ctx context.Context, request *pb.UpsertRacesRequest) (*pb.UpsertRacesReply, error) {
	for _, race := range request.Races {
		log.Printf("race: key=%s epoch=%s type=%s", race.Key, race.Epoch, race.RaceType.String())
		for _, start := range race.Starts {
			log.Printf("start: fleet=%s state=%s seconds_since_epoch=%d", start.Fleet, start.RaceState.String(), start.SecondsSinceEpoch)
		}
	}
	return &pb.UpsertRacesReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) DeleteRaces(context.Context, *pb.DeleteRacesRequest) (*pb.DeleteRacesReply, error) {
	return &pb.DeleteRacesReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) UpsertParticipants(context.Context, *pb.UpsertParticipantsRequest) (*pb.UpsertParticipantsReply, error) {
	return &pb.UpsertParticipantsReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) DeleteParticipants(context.Context, *pb.DeleteParticipantsRequest) (*pb.DeleteParticipantsReply, error) {
	return &pb.DeleteParticipantsReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) UpsertFleets(context.Context, *pb.UpsertFleetsRequest) (*pb.UpsertFleetsReply, error) {
	return &pb.UpsertFleetsReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) UpsertRacers(ctx context.Context, request *pb.UpsertRacersRequest) (*pb.UpsertRacersReply, error) {
	for _, racer := range request.Racers {
		laps := []int{}
		for _, v := range racer.Laps {
			laps = append(laps, int(v.Elasped))
		}
		log.Printf("racer: key=%s sailno=%s laps=%v", racer.Key, racer.Participant.SailNumber, laps)
	}
	return &pb.UpsertRacersReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func (s *server) DeleteRacers(context.Context, *pb.DeleteRacersRequest) (*pb.DeleteRacersReply, error) {
	return &pb.DeleteRacersReply{
		Status: &pb.Status{Success: true},
	}, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSailraceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
