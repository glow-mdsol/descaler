# Descaler  
A URL unpacking tool, defaults to outputting in JSON

## Installation
`go install`

## Usage
```
$ ./descaler -url https://mcc-innovate.imedidata.com/apps/dalton/client_division_schemes/48d1c900-.../configuration_types/85f488d4-.../configuration_type_roles\?c_selections\=com%3Amdsol%3Aconfiguration_type_roles%3A1ea3ee78...%2Ccom%3Amdsol%3Aclient_divisions%3A48d1c900-...%2Ccom%3Amdsol%3Aactions%3Aconfigure_roles\&client_division_uuid\=48d1c900... | python -m json.tool
  {
      "base_url": "mcc-innovate.imedidata.com",
      "fragment": "",
      "path": "/apps/dalton/client_division_schemes/48d1c900-.../configuration_types/85f488d4-.../configuration_type_roles",
      "queries": {
          "c_selections": [
              "com:mdsol:configuration_type_roles:1ea3ee78-...",
              "com:mdsol:client_divisions:48d1c900-...",
              "com:mdsol:actions:configure_roles"
          ],
          "client_division_uuid": [
              "48d1c900-..."
          ]
      }
  }
```

## Improvements
* maybe add a query statement