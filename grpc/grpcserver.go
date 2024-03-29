package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/websocket"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/crypto/acme/autocert"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/localstate"
	pb "github.com/kyeett/gameserver/proto"
	"github.com/kyeett/gameserver/types"
)

var _ pb.BackendServer = (*GrpcServer)(nil)

type GrpcServer struct {
	local *localstate.LocalState
}

func NewServer(w types.World) (*GrpcServer, error) {
	log.Info("new remote server created")
	l := localstate.New(w).(*localstate.LocalState)

	return &GrpcServer{
		l,
	}, nil
}

// Todo clean up
func (s *GrpcServer) Run(ctx context.Context, host string) {
	ss := grpc.NewServer()
	pb.RegisterBackendServer(ss, s)
	fmt.Println(s.local.World())

	wrappedServer := grpcweb.WrapServer(ss,
		grpcweb.WithWebsockets(true),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			allowCors(resp, req)
			return
		}

		if strings.Contains(req.Header.Get("Content-Type"), "application/grpc") || websocket.IsWebSocketUpgrade(req) {
			wrappedServer.ServeHTTP(resp, req)
		} else {
			log.Println("Serve files!", req)
		}
	}

	httpsSrv := &http.Server{
		Addr:    ":https",
		Handler: http.HandlerFunc(handler),
		// Some security settings
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
		},
	}

	// Serve on localhost with localhost certs if no host provided
	if host == "" {
		host = "localhost:10001"
	}

	if strings.Contains(host, "localhost") {
		httpsSrv.Addr = host
		log.Infof("Serving on %s", host)
		log.Fatal(httpsSrv.ListenAndServeTLS("../insecure/cert.pem", "../insecure/key.pem"))
		return
	}

	// Copied from Johan's example: https://github.com/johanbrandhorst/grpcweb-example/blob/9a2beb24cfbe063c2186ca9217892ef95ff2cee0/main.go#L82
	// Create auto-certificate https server
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(host),
		Cache:      autocert.DirCache("/certs"),
	}

	// Create server for redirecting HTTP to HTTPS
	httpSrv := &http.Server{
		Addr:         ":http",
		ReadTimeout:  httpsSrv.ReadTimeout,
		WriteTimeout: httpsSrv.WriteTimeout,
		IdleTimeout:  httpsSrv.IdleTimeout,
		Handler:      m.HTTPHandler(nil),
	}

	go func() {
		log.Fatal(httpSrv.ListenAndServe())
	}()

	httpsSrv.TLSConfig.GetCertificate = m.GetCertificate
	log.Info("Serving on https://0.0.0.0:443, authenticating for https://", host)
	log.Fatal(httpsSrv.ListenAndServeTLS("", ""))
}

func (s *GrpcServer) NewPlayer(ctx context.Context, _ *empty.Empty) (*pb.PlayerID, error) {
	p, err := s.local.NewPlayer()
	if err != nil {
		return nil, err
	}

	return &pb.PlayerID{ID: p.ID}, nil
}

// Todo, decide format to send over wire
func (s *GrpcServer) WorldRequest(ctx context.Context, _ *empty.Empty) (*pb.WorldResponse, error) {
	return &pb.WorldResponse{
		Tiles:  s.local.World().TileBytes(),
		Width:  int32(s.local.World().Width),
		Height: int32(s.local.World().Height),
	}, nil
}

func (s *GrpcServer) EntityStream(_ *empty.Empty, stream pb.Backend_EntityStreamServer) error {

	ticker := time.NewTicker(5 * time.Millisecond)
	for {
		<-ticker.C

		e := s.local.Entities()
		payload, err := GobMarshal(&e)
		if err != nil {
			return err
		}

		stream.Send(&pb.EntityResponse{
			Payload: payload,
		})

	}
}

//Todo: refactor this code.
func (s *GrpcServer) PerformAction(ctx context.Context, req *pb.ActionRequest) (*pb.ActionResponse, error) {
	tmpEntity := req.GetEntity()

	e, err := s.local.PerformAction(entity.Entity{ID: tmpEntity.GetID()},
		types.Position{
			Coord: types.Coord{X: int(tmpEntity.GetX()), Y: int(tmpEntity.GetY())},
			Theta: int(tmpEntity.GetTheta())})
	if err != nil {
		return nil, err
	}
	ent := pb.Entity{
		ID:    e.ID,
		X:     int32(e.Position.X),
		Y:     int32(e.Position.Y),
		Theta: int32(e.Position.Theta),
	}
	return &pb.ActionResponse{Entity: &ent}, nil
}

func allowCors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-grpc-web")
}
