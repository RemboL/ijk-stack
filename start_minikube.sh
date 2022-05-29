#!/bin/bash

minikube start --extra-config=kubelet.housekeeping-interval=10s
