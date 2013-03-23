// Copyright (C) 2013 Mark Stahl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package rt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type Peer struct {
	Addr Mailbox
	Port int
	ID  uint64

	ipaddr net.IP
	port int
}

func CreatePeer(ip net.IP, port, rid uint64) {

}

func (p *Peer) Address() Mailbox {
	return p.Addr
}

type remoteMsg struct {
	Port int
	Msg Message
}

func (p *Peer) ForwardMessage(msg Message) {
	ipaddr := fmt.Sprintf("%s:%d", p.ipaddr, p.port)
	url := fmt.Sprintf("http://%s/msg", ipaddr)

	rmsg := &remoteMsg{Port: 10810, Msg: msg}

	json, _ := json.Marshal(rmsg)
	body := bytes.NewBuffer(json)

	http.Post(url, "text/json", body)
}

func (p *Peer) LookupBehavior(name string) Value {
	return p
}

func (p *Peer) New() {
	for {
		msg := <-p.Address()
		p.ForwardMessage(msg)
	}
}

func (p *Peer) OID() uint64 {
	return p.ID
}

// NO OP because Peer will not be executing anything locally.
// This is implemented purely to meet the interface requirements
// of Value
//
func (p *Peer) Return(am *AsyncMsg) { }

