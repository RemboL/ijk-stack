#!/bin/bash

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.13/samples/addons/jaeger.yaml


~/istio-1.11.4/bin/istioctl install --set meshConfig.defaultConfig.tracing.zipkin.address=jaeger-collector.istio-system:9411 --set meshConfig.defaultConfig.tracing.sampling=100.0 -y

kubectl rollout restart `kubectl get deploy -o name`
