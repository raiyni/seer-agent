package main

import (
	"net/http"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

type Node struct {
	Image  string            `json:"Image"`
	Names  []string          `json:"Names"`
	ID     string            `json:"Id"`
	Labels map[string]string `json:"Labels"`
}

func getImages() []Node {
	ns := make([]Node, 0)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	images, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		if err != nil {
			panic(err)
		}

		n := Node{
			Image:  image.Image,
			Names:  image.Names,
			ID:     image.ID,
			Labels: image.Labels,
		}

		ns = append(ns, n)
	}

	return ns
}

func imageInfo(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, getImages(), "  ")
}

func main() {

	e := echo.New()
	e.GET("/", imageInfo)
	e.Logger.Fatal(e.Start(":3333"))
}
