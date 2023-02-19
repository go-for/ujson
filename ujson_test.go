package ujson

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	str := `{
        "Image": {
            "Width":  800,
            "Height": 600,
            "Title":  "View from 15th Floor",
            "Thumbnail": {
                "Url":    "http://www.example.com/image/481989943",
                "Height": 125,
                "Width":  100
            },
            "Animated" : false,
            "IDs": [116, 943, 234, 38793],
            "GeoInfo": {
                "Latitude":  37.7668,
                "Longitude": -122.3959
            }
        }
}`
	result, err := Unmarshal([]byte(str))
	fmt.Printf("%#v\n", result)
	fmt.Printf("%v\n", err)
}
