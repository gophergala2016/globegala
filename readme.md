## Globegala

A world map that tracks where all the gophergala2016 submissions are coming from.

You can view it at http://globegala.herokuapp.com/

#### Running in Local

1. `go get github.com/gophergala2016/globegala`  
2. Set enviroment variable `EXPORT access_token=$YOUR_GITHUB_TOKEN  
3. To run the server `cd $GOPATH/github.com/gophergala2016/globegala` then `./run.sh`  
 
#### Notes

Some of the contributors do not have there location set in Github. We placed them in our `Gopher Treasure Hunt` event.
Github doesn't provide lat, lng. Instead it provides city/country name. Globegala placed people according to there city/country centers. 
We don't use Github Real-time API. Your changes won't effect that instance also response is cached.
