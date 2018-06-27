$(document).ready(function() {
  $("#searchbox-clients").on("keyup", function() {
    filterRows()
  });

  $("#select-category").on("change", function() {
    filterRows()
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
    } else if (categoryValue != "all" && searchValue == "") {
      $(".filter-element").filter(function() {
        $(this).toggle($(this).attr("data-category").indexOf(categoryValue) > -1)
      });
    } else {
      $(".filter-element").filter(function() {
        const searchResult = $(this).attr("data-name").toLowerCase().indexOf(searchValue) > -1
        const categoryResult = $(this).attr("data-category").indexOf(categoryValue) > -1
        $(this).toggle(searchResult && categoryResult)
      });
    }
  }
});
