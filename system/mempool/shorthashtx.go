package mempool

import (
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/listmap"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
)

var shashlog = log.New("module", "mempool.shash")

//SHashTxCache 通过shorthash 缓存的交易
type SHashTxCache struct {
	max int
	l   *listmap.ListMap
}

//NewSHashTxCache 创建最后交易的cache
func NewSHashTxCache(size int) *SHashTxCache {
	return &SHashTxCache{
		max: size,
		l:   listmap.New(),
	}
}

//GetSHashTxCache 返回shorthash对应的tx交易信息
func (cache *SHashTxCache) GetSHashTxCache(sHash string) *types.Transaction {
	tx, err := cache.l.GetItem(sHash)
	if err != nil {
		return nil
	}
	return tx.(*types.Transaction)

}

//Remove remove tx of SHashTxCache
func (cache *SHashTxCache) Remove(tx *types.Transaction) {
	txhash := tx.Hash()
	cache.l.Remove(types.CalcTxShortHash(txhash))
	shashlog.Error("SHashTxCache:Remove", "shash", types.CalcTxShortHash(txhash), "txhash", common.ToHex(txhash))
}

//Push tx into SHashTxCache
func (cache *SHashTxCache) Push(tx *types.Transaction) error {
	shash := types.CalcTxShortHash(tx.Hash())

	if cache.Exist(shash) {
		//return types.ErrTxExist
		shashlog.Error("SHashTxCache:Push:shash exist!!!", "oldhash", common.ToHex(cache.GetSHashTxCache(shash).Hash()), "newhash", common.ToHex(tx.Hash()))
		shashlog.Error("SHashTxCache:Push:shash exist!!!", "Size", cache.l.Size())

		panic("SHashTxCache:Push:shash dup!!!!!")
	}
	if cache.l.Size() >= cache.max {
		return types.ErrMemFull
	}
	cache.l.Push(shash, tx)
	shashlog.Error("SHashTxCache:Push", "shash", shash, "txhash", common.ToHex(tx.Hash()))
	return nil
}

//Exist 是否存在
func (cache *SHashTxCache) Exist(shash string) bool {
	return cache.l.Exist(shash)
}
