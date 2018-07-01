function search() {
  var input = document.getElementById("search-box").value.toLowerCase();
  var cat = document.getElementById("select-category").value;
  var elements = document.getElementsByClassName("filter-element");
  var counter = 0;
  for (var i = 0; i < elements.length; i++) {
    if (cat == "all") {
      if (elements[i].getAttribute("data-name").toLowerCase().includes(input)) {
        elements[i].style.display = "block";
      } else {
        elements[i].style.display = "none";
        counter++;
      }
    } else {
      if (elements[i].getAttribute("data-name").toLowerCase().includes(input) && elements[i].getAttribute("data-category") == cat) {
        elements[i].style.display = "block";
      } else {
        elements[i].style.display = "none";
        counter++;
      }
    }
  }

  if (counter == elements.length) {
    document.getElementById("info-equipmentseite").innerHTML = "Kein Equipment gefunden";
  }
}

function searchEquipAdmin() {
  var input = document.getElementById("search-box-admin").value.toLowerCase();
  var cat = document.getElementById("select-category-admin").value;
  var elements = document.getElementsByClassName("filter-element-admin");
  for (var i = 0; i < elements.length; i++) {
    if (cat == "all") {
      if (elements[i].getAttribute("data-name").toLowerCase().includes(input)) {
        elements[i].style.display = "flex";
      } else {
        elements[i].style.display = "none";
      }
    } else {
      if (elements[i].getAttribute("data-name").toLowerCase().includes(input) && elements[i].getAttribute("data-category") == cat) {
        elements[i].style.display = "flex";
      } else {
        elements[i].style.display = "none";
      }
    }
  }
}

function searchKundeAdmin() {
  var input = document.getElementById("search-box-admin-kunde").value.toLowerCase();
  var cat = document.getElementById("select-category-admin-kunde").value;
  var elements = document.getElementsByClassName("filter-element-admin-kunde");
  for (var i = 0; i < elements.length; i++) {
    if (cat == "all") {
      if (elements[i].getAttribute("data-name").toLowerCase().includes(input)) {
        elements[i].style.display = "flex";
      } else {
        elements[i].style.display = "none";
      }
    } else {
      if (elements[i].getAttribute("data-name").toLowerCase().includes(input) && elements[i].getAttribute("data-category") == cat) {
        elements[i].style.display = "flex";
      } else {
        elements[i].style.display = "none";
      }
    }
  }
}
