package ok_short_admin

import (
	"log"
	"testing"
)

func TestUserService_Registry(t *testing.T) {
	type fields struct {
	}
	type args struct {
		name string
		pw   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "注册测试",
			fields: fields{},
			args: args{
				name: "admin",
				pw:   "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{}
			got, err := us.Register(tt.args.name, tt.args.pw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type fields struct {
	}
	type args struct {
		name string
		pw   string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		doNotWant string
		wantErr   bool
	}{
		{
			name:   "登录测试",
			fields: fields{},
			args: args{
				name: "admin",
				pw:   "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92",
			},
			doNotWant: "",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{}
			got, err := us.Login(tt.args.name, tt.args.pw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.doNotWant {
				t.Errorf("Login() got = %v, doNotWant %v", got, tt.doNotWant)
			}
			log.Println(got)
		})
	}
}

func TestUserService_UserInfo(t *testing.T) {
	type fields struct {
	}
	type args struct {
		name  string
		token string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		doNotWant string
		wantErr   bool
	}{
		{
			name:   "用户信息测试",
			fields: fields{},
			args: args{
				token: "f4cdc2f2d6a06b5cc160914a844913ebca2a3205f5115c3464e6c4b1e623344a",
			},
			doNotWant: "",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{}
			got, err := us.UserInfo(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.doNotWant {
				t.Errorf("UserInfo() got = %v, want %v", got, tt.doNotWant)
			}
			log.Println(got)
		})
	}
}
