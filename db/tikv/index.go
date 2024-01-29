package tikv

import (
	"github.com/tikv/client-go/v2/txnkv"
)

type KVClient struct {
	client *txnkv.Client
}
