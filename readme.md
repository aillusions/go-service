
https://www.youtube.com/watch?v=TYx0BTyFtmc

docker build -t aillusions/go-service:0.0.1 .

kubectl apply -f pod.yaml
kubectl logs -f hello-go-pod

docker build -t aillusions/go-service:0.0.2 .
kubectl apply -f pod.yaml


brew install skaffold
skaffold dev

skaffold delete

while true; do cont=$(curl -s http://localhost:30001); echo "$var: $cont"; var=$((var+1)); sleep 0; done