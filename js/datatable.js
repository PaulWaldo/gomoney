$(document).ready(function () {
  $("#transactions").DataTable({
  "serverSide":true,
  "ajax":"/transactions"
  });
})
