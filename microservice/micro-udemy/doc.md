1. kubectl delete svc broker-service
2. kubectl expose deployment broker-service --type=LoadBalancer --port=80
   80 --target-port=8080
3. minikube tunnel
4. minikube addons enable ingress
5. kubectl apply -f ingress.yml
6. kubectl get ingress
7. sudo vi /etc/hosts
8. minikube tunnel
