# Sunset

Retrieve crepscular timings from sunrise-sunset.org

## Design

This package is designed to be able to automatically retrieve Sunset timings based upon your current GeoIP. Each type and method is separately exported, so you can insert your own logic anywhere in the chain. For an example of how the chain ties together, see `sunset/sunset.go`

For an explanation of the different times returned, please see [MrReid.org's explanations](http://wordpress.mrreid.org/2013/02/05/dawn-dusk-sunrise-sunset-and-twilight/)

## APIs

`freegeoip.go`: GeoIP infromation from [freegeoip.net](https://freegeoip.net/). Works around IPv6 limitations.

`ipify.go`: Your current IPv4 from [ipify.org](https://api.ipify.org?format=json). Used to work around FreeGeoIP's IPv6 issues.

`sunrise-sunset.go`: Sunrise timings from [sunrise-sunset.org](https://api.sunrise-sunset.org/json?lat=36.7201600&lng=-4.4203400&formatted=0). Times are in UTC and must be converted with `.Local()`.


## Use

Go Get: `go get -u github.com/bmhatfield/sunset`
CLI: `go run print-sunset.go`
API: `sunset.Time()`
