package singleton

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
)

var lock = &sync.Mutex{}

type single struct {
	ctx    context.Context
	config config.Config
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			// Using the SDK's default configuration, loading additional config
			// and credentials values from the environment variables, shared
			// credentials, and shared configuration files
			cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}
			singleInstance = &single{
				ctx:    context.TODO(),
				config: cfg,
			}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
