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
      { data: "payee" },
      // { data: "Type" },
      { data: "amount" },
      { data: "memo" },
      { data: "date" },
    ],
  });
});
