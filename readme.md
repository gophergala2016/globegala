## Globegala

Shows all the gophergala2016 hackathon contributors in a world map 🌎. 

Globegala uses github API to get the data. You can view it at http://globegala.herokuapp.com/

#### Preview

<img width="854" alt="screen shot 2016-01-24 at 4 25 30 pm" src="https://cloud.githubusercontent.com/assets/4488777/12540106/84bbdaf6-c2b7-11e5-8d5d-3c8b6a8b1d5b.png">

#### Running in Local

1. `go get github.com/gophergala2016/globegala`  
2. Set enviroment variable `EXPORT access_token=$YOUR_GITHUB_TOKEN (This is optional)
3. To run the server `cd $GOPATH/src/github.com/gophergala2016/globegala` then `./run.sh`  
 
#### Notes

Github doesn't provide lat, lng. Instead it provides city/country name. Globegala placed people according to there city/country centers. 
Some of the contributors do not have there location set in Github. We placed them in our `Gopher Treasure Hunt` event.
We don't use Github Real-time API. Your changes won't effect that instance also response is cached.

<img width="334" alt="screen shot 2016-01-24 at 4 26 20 pm" src="https://cloud.githubusercontent.com/assets/4488777/12540104/82f35230-c2b7-11e5-8f4b-5abc4c562b36.png">

