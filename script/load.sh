#!/bin/bash

echo "loading data ..."
curl localhost:8080/mutate -H "X-Dgraph-CommitNow: true" -XPOST -d "$(cat data/links-and-topics.nq)" | jq '.'

echo "loading schema ..."
curl localhost:8080/alter -XPOST -d $'
type: string @index(exact) .
externalId: string @index(exact) .
url: string @index(term) .
name: string @index(term) .
title: string @index(term) .
includes: uid @reverse .
'
