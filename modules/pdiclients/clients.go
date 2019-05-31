package pdiclients

import pdiutil "github.com/Soontao/pdi-util"

// PDIClients type
//
// central manage all tenant client
type PDIClients struct {
	clients map[string]*PDIClientInfo
}

// PDIClientInfo type
type PDIClientInfo struct {
	username string
	password string
	hostname string
	client   *pdiutil.PDIClient
}

// NewPDIClients instance
func NewPDIClients() *PDIClients {
	return &PDIClients{
		clients: map[string]*PDIClientInfo{},
	}
}

// GetPDIClient ref
func (info *PDIClientInfo) GetPDIClient() *pdiutil.PDIClient {
	return info.client
}

// GetOrAddNewClient to collection
//
// Please concern about lock
func (cs *PDIClients) GetOrAddNewClient(host, username, password string) (client *pdiutil.PDIClient, err error) {

	if info, ok := cs.clients[host]; ok {
		// password changed
		if info.username != username || info.password != password {
			info.username = username
			info.password = password
			info.hostname = host
			info.client, err = pdiutil.NewPDIClient(username, password, host, "")
		}
		client = info.client
	} else {
		// create new mapping
		info = &PDIClientInfo{username, password, host, nil}
		// pdi will auth detect the release version
		if info.client, err = pdiutil.NewPDIClient(username, password, host, ""); err == nil {
			cs.clients[host] = info
			client = info.client
		}
	}
	return client, err
}

// Remove tenant from collection
func (cs *PDIClients) Remove(host string) (rt error) {
	// if exist
	if _, ok := cs.clients[host]; ok {
		delete(cs.clients, host)
	}
	return rt
}
