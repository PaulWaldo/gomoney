$(document).ready(function () {
  $("#transactions").DataTable({
    serverSide: true,
    ajax: "/transactions",
    columns: [
      { data: "Payee" },
      // { data: "Type" },
      { data: "Amount" },
      { data: "Memo" },
      { data: "Date" },
    ],
  });
});
