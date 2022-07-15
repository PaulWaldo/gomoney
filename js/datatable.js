$(document).ready(function () {
  $("#transactions").DataTable({
    serverSide: true,
    ajax: "/transactions",
    scrollY: "50em",
    deferRender: true,
    // scroller: true,
    scroller: {
        displayBuffer: 20
    },
    columns: [
      { data: "Payee" },
      // { data: "Type" },
      { data: "Amount" },
      { data: "Memo" },
      { data: "Date" },
    ],
  });
});
