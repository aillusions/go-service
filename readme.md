
docker build -t go-service:0.0.1 .

kubectl apply -f pod.yaml

kubectl logs -f hello-go-pod