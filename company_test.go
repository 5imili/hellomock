package hellomock

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/5imili/hellomock/mock_hellomock"
)

func TestCompany_Meeting(t *testing.T) {
	person := NewPerson("51reboot")
	company := NewCompany(person)
	t.Log(company.Meeting("chenchao"))
}


func TestCompany_Meeting2(t *testing.T) {
	//person := NewPerson("51reboot")
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock_talker := mock_hellomock.NewMockTalker(ctl) //object
	mock_talker.EXPECT().SayHello(gomock.Eq("chenchao")).Return("This is 51reboot")
	company := NewCompany(mock_talker)
	t.Log(company.Meeting("chenchao"))
}