#!/usr/bin/env bash

rm -rf ../product_research/*

for env in test development staging production; do
  spiff merge product_research/marketplaces_$env.yml product_research/mail_packages.yml output/categories_$env.yml > ../product_research/$env.yml;
done
