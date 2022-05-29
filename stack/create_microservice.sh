#!/bin/bash



echo "apiVersion: apps/v1
kind: Deployment
metadata:
  name: $1
  labels:
    app: $1
spec:
  replicas: 2
  selector:
    matchLabels:
      app: $1
  template:
    metadata:
      labels:
        app: $1
    spec:
      containers:
      - name: $1
        image: $1:0.0.1
        ports:
        - containerPort: 8080
        readinessProbe:
          periodSeconds: 5
          httpGet:
            path: /health
            port: 8080
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: $1
spec:
  selector:
    app: $1
  ports:
  - name: \"http-api\"
    port: 8080
" > $1.yaml
