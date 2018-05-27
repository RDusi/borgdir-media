function previewFile() {
  var preview = document.getElementById("inputUploadBild");
  preview.click();
}

function loadFile() {
  var img = document.getElementById("showbild");
  var preview = document.getElementById("inputUploadBild");
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
