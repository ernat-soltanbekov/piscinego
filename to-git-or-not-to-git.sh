#!/bin/bash
curl -sL https://01.tomorrow-school.ai/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'curl -sL https://01.tomorrow-school.ai/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
