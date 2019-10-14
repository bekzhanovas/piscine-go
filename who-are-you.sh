#! /bin/bash
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq '.[51].name'
