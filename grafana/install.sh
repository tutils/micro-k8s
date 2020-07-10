#!/bin/sh

helm install grafana stable/grafana --set adminPassword=admin,ingress.enabled=true,ingress.annotations.'kubernetes\.io/ingress\.class'=traefik,ingress.'hosts[0]'=grafana.hkdev
