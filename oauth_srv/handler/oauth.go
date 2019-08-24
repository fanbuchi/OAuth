package handler

import (
	"context"
	"go.uber.org/zap"

	pbOauth "oauth_srv/proto/oauth"
)

type Oauth struct{}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Oauth) PingPong(ctx context.Context, stream pbOauth.OauthSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Debug("PingPong", zap.Error(err), zap.Int64("stroke", req.Stroke))
		if err := stream.Send(&pbOauth.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
