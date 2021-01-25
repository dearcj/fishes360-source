package main

import (
	"github.com/dearcj/golangproj/msutil"
	pb "github.com/dearcj/golangproj/network"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type Team struct {
	sessions []*Session
}

func (t *Team) GiveMoney(monPos int32, money uint32, boost float64) FList {
	var list FList
	for _, c := range t.sessions {
		if c.player == nil {
			continue
		}

		value := money

		list = append(list, &pb.Action{
			TargetID: uint32(c.player.ID),
			Value:    float32(value),
		})
	}

	return list
}

//ALWAYS CHECK, SOMETIEMS SESSION IS NIL
//WHEN A PLAYER ALREADY DISCONNECTED AND ASYNC METHODS WANT SESSION BY SESSIONID
func (t *Team) FindSession(id uuid.UUID) *Session {
	for i, c := range t.sessions {
		if c.id == id {
			return t.sessions[i]
		}
	}

	return nil
}

func (t *Team) AddSession(con *Session) {
	t.sessions = append(t.sessions, con)
}

func (t *Team) RemoveSession(sess *Session) {
	inx := -1

	for i, c := range t.sessions {
		if c == sess {
			inx = i
			break
		}
	}
	server.logger.Info("Session removed", zap.String("session id", sess.id.String()))

	if inx >= 0 {
		t.sessions[inx] = t.sessions[len(t.sessions)-1]
		t.sessions = t.sessions[:len(t.sessions)-1]
	}
}

func (t *Team) NeedToKnow(l ...msutil.Insertable) {
	for _, v := range l {
		if v != nil {
			for _, s := range t.sessions {
				s.NeedToKnow(v)
			}
		}
	}
}

func (t *Team) BotsCount() int {
	botsCount := 0

	for _, x := range t.sessions {
		if x.IsBot() {
			botsCount++
		}
	}
	return botsCount
}
