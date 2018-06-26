$(document).ready(function() {
  $("#search-box").on("keyup", function() {
    filterRows()
  });

  $("#select-category").on("change", function() {
    filterRows()
  });

  $("#select-sortorder").on("change", function() {
    var value = $(this).val()
    var $wrapper = $(".elements");

    $wrapper.find(".filter-element").sort(function(a, b) {
        return ($(b).attr("data-" + value).toLowerCase()) < ($(a).attr("data-" + value).toLowerCase()) ? 1 : -1;
      })
      .appendTo($wrapper);
  });

  function filterRows() {
    const searchValue = $("#search-box").val().toLowerCase();
    const categoryValue = $("#select-category").val()

    if (categoryValue == "all" && searchValue == "") {
      $(".filter-element").filter(function() {
        $(this).toggle(true)
      });
    } else if (categoryValue == "all" && searchValue != "") {
      $(".filter-element").filter(function() {
        $(this).toggle($(this).attr("data-name").toLowerCase().indexOf(searchValue) > -1)
      });
    } else {
      $(".filter-element").filter(function() {
        const searchResult = $(this).attr("data-name").toLowerCase().indexOf(searchValue) > -1
        $(this).toggle(searchResult && categoryResult)
      });
    }
  }
});
