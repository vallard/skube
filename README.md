# Skube
A small kubernetes client to include in simple Kuberentes utilities. 

## Usage

```golang

package main

import (
	"fmt"
	"ioutil"
	
	"github.com/vallard/skube"
)

func main() {
	// kubernetes token
	token := "asdfasdfasdf"
	// read in the kubernetes certificate authority
	ca, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		panic()
	}
	// create a new skube client
	k := skube.New("https://mykubernetes:6443", token, ca)
	
	// do skube stuff with it. 
	deployments, err := k.ListDeployments("")
	
	if err ! = nil {
		panic()
	}
	
	for _, d := range deployments {
		fmt.Println(d)
	}