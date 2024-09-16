package main

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/stenstromen/rancher-renewer/api"
	"github.com/stenstromen/rancher-renewer/file"
)

func main() {
	apiKey := os.Getenv("RANCHER_API_KEY")
	rancherURL := os.Getenv("RANCHER_URL")

	if apiKey == "" || rancherURL == "" {
		log.Fatal("RANCHER_API_KEY and RANCHER_URL must be set as environment variables")
	}

	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Failed to find home directory: %v", err)
	}

	kubeconfigPath := home + "/.kube/config"
	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	for name, cluster := range config.Contexts {
		if cluster.Cluster == rancherURL {
			authInfo := config.AuthInfos[name]
			token := authInfo.Token

			tokenInfo, err := api.GetRancherTokenInfo(apiKey, rancherURL, token)
			if err != nil {
				log.Printf("Failed to check token for context %s: %v", name, err)
				continue
			}

			if api.TokenIsExpiringSoon(tokenInfo.ExpiresAt) {
				log.Printf("Token for context %s is expiring soon, renewing...", name)

				newToken := tokenInfo.Token
				err = file.UpdateKubeconfig(rancherURL, newToken)
				if err != nil {
					log.Printf("Failed to update kubeconfig for context %s: %v", name, err)
				} else {
					log.Printf("Token for context %s successfully renewed and updated", name)
				}
			} else {
				log.Printf("Token for context %s is still valid", name)
			}
		}
	}
}
