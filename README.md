# rancher-renewer

Utility to renew Rancher tokens in local kubeconfig.

Renew token if it expires in less than 7 days.

## Environment variables

```bash
export RANCHER_URL="https://rancher.example.com"
export RANCHER_API_KEY="your-api-key"
export KUBECONFIG="your-kubeconfig-file" # optional, default to ~/.kube/config
```
