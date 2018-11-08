/*nav.js handles all the javascript for the page navigation pane*/ 


// dynamically loads the navigation on page refresh
document.addEventListener("load", requestNav())
function requestNav() {
    let xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {

      navtext = document.querySelector("nav");
    
      let arr = JSON.parse(this.response)
      for (i in arr) {
        navtext.innerHTML += '<div class="navlink">' + arr[i].Name + '</div>';
      }

    }
  };
  xhttp.open("GET", "nav", true);
  xhttp.send();
}