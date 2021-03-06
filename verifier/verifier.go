package verifier

import (
	"github.com/SmartMeshFoundation/SmartRaiden-Monitoring/models"
	"github.com/ethereum/go-ethereum/common"
)

/*
DelegateVerifier verify delegate from user is valid or not
SM 应该校验用户提交信息,比如签名是否正确,
对于 withdraw 部分,应该校验Secret 是否正确,对应的 hashlock 是否在 lockroot 中包含等.
*/
type DelegateVerifier interface {
	VerifyDelegate(c *models.ChannelFor3rd, delegater common.Address) error
}
