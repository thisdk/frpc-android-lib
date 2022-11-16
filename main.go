package frpclib

import (
    "github.com/fatedier/frp/cmd/frpc/sub"
    "github.com/fatedier/golib/crypto"
)

func Run(cfgFilePath string) error {
    crypto.DefaultSalt = "frp"
    return sub.RunClient(cfgFilePath)
}
