/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package localnetv1

import (
	"net"
	"sort"
)

func NewIPSet(ips []string) (set *IPSet) {
	set = &IPSet{}
	set.AddAll(ips)
	return
}

// Add adds an address to this set, returning the parsed IP. `ìp` will be nil if it couldn't be parsed.
func (set *IPSet) Add(s string) (ip net.IP) {
	ip = net.ParseIP(s)
	if ip == nil {
		return
	}

	if len(ip) == net.IPv6len {
		insertString(&set.V6, s)
	} else {
		insertString(&set.V4, s)
	}

	return
}

func (set *IPSet) AddAll(ips []string) {
	for _, ip := range ips {
		set.Add(ip)
	}
}

func (set *IPSet) AddSet(set2 *IPSet) {
	if set2 == nil {
		return
	}

	for _, ip := range set2.V4 {
		insertString(&set.V4, ip)
	}
	for _, ip := range set2.V6 {
		insertString(&set.V6, ip)
	}
}

func insertString(a *[]string, s string) {
	idx := sort.SearchStrings(*a, s)

	if idx != len(*a) && (*a)[idx] == s {
		// already there
		return
	}

	// insert
	(*a) = append(*a, "")
	copy((*a)[idx+1:], (*a)[idx:])
	(*a)[idx] = s
	return
}
