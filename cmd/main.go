package main

import (
	"bitgo_client/pkg"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)


func main() {
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	parts := strings.Split(cwd, "/")
	p := parts[:len(parts) - 1]
	exPath := strings.Join(p, "/") + "/json_samples/"
	walletPath := exPath + "wallet.json"
	eprisePath := exPath + "enterprise.json"

	var token string
	flag.StringVar(&token, "token", "", "BitGo's access token")
	isProd := flag.Bool("isProd", false, "BitGo's API type, default: test")
	flag.Parse()

	// Get user info
	c := pkg.NewBitGoClient(token, *isProd)
	res, err := c.GetUserBy("psauxgrepkill@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(res)
	fmt.Println("******")

	// Create new wallet address
	payload, err := ioutil.ReadFile(walletPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	res, err = c.NewWalletAddress("tltc", "5c8fbbaa51cd44c9031f0695f8f9877a", string(payload))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(res)
	fmt.Println("******")

	// Create enterprise
	payload, err = ioutil.ReadFile(eprisePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	res, err = c.CreateEnterprise(string(payload))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(res)
}

