package cmd

// Version - server health check
func Version() string {
	url := serverAddr + "/version"
	s := PrettyHTTPResp(netClient.Get(url))

	return s
}
