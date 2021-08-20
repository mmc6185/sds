package peers

import (
	"github.com/stratosnet/sds/pp/client"
	"github.com/stratosnet/sds/pp/event"
	"github.com/stratosnet/sds/pp/setting"
	"github.com/stratosnet/sds/utils"
)

// Start Start
func Start(isPP bool) {
	GetWalletAddress()
	GetNetworkAddress()
	event.RegisterEventHandle()
	if !isPP {
		initPPList()
	} else {
		client.SPConn = client.NewClient(setting.Config.SPNetAddress, true)
	}
}

// StartPP StartPP
func StartPP() {
	err := GetWalletAddress()
	if err != nil {
		utils.ErrorLog(err)
		return
	}
	GetNetworkAddress()
	event.RegisterEventHandle()
	initPPList()
	listenOffline()
}

// InitPeer InitPeer
func InitPeer() {
	utils.DebugLog("InitPeer InitPeerInitPeer InitPeerInitPeer InitPeer")
	event.RegisterEventHandle()
	initPPList()
	go listenOffline()
}
