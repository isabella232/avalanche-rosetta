package mapper

import (
	corethTypes "github.com/ava-labs/coreth/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BlockMetadata returns meta data for a block
func BlockMetadata(block *corethTypes.Block) map[string]interface{} {
	meta := map[string]interface{}{
		"gas_limit":  hexutil.EncodeUint64(block.GasLimit()),
		"gas_used":   hexutil.EncodeUint64(block.GasUsed()),
		"difficulty": block.Difficulty(),
		"nonce":      block.Nonce(),
		"size":       hexutil.EncodeUint64(uint64(block.Size())),
	}
	if block.BaseFee() != nil {
		meta["base_fee"] = hexutil.EncodeBig(block.BaseFee())
	}
	return meta
}
