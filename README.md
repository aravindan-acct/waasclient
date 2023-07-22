# Barracuda WAF-As-A-Service API Client in Golang

### What is this ?

This package can be used to make API calls to Barracuda WAF-As-A-Service or a Barracuda WAF Instance.

### How to ?

An implementation of the package is included in the waas_api_test directory. 

For WAAS: 

- Set the environment variables for the WAF-As-A-Service credentials:
    1. `export WAAS_EMAIL=<waas account email address>`

    2. `export WAAS_PASSWD=<waas account password>`

For WAF:

- Set the environment variables for the WAF-As-A-Service credentials:
    1. `export WAF_USERNAME=<waf username>`

    2. `export WAF_PASSWD=<waf password>`

- Navigate to the directory and run `go run .`
    
    `cd waas_api_test`

    `go run . waf` or `go run . waas` 

