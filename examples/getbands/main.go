package main

import (
	"fmt"
	"github.com/Beta5051/naverbandgo"
)

func main() {
	client := naverbandgo.NewClient("발급 받은 엑세스 토큰", nil)


	bands, err := client.GetBands()
	if err != nil {
		panic(err)
	}

	for _, band := range bands {
		fmt.Printf("%s - %s\n", band.Name, band.BandKey)
	}
}