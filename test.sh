#!/bin/bash

URL="http://localhost:8080/submit-data"


http POST $URL \
Name="Rabi" LastName="Muhammad Siddique" \
Content-Type="application/json"
