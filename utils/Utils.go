package utils

import (
	"github.com/donething/utils-go/dohttp"
	"time"
)

var Client = dohttp.New(30*time.Second, false, false)
