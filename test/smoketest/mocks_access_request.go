package smoketest

var accessRequest = `{
  "apiVersion":"dsr/v1",
  "kind":"AccessRequest",
  "metadata": {
    "uid":"D6EFAECD-2FCB-486D-B49A-4F699BFC24D1",
    "tenant":"axonic"
  },
  "request": {
    "property": "example.com",
    "environment": "production",
    "jurisdiction": "CCPA",
    "regulation": "CCPA",
    "subject": {
      "firstName": "JOE",
      "lastName": "SMITH",
      "phone": "1234567890",
      "email": "joe@example.com",
      "addressLine1": "123 MAIN ST",
      "addressLine2": "APT 34",
      "city": "ENGLEWOOD",
      "stateRegionCode": "MA",
      "postalCode": "10123-1234"
    },
    "identities": [
      {
        "identitySpace": "axonicID",
        "identityValue": "00015c60-37c5-11e9-83e4-0e7679b64802"
      }
    ],
    "callbacks": [
      {
        "url": "https://example.com/callback",
        "headers": {
          "Authorization": "Bearer $auth"
        }
      }
    ],
    "claims": {
    }
  }
}`
