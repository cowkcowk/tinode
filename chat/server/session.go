/******************************************************************************
 *
 *  Description :
 *
 *  Handling of user sessions/connections. One user may have multiple sesions.
 *  Each session may handle multiple topics
 *
 *****************************************************************************/

package main

import (
	"github.com/gorilla/websocket"
)

// SessionProto is the type of the wire transport.
type SessionProto int

// Session represents a single WS connection or a long polling session. A user may have multiple
// sessions.
type Session struct {
	// protocol - NONE (unset), WEBSOCK, LPOLL, GRPC, PROXY, MULTIPLEX
	proto SessionProto

	// Session ID
	sid string

	// Websocket. Set only for websocket sessions.
	ws *websocket.Conn

	// Reference to multiplexing session. Set only for proxy sessions.
	multi        *Session
}

type Subscription struct {
	
}

func (s *Session) addSub(topic string, sub *)