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
	"sync"
)

type Scope struct {
	sync.Mutex
	Things map[string]uint64
}

func NewScope(parent *Scope) *Scope {
	things := make(map[string]uint64)

	if parent != nil {
		for key, val := range parent.Things {
			things[key] = val
		}
	}

	return &Scope{Things: things}
}

func (s *Scope) Insert(name string, oid uint64) {
	s.Lock()

	s.Things[name] = oid

	s.Unlock()
}

func (s *Scope) Lookup(name string) (oid uint64, comms chan Message) {
	s.Lock()

	oid = s.Things[name]
	comms = RT.Heap[oid]

	s.Unlock()

	return
}
