#!/bin/bash -e

# shutdown minecraft
screen -r minecraft -X stuff 'stop\n'
sleep 30
