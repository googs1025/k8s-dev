package reconcilermgr

import (
	"os"

	"github.com/rebirthmonkey/go/pkg/log"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	crmgr "sigs.k8s.io/controller-runtime/pkg/manager"
)

type ReconcilerManager struct {
	Concurrence    int
	Portable       bool
	APIServerURL   string
	APIExtsEnabled bool
	APIExtsURL     string
	APIToken       string
}

type PreparedReconcilerManager struct {
	crmgr.Manager
	client.Client
	//Scheme *runtime.Scheme

	*ReconcilerManager
}

func (rmgr *ReconcilerManager) PrepareRun(scheme *runtime.Scheme) *PreparedReconcilerManager {
	log.Info("[ReconcilerManager] PrepareRun")

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     ":8001",
		Port:                   9443,
		HealthProbeBindAddress: ":8002",
		LeaderElection:         false,
		LeaderElectionID:       "465bd3f6.wukong.com",
	})
	if err != nil {
		log.Error("unable to start manager")
		os.Exit(1)
	}

	return &PreparedReconcilerManager{
		Manager: mgr,
		//Scheme:            scheme,
		Client:            mgr.GetClient(),
		ReconcilerManager: rmgr,
	}
}

func (rmgr *PreparedReconcilerManager) Run() error {
	log.Info("[PreparedReconcilerManager] Run")

	if err := rmgr.Manager.Start(ctrl.SetupSignalHandler()); err != nil {
		log.Error("problem running manager")
		os.Exit(1)
	}

	return nil
}
