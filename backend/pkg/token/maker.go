package token

import "time"

type TokenMaker interface {
	CreateToken(id int32, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
