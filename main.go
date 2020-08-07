package main

import (
	"github.com/thaitanloi365/go-monitor/cmd"
)

func main() {
	cmd.Execute()
	// os.Setenv("DOCKER_HOST", "unix:///var/run/docker.sock")
	// cli, err := client.NewEnvClient()
	// if err != nil {
	// 	panic(err)
	// }

	// list, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
	// 	All: true,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// for _, item := range list {
	// 	fmt.Println(item)

	// }

	// reader, err := cli.ContainerLogs(context.Background(), "ezielog-backend", types.ContainerLogsOptions{
	// 	ShowStdout: true,
	// 	ShowStderr: true,
	// 	Follow:     true,
	// 	Timestamps: false,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// scanner := bufio.NewScanner(reader)
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// 	c.Response().WriteHeader(http.StatusOK)

	// 	// enc := json.NewDecoder(c.Response())
	// 	for scanner.Scan() {
	// 		c.Response().Write([]byte(scanner.Text()))
	// 		// var s = map[string]interface{}{
	// 		// 	"sss": scanner.Text(),
	// 		// }

	// 		// if err := enc.Encode(s); err != nil {
	// 		// 	fmt.Println(err)
	// 		// }
	// 		c.Response().Flush()
	// 	}

	// 	return nil
	// })
	// e.Logger.Fatal(e.Start(":1323"))

}
