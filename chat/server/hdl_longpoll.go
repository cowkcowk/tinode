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

func (sess *Session)