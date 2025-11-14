#! /bin/bash

INTERVIEW=$(grep -h "SEE INTERVIEW" streets/* | grep -oE '[0-9]+')
echo "$INTERVIEW"
cat interviews/interview-"$INTERVIEW"
echo "$MAIN_SUSPECT"
echo ""
