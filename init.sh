#!/bin/bash
rm -r ~/.checklistcli
rm -r ~/.checklistd

checklistd init mynode --chain-id checklist

checklistcli config keyring-backend test

checklistcli keys add me
checklistcli keys add you

checklistd add-genesis-account $(checklistcli keys show me -a) 1000foo,100000000stake
checklistd add-genesis-account $(checklistcli keys show you -a) 1foo

checklistcli config chain-id checklist
checklistcli config output json
checklistcli config indent true
checklistcli config trust-node true

checklistd gentx --name me --keyring-backend test
checklistd collect-gentxs