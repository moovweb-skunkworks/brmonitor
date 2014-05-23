HTMLElement.prototype.createShadowRoot = HTMLElement.prototype.webkitCreateShadowRoot ||
                                         HTMLElement.prototype.mozCreateShadowRoot ||
                                         HTMLElement.prototype.createShadowRoot;

var Bathroom = function() {
  this.stalls = [];
}

Bathroom.prototype.getStalls = function() {
  var bathroom = this;
  var stallsRequest = new XMLHttpRequest();
  stallsRequest.open("GET", "/stalls", true);
  stallsRequest.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

  stallsRequest.addEventListener("load", function() {
    var bathroomResponse = JSON.parse(this.responseText);
    console.log(bathroomResponse.stalls);
    console.log(bathroom.stalls);
    if(bathroomResponse.stalls != bathroom.stalls) {
      bathroom.stalls = bathroomResponse.stalls;
      bathroom.render();
    }
  });

  stallsRequest.send();
}

Bathroom.prototype.makeStall = function(stall, i, stalls) {
  var stallTemplate = document.querySelector("template#stall");
  stallTemplate.createShadowRoot();
  var stallHtml = stallTemplate.content.querySelector(".stall").cloneNode(true);

  if(stall.id == 1) {
    stallHtml.classList.add("big");
  }
  if(stall.status == true) {
    stallHtml.classList.add("open");
  }
  stallHtml.querySelector(".stall-id").innerHTML = stall.id.toString();

  return stallHtml;
}

Bathroom.prototype.render = function() {
  var display = document.querySelector(".status-display");
  var brTemplate = document.querySelector("template#bathroom");
  brTemplate.createShadowRoot();
  var brHtml = brTemplate.content.querySelector(".bathroom").cloneNode(true);

  var stallsHtml = this.stalls.map(this.makeStall);

  stallsHtml.forEach(function(el) {
    brHtml.appendChild(el);
  });

  display.innerHTML = "";
  display.appendChild(brHtml);
}


document.addEventListener("DOMContentLoaded", function() {
  window.bathroom = new Bathroom();
  window.setInterval(function() {
    window.bathroom.getStalls();
  },
  2000);
  console.log(bathroom.getStalls());
});


