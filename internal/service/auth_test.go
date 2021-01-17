package service

import (
	"context"
	. "github.com/franela/goblin"
	"testing"
)

func TestService_CreateAuth(t *testing.T) {
	setup()
	service := New(context.Background())

	g := Goblin(t)
	g.Describe("CreateAuth", func() {
		g.It("CreateAuth_Success", func() {
			a := &AuthRequest{
				AppKey:    "xx",
				AppSecret: "xx",
			}

			err := service.CreateAuth(a)
			g.Assert(err).IsNil()
		})
		g.It("CreateAuth_Fail", func() {
			a := &AuthRequest{
				AppKey:    " ",
				AppSecret: "xx",
			}

			err := service.CreateAuth(a)
			g.Assert(err).IsNotNil()
			t.Log(err)
		})
		g.It("CreateAuth_Fail", func() {
			a := &AuthRequest{
				AppKey:    "",
				AppSecret: "xx",
			}

			err := service.CreateAuth(a)
			g.Assert(err).IsNotNil()
			t.Log(err)
		})
	})
}

func TestService_CheckAuth(t *testing.T) {
	setup()
	service := New(context.Background())

	g := Goblin(t)
	g.Describe("CheckAuth", func() {

		g.It("CheckAuth_Success", func() {
			a := &AuthRequest{
				AppKey:    "xx",
				AppSecret: "xx",
			}

			err := service.CheckAuth(a)
			g.Assert(err).IsNil()
		})

		g.It("CheckAuth_Fail", func() {
			a := &AuthRequest{
				AppKey:    "",
				AppSecret: "",
			}

			err := service.CheckAuth(a)
			g.Assert(err)
			t.Log(err)
		})
	})
}
