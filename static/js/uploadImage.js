function previewFile() {
  var preview = document.getElementById("inputProfilbild");
  preview.click();
}

function loadFile() {
  var img = document.getElementById("profilbild");
  var preview = document.getElementById("inputProfilbild");
  var file = preview.files[0];
  var reader = new FileReader();

  reader.onloadend = function() {
    img.src = reader.result;
  }

  if (file) {
    reader.readAsDataURL(file);
  } else {
    preview.src = "";
  }
}
