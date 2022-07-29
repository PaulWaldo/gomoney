// var selectedAccountId = "";
var table;

function getTransactions(accountId) {
  selectedAccountId = accountId;
  var url = `/transactions/${selectedAccountId}`;
  table.ajax.url(url).load();
}

$(document).ready(function () {
  table = $("#transactions").DataTable({
    serverSide: true,
    ajax: `/transactions/`,
    // ${selectedAccountId}`,
    scrollY: "50em",
    deferRender: true,
    // scroller: true,
    scroller: {
      displayBuffer: 20,
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
