package main

import (
	"math/rand"
	"time"

	"github.com/rebirthmonkey/k8s-dev/pkg/kubecontroller"
	_ "github.com/rebirthmonkey/k8s-dev/scaffold/kubecontroller/pkg/apiexts/all"
	_ "github.com/rebirthmonkey/k8s-dev/scaffold/kubecontroller/pkg/reconcilers/all"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	kubecontroller.NewApp("kubecontroller").Run()
}
