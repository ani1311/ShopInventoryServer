let URL = "http://localhost:8080"
let url = document.location.href
let params = url.split('?')
let paramData = {}
let markers = []

for (var i = 0, l = params.length; i < l; i++) {
    let tmp = params[i].split('=');
    paramData[tmp[0]] = tmp[1];
}

if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(getShopData);
} else {
    console.log("FUCK YOU")
}


function getShopData(pos) {
    map, infoWindow = initMap(pos)
    let request = new XMLHttpRequest()
    if (paramData == undefined) return
    request.open("GET", URL + "/shopWithItem" + "?barcode=" + paramData['barcode'] + "&longitude=" + pos["coords"]["longitude"] + "&latitude=" + pos["coords"]["latitude"])
    request.send()
    request.onload = () => {
        data = JSON.parse(request.response)
        data["data"].forEach((shop, index) => {
            createMarker(map, shop["latitude"], shop["longitude"], shop["name"], infoWindow)
        })
    }
    markers.forEach((marker,index) => {
        google.maps.event.addListener(marker, 'click', function() {
            infowindow = new google.maps.InfoWindow({
                content: marker.title
            });
            infowindow.open(map, marker);
        });
    })
}


function loadMapWithShops(data) {
    data.forEach()
}

function initMap(pos) {
    if (pos == undefined) return
    map = new google.maps.Map(document.getElementById('map'), {
        center: { lat: pos["coords"]["latitude"], lng: pos["coords"]["longitude"] },   //initial map center
        zoom: 17,    //for street level zooming
        mapTypeId: google.maps.MapTypeId.ROADMAP
    });
    createMarker(map, pos["coords"]["latitude"], pos["coords"]["longitude"], "ME")
    infoWindow = new google.maps.InfoWindow;
    return map, infoWindow
}

function createMarker(map, lat, lng, name, infoWindow) {
    // var marker = new google.maps.Marker({
    // position: latLng,
    // map: map,
    // title: data.title
    // });
    markers.push(new google.maps.Marker({
        position: new google.maps.LatLng(lat, lng),
        map: map,
        title : name
    }))
}

function handleLocationError(browserHasGeolocation, infoWindow, pos) {
    infoWindow.setPosition(pos);
    infoWindow.setContent(browserHasGeolocation ?
        'Error: The Geolocation service failed.' :
        'Error: Your browser doesn\'t support geolocation.');
    infoWindow.open(map);
}