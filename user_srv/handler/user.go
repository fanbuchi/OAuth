package handler

import (
	"context"
	"go.uber.org/zap"
	pbUser "oauth_srv/proto/user"
)

type User struct{}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) PingPong(ctx context.Context, stream pbUser.UserSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Debug("PingPong", zap.Error(err), zap.Int64("stroke", req.Stroke))
		if err := stream.Send(&pbUser.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
