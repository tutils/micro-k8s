#!/bin/sh

helm install stable/prometheus --set server.ingress.enabled=true,server.ingress.annotations.'kubernetes\.io/ingress\.class'=traefik,server.ingress.'hosts[0]'=prometheus.hkdev
