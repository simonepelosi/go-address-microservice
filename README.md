# go-address-microservice

## API Documentation

### 1.1 GET CLOSEST ADDRESSES
```bash
curl --location --request POST 'http://localhost:8000/check-closest' \
--header 'Content-Type: application/json' \
--data-raw '{ "addresses": [ "Via di S. Basilio, 15, 00187 Roma RM, Italy", "Via Forestella, 13, 03026 Pofi FR, Italy", "Address C", "Address D"]}'
```
**RESPONSE**
```bash
{
    "AddressOne": {
        "Name": "Via di S. Basilio, 15, 00187 Roma RM, Italy",
        "Lat": 41.9053319,
        "Lng": 12.489857
    },
    "AddressTwo": {
        "Name": "Via Forestella, 13, 03026 Pofi FR, Italy",
        "Lat": 41.5532695,
        "Lng": 13.4022403
    },
    "DistanceInKm": 85
}
```

## RUN TESTS
Move into handlers folder and type:
```bash
go test
```