let URL = "http://localhost:8080"


function addToItems(name,barcode){
    let listOfItems = document.getElementById("items")
    let newItem = document.createElement('li');
    newItem.className = "list-group-item"
    newItem.id = barcode
    newItem.appendChild(document.createTextNode(name));
    itemButton = document.createElement('BUTTON');
    itemButton.appendChild(document.createTextNode("GET"));
    itemButton.addEventListener('click',() =>{
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition((position) => {
                window.location.href = "/shopsWithItem.html?barcode=" + barcode
            });
        } else {
            console.log("error")
          }
    })
    newItem.appendChild(itemButton);
    listOfItems.appendChild(newItem);
}


function clearItemList(){
    itemList = document.getElementById("items")
    while( itemList.firstChild ){
        itemList.removeChild( itemList.firstChild );
      }
}

function getItems(itemName){
    let request = new XMLHttpRequest()
    request.open("GET",URL + "/getItems" + "?name=" + itemName)
    request.send()
    request.onload = () => {
        data = JSON.parse(request.response)
        console.log(data)
        data["data"].forEach((item,index) => {
            addToItems(item.name,item.barcode)
        })
    }
}



let getItemButton = document.getElementById("getItemButton")
getItemButton.addEventListener('click',function() {
    clearItemList()
    itemNaxtText = document.getElementById("itemNameText")
    getItems(itemNaxtText.value)
    itemNaxtText.value = ''
})
