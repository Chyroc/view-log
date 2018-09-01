package command

func init() {
	c := Command{Client: lsClient, Server: lsServer, Min: 0, Max: 0}
	registerCommand("ls", c)
	registerCommand("ll", c)
}

func lsClient(args []string) (string, error) {
	return "client", nil
}

func lsServer(args []string) (string, error) {
	return "server", nil
}
