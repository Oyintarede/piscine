#! /bin/bash

KEY_INTERVIEW=$(grep -R "L337" -n interviews/ \
    | head -1 \
    | sed -E 's/.*interview[-]?([0-9]+).*/\1/')
export KEY_INTERVIEW
echo "$KEY_INTERVIEW"
cat "interviews/interview-$KEY_INTERVIEW"
echo "$MAIN_SUSPECT"