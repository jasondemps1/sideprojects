#!/bin/sh

Replace() {
    sed -i "s/__$1__/$2/g" $3
    echo "Replaced: $1 => $2"
}


Replace "LISTENER_ADDRESS" $LISTENER_ADDRESS $1
Replace "LISTENER_PORT" $LISTENER_PORT $1
Replace "CLUSTER_ADDRESS" $CLUSTER_ADDRESS $1
Replace "CLUSTER_PORT" $CLUSTER_PORT $1

cat $1

/usr/local/bin/envoy -c /etc/envoy/envoy.yaml 2>&1