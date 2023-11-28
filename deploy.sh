#!/bin/bash
echo "Deploying K8S cluster..."
echo

echo "Deploying postgresql..."
for file in ./k8s/postgresql/*
do
  kubectl apply -f $file
done

echo
echo "Waiting for postgresql to start..."
kubectl wait --for=condition=ready --timeout=-1s pod -l app=db 

pods=$(kubectl get pod -l=app=db)
regex="db-([^\s]*)"
dbPod=$(echo "$pods" | grep -oP "$regex")

echo $dbPod
kubectl exec --stdin --tty $dbPod -- /bin/bash -c "psql -U postgres -d postgres -c 'CREATE DATABASE keycloak'"


echo "Deploying postgresql..."
for file in ./k8s/pgadmin/*
do
  kubectl apply -f $file
done
