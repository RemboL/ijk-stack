#!/bin/bash

~/istio-1.11.4/bin/istioctl install -y

kubectl label namespace default istio-injection=enabled

kubectl rollout restart `kubectl get deploy -o name`
