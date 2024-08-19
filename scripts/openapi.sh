#!/bin/sh
oapi-codegen -generate types -o "./internal/openapi_types.gen.go" -package "api" "./api/api.yml"
