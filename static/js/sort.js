$(document).ready(function() {
  $("#select-sortorder").on("change", function() {
    document.getElementById("info-equipmentseite").innerHTML = "";
    var value = $(this).val()
    var $wrapper = $(".elements");

    $wrapper.find(".filter-element").sort(function(a, b) {
        return ($(b).attr("data-" + value).toLowerCase()) < ($(a).attr("data-" + value).toLowerCase()) ? 1 : -1;
      })
      .appendTo($wrapper);
  });
});
