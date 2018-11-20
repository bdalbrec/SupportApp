/* This script creates a simple search by scraping the page for all <a> tags and putting them into an array

Author: Brian Albrecht
Date: 1/25/15
Filename: searchscript.js
*/

// Global Variables

var searchList = [];  // making this global so it can easily be passed from createSearchArray to listElements



//This function gets all the elements with the <a> tag within the toggle-view <ul> element and puts them into an array.
function createSearchArray() {
    var getLinks = document.getElementById("toggle-view");
    searchList = getLinks.getElementsByTagName("a");
    // var linksList = document.links      ---------this also works
}



// This function loops through the searchList array and lists all the elements within the <a> tag that match what is entered in the searchBox
function listElements() {
    document.getElementById("results").innerHTML = ""; // clears out the results list to make room for new list upon iteration of the loop
    var searchTerm = document.getElementById("searchBox").value.toLowerCase(); //creates a search term from the contents of the search box and converts it to lower case
    searchResults.style.visibility = "hidden"; // this needs to be here to hide the box if the list is set back to blank. Input triggering the loop will make it visible
    if (searchBox.value != "") { // only do the work if there is a value in the searchbox
       // var resultsArray = [];
        //var resultsCounter = 0;
        // step through the array and see if searchTerm matches each entry. If it matches print out that array index value as a link in the results box.
        for (var i = 0; i < searchList.length; i++) {
            if (searchList[i].innerHTML.toLowerCase().indexOf(searchTerm) > -1 || searchList[i].title.toLowerCase().indexOf(searchTerm) > -1) { // compares the now lower case search term to lower case values in the search array
                var a = document.createElement('a'); // creates the link element
                var listItem = document.createElement('li'); //creates the list item
                a.text = (searchList[i]).innerHTML; // adds the text to the lin
                a.title = (searchList[i]).textContent; // title to link
                a.href = searchList[i]; // adds the address to the link
                a.target = "_blank";
                listItem.appendChild(a); // appends the link created to the <li> tag.
                var resultsBox = document.getElementById("results"); 
                resultsBox.appendChild(listItem); // inserts the <li> tag into the <ul> with ID results
                searchResults.style.visibility = "visible"; // makes the searchResults div visible if there is content
            }
        }    
    }
  
}


// This inserts placeholder text into the search box for crappy old browsers

function searchPlaceholder() {
    if (!Modernizr.input.placeholder) {
        document.getElementById("searchBox").value = "Search - IE9+";
    }
}


// Clear the search box and hide the search list if the user clicks off of the search box.

function cleanPage() {
    document.getElementById("searchBox").value = "";
    searchResults.style.visibility = "hidden";
    document.getElementById("searchBox").focus();
}


// Event Listeners to call functions that create the search array, write it to the output, and fill in placeholder text when the page loads
var inputButton = document.getElementById("searchBox");
var outOfBounds = document.getElementById("wrapper"); 
if (document.addEventListener) {
    window.addEventListener("load", createSearchArray, false); 
    inputButton.addEventListener("input", listElements, false);
    document.addEventListener("load", searchPlaceholder, false);
    wrapper.addEventListener("click", cleanPage, false);
} else if (window.attachEvent) {
    window.attachEvent("onload", createSearchArray);
    inputButton.attachEvent("oninput", listElements);
    window.attachEvent("onload", searchPlaceholder);
    wrapper.attachEvent("onclick", cleanPage);
}