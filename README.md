# Summary
Searchs hyundai vehicle inventory

This app will search hyundai inventory and send text messages when found something meeting criteria.
It is initially set to run once a minute, you can edit this though in the main method
You can specify the criteria you want in the meetsCondition method

# Before Running
You can get an updated curl response with the following page: https://mholt.github.io/curl-to-go/
You can get a good curl command by going to https://www.hyundaiusa.com/us/en/inventory-search/vehicles-list?model=Palisade&year=2022&trims=Caligraphy and searching
for the GET at domain:hyundaiusa.com file:vehicle-list.json?....
Making use of twilo to send notifications: https://www.twilio.com/blog/send-sms-30-seconds-golang

You need to set your TWILIO account secrets as well as phone numbers

```
export TWILO_ACCOUNT_SID=<ACCOUNT_SID>
export TWILO_AUTH_TOKEN=<AUTH_TOKEN>
export VEHICLE_SEARCH_PHONE_TO_NUMNBER=<Phone number to send text to, something like +15557771234)
export VEHICLE_SEARCH_PHONE_FROM_NUMNBER=<Phone number to send text from, i was using a free twilo phone i created))
```

# Building
go build

# Running
go run vehicle-search

# Expected Results
Terminal will output something similiar to the following, this same string will be the thing pushed over SMS

``` 
https://www.hyundaiusa.com/us/en/inventory-search/vehicles-list?model=Palisade&year=2022
Name: Jack Giambalvo Hyundai
Distance: 85.629997
Phone: 717-755-2944
{
    "Vin": "KM8R7DHE5NU485404",
    "ModelNm": "Palisade",
    "TrimDesc": "CALLIGRAPHY",
    "ModelCd": "J1472A65",
    "Price": "$51,040.00",
    "ExteriorColorCd": "P7V",
    "InteriorColorCd": "WDN",
    "DrivetrainDesc": "ALL WHEEL DRIVE",
    "TransmissionDesc": "AUTO",
    "TotalPackages": 0,
    "TotalOptions": 0,
    "Packages": null,
    "Colors": [
        {
            "SAPExterioColorCode": "P7V",
            "ExtColorLongDesc": "STEEL GRAPHITE"
        }
    ],
    "PlannedDeliveryDate": "2022-06-23T04:45:00",
    "InventoryStatus": "AA"
}
```
