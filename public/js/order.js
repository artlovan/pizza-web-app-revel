(function() {

  $('.pizza-size').delegate('.size', 'click', function(e) {
    $('.size').removeClass('selected');
    $(e.target).addClass('selected');
    calcPrice();
  });

  $('.toppings').delegate('[data-topping]', 'click', function(e) {
    $(e.target).toggleClass('selected');
    calcPrice();
  });

  function calcPrice() {
    var sizePrice = 0,
      toppingPrice = 0;

    var size = $('.pizza-size .size.selected');
    if (size.length > 0) {
      sizePrice = size[0].dataset['price'];
    }

    toppingPrice = $('.toppings [data-topping].selected').length * 1.5;
    if (toppingPrice === Math.floor(toppingPrice)) {
      toppingPrice += '.00';
    } else {
      toppingPrice += '0';
    }

    $('#sizePrice').html('$' + sizePrice);
    $('#toppingPrice').html('$' + toppingPrice);

    var totalPrice = parseFloat(sizePrice) + parseFloat(toppingPrice);
    if (totalPrice === Math.floor(totalPrice)) {
      totalPrice += '.00';
    } else {
      totalPrice += '0';
    }
    $('#totalPrice').html('$' + totalPrice);
  }

  $('#submit').click(function() {
    var order = {};
    order.Size = $('.pizza-size .selected').next().text();
    order.Toppings = [];
    $('.toppings .selected').each(function(idx, item) {
      order.Toppings.push(item.dataset['topping']);
    });

    $.ajax({
        method: 'post',
        url: '/orders/api/create',
        contentType: 'application/json',
        data: JSON.stringify(order)
    });
  });
})();
