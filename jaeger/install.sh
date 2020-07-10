#!/bin/sh

helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo update
helm install jaegertracing jaegertracing/jaeger --set query.ingress.enabled=true,query.ingress.annotations.'kubernetes\.io/ingress\.class'=traefik,query.ingress.'hosts[0]'=jaeger.hkdev
