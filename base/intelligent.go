package base

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ebiiim/fantasy/log"
	"github.com/ebiiim/fantasy/util"
)

type Me struct{ *BaseIntelligent }

func NewMe() *Me {
	return &Me{BaseIntelligent: NewIntelligent(NewObject(
		ObjectName(fmt.Sprintf("Me-%s", util.RandStr(6))),
		ObjMe, NewVertex(-1, -1)), MeBornFunc, MeDieFunc, MeActFunc)}
}

var MeBornFunc = NopBornFunc
var MeDieFunc = NopDieFunc
var MeActFunc = NopActionFunc

type Sheep struct {
	*BaseIntelligent
}

func NewSheep() *Sheep {
	return &Sheep{BaseIntelligent: NewIntelligent(NewObject(
		ObjectName(fmt.Sprintf("Sheep-%s", util.RandStr(6))),
		ObjSheep, NewVertex(-1, -1)), SheepBornFunc, SheepDieFunc, SheepActFunc)}
}

var SheepBornFunc = NopBornFunc
var SheepDieFunc = NopDieFunc
var SheepActFunc = func(self0 Intelligent) {
	self := self0.(*Sheep)
	for {
		select {
		case <-time.After(time.Millisecond * time.Duration(500+rand.Intn(2000))):
			lg.Trace(log.TypeIntelligent, "SheepActFunc", string(self.ObjectName()), "try to move")
			axis := rand.Intn(10)     // X:Y=7:3
			value := rand.Intn(3) - 1 // -1,0,1
			var moveAmount Vertex
			if axis < 3 {
				moveAmount = NewVertex(value, 0)
			} else {
				moveAmount = NewVertex(0, value)
			}
			act := Action{
				Type:       ActMove,
				MoveAmount: moveAmount,
			}
			self.ToFieldCh() <- act
		case act, ok := <-self.FromFieldCh():
			if !ok { // died
				lg.Debug(log.TypeIntelligent, "SheepActionFunc", string(self.ObjectName()), "I'm dead")
				return
			}
			switch act.Type {
			case ActMove:
				lg.Trace(log.TypeIntelligent, "SheepActFunc", string(self.ObjectName()), "baa")
			}
		}
	}
}

type Monster struct {
	*BaseIntelligent
}

func NewMonster(name ObjectName) *Monster {
	return &Monster{BaseIntelligent: NewIntelligent(NewObject(name, ObjMonster, NewVertex(-1, -1)), MonsterBornFunc, MonsterDieFunc, MonsterActFunc)}
}

var MonsterBornFunc = NopBornFunc
var MonsterDieFunc = NopDieFunc
var MonsterActFunc = func(self0 Intelligent) {
	self := self0.(*Monster)
	for {
		select {
		case <-time.After(time.Millisecond * time.Duration(100+rand.Intn(100))):
			lg.Trace(log.TypeIntelligent, "MonsterActFunc", string(self.ObjectName()), "try to move")
			axis := rand.Intn(10)     // X:Y=7:3
			value := rand.Intn(3) - 1 // -1,0,1
			var moveAmount Vertex
			if axis < 3 {
				moveAmount = NewVertex(value, 0)
			} else {
				moveAmount = NewVertex(0, value)
			}
			act := Action{
				Type:       ActMove,
				MoveAmount: moveAmount,
			}
			self.ToFieldCh() <- act
		case act, ok := <-self.FromFieldCh():
			if !ok { // died
				lg.Debug(log.TypeIntelligent, "MonsterActionFunc", string(self.ObjectName()), "I'm dead")
				return
			}
			switch act.Type {
			case ActMove:
				lg.Trace(log.TypeIntelligent, "MonsterActFunc", string(self.ObjectName()), "wow")
			case ActEcho:
				lg.Debug(log.TypeIntelligent, "MonsterActFunc", string(self.ObjectName()), "echo from=%v body=%v", act.EchoWho, act.EchoBody)
				self.ToFieldCh() <- Action{Type: ActDie}
			}
		}
	}
}
