#!/bin/sh
read -p "Please input FORTISWITCH_ACCESS_HOSTNAME:" host
read -p "Please input FORTISWITCH_ACCESS_USERNAME:" username
read -p "Please input FORTISWITCH_ACCESS_PASSWORD:" password
read -p "Please input FORTISWITCH_INSECURE:" insecure
read -p "Please input FORTISWITCH_CA_CABUNDLE:" cabundlefile

p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
export "FORTISWITCH_ACCESS_HOSTNAME"=$host
export "FORTISWITCH_ACCESS_USERNAME"=$username
export "FORTISWITCH_ACCESS_PASSWORD"=$password
export "FORTISWITCH_INSECURE"=$insecure
export "FORTISWITCH_CA_CABUNDLE"=$cabundlefile

echo $FORTISWITCH_ACCESS_HOSTNAME
echo $FORTISWITCH_ACCESS_USERNAME
echo $FORTISWITCH_ACCESS_PASSWORD
echo $FORTISWITCH_INSECURE
echo $FORTISWITCH_CA_CABUNDLE

make -C ../  testacc
