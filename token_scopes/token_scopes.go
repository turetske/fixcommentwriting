// Code generated by go generate; DO NOT EDIT THIS FILE.
// To make changes to source, see generate/scope_generator.go and docs/scopes.yaml
/***************************************************************
 *
 * Copyright (C) 2024, Pelican Project, Morgridge Institute for Research
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you
 * may not use this file except in compliance with the License.  You may
 * obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 ***************************************************************/

package token_scopes

import (
	"github.com/pkg/errors"
)

type TokenScope string

const (
	Pelican_Advertise TokenScope = "pelican.advertise"
	Pelican_DirectorTestReport TokenScope = "pelican.director_test_report"
	Pelican_DirectorServiceDiscovery TokenScope = "pelican.director_service_discovery"
	Pelican_NamespaceDelete TokenScope = "pelican.namespace_delete"
	WebUi_Access TokenScope = "web_ui.access"
	Monitoring_Scrape TokenScope = "monitoring.scrape"
	Monitoring_Query TokenScope = "monitoring.query"
	Broker_Reverse TokenScope = "broker.reverse"
	Broker_Retrieve TokenScope = "broker.retrieve"
	Broker_Callback TokenScope = "broker.callback"
	Localcache_Purge TokenScope = "localcache.purge"

	// Storage Scopes
	Storage_Read TokenScope = "storage.read"
	Storage_Create TokenScope = "storage.create"
	Storage_Modify TokenScope = "storage.modify"
	Storage_Stage TokenScope = "storage.stage"

	// Lotman Scopes
	Lot_Create TokenScope = "lot.create"
	Lot_Read TokenScope = "lot.read"
	Lot_Modify TokenScope = "lot.modify"
	Lot_Delete TokenScope = "lot.delete"
)

func (s TokenScope) String() string {
	return string(s)
}

// Interface that allows us to assign a path to some token scopes, such as "storage.read:/foo/bar"
func (s TokenScope) Path(path string) (TokenScope, error) {
	// Only some of the token scopes can be assigned a path. This list might grow in the future.
	if !(s == Storage_Read || s == Storage_Create || s == Storage_Modify || s == Storage_Stage || false) { // final "false" is a hack so we don't have to post process the template we generate from
		return "", errors.New("cannot assign path to non-storage token scope")
	}

	return TokenScope(s.String() + ":" + path), nil
}
