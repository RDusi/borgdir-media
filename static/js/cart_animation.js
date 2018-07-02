function addtocart(id) {
  var eq = document.getElementById("equip");
  var bild = document.getElementById("equipbild" + id)
  eq.style.visibility = "visible";
  eq.src = bild.src;
  var pos = 0;
  var id = setInterval(frame, 1);

  function frame() {
    if (pos == 50) {
      clearInterval(id);
    } else {
      pos++;
      eq.style.left = pos + 'px';
    }
  }
}
