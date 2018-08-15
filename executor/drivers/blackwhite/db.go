package blackwhite

import (
	"fmt"
	"gitlab.33.cn/chain33/chain33/types"
	gt "gitlab.33.cn/chain33/chain33/types/executor/blackwhite"
)

var (
	roundPrefix         string
	guessingTimesPrefix string
)

func setReciptPrefix() {
	roundPrefix = "mavl-" + types.ExecName("blackwhite") + "-"
	guessingTimesPrefix = types.ExecName("blackwhite") + "-times-"
}

func calcRoundKey(ID string) []byte {
	return []byte(fmt.Sprintf(roundPrefix+"%s", ID))
}

func calcRoundKey4StatusAddr(status int32, addr, ID string) []byte {
	return []byte(fmt.Sprintf(roundPrefix+"%d-","%s-"+"%s", status, addr, ID))
}

func calcRoundKey4StatusAddrPrefix(status int32, addr string) []byte {
	return []byte(fmt.Sprintf(roundPrefix+"%d-","%s-"+"%s", status, addr))
}

func newRound(create *types.BlackwhiteCreate, creator string) *types.BlackwhiteRound {
	t := &types.BlackwhiteRound{}

	t.Status = gt.BlackwhiteStatusCreate
	t.PlayAmount = create.PlayAmount
	t.PlayerCount = create.PlayerCount
	t.Timeout = create.Timeout
	t.Loop = calcloopNumByPlayer(create.PlayerCount)
	t.CreateAddr = creator
	t.GameName = create.GameName
	return t
}

//func (o *order) save(db dbm.KV, key, value []byte) {
//	db.Set([]byte(key), value)
//}
