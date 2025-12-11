#!/bin/bash
if [ -z "HERO_ID" ]; then
    echo "Error: HERO_ID environment variable is not set."
    exit 1
fi
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq " .[] | select(.id == $HERO_ID) | (.connections.relatives)" | tr -d "\""