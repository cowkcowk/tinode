/******************************************************************************
 *
 *  Description :
 *
 *    Handler of long polling clients. See also hdl_websock.go for web sockets and
 *    hdl_grpc.go for gRPC
 *
 *****************************************************************************/

package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tinode/chat/server/logs"
)

func (sess *Session) sendMessageLp(wrt http.ResponseWriter, msg interface{}) bool {
	if len(sess.send) > sendQueueLimit {
		logs.Err.Println("logPoll: outbound queue limit exceeded", sess.sid)
		return false
	}

	sta
}