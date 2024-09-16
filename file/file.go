package file

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/tools/clientcmd"
)

func UpdateKubeconfig(rancherURL string, newToken string) error {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("Failed to find home directory: %v", err)
		}
		kubeconfigPath = home + "/.kube/config"
	}

	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	for name, cluster := range config.Contexts {
		if cluster.Cluster == rancherURL {
			config.AuthInfos[name].Token = newToken
		}
	}

	return clientcmd.WriteToFile(*config, kubeconfigPath)
}
