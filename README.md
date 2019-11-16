# k8s-test

This repository has a simple test steps to learn kubernetes.
Before start you should have a installed `kubernetes` and `minikube` on your machine.

## Prepare
Gets repository clone
```bash
git clone git@github.com:ubozov/k8s-test.git
cd k8s-test/
```

Starts minikube and installs kubernetes dashboard
```bash
 minikube start
 kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
 ```

Creates kubernetes user (you should copying user token)
```bash
kubectl apply -f ./kube/dashboard-adminuser.yaml
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')
```

Starts dashboard
```bash
kubectl proxy
```
`kubectl` will make dashboard available at http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/ . You should put token to login page.

## Deployment
Creates test pods and launches service via NodePort (port: 32001)
```bash
kubectl apply -f ./kube/app/deployment.yaml
kubectl apply -f ./kube/app/service.yaml
```

## Test
Below command returns the IP address of local minikube cluster, you should pass on browser this IP and port 32001 for check activity
```bash
minikube ip
```


