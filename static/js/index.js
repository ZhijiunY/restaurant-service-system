const menu = document.querySelector(".menu");
const navOpen = document.querySelector(".hamburger");
const navClose = document.querySelector(".close");

const navLeft = menu.getBoundingClientRect().left;
navOpen.addEventListener("click", () => {
  if (navLeft < 0) {
    menu.classList.add("show");
  }
});

navClose.addEventListener("click", () => {
  if (navLeft < 0) {
    menu.classList.remove("show");
  }
});

// Fixed Nav
const nav = document.querySelector(".nav");
const navHeight = nav.getBoundingClientRect().height;
window.addEventListener("scroll", () => {
  const scrollHeight = window.pageYOffset;
  if (scrollHeight > navHeight) {
    nav.classList.add("fix-nav");
  } else {
    nav.classList.remove("fix-nav");
  }
});

// Scroll To
const links = [...document.querySelectorAll(".scroll-link")];
links.map(link => {
  link.addEventListener("click", e => {
    const href = e.target.getAttribute("href");
    if (href.startsWith("#")) {
      e.preventDefault();
      // 執行頁面內跳轉的代碼
      const id = href.slice(1);
      const element = document.getElementById(id);
      const fixNav = nav.classList.contains("fix-nav");
      let position = element.offsetTop - navHeight;

      if (!fixNav) {
        position = position;
      }

      window.scrollTo({
        top: position,
        left: 0,
      });
      menu.classList.remove("show");
    }
  });
});

AOS.init();

// TotalPrice
document.addEventListener('DOMContentLoaded', function() {
  const quantities = document.querySelectorAll('.quantity');
  const totalPriceElement = document.getElementById('totalPrice');

  function updateTotalPrice() {
    let total = 0;
    quantities.forEach(quantity => {
        const price = parseFloat(quantity.dataset.price);
        const amount = parseInt(quantity.value);
        total += price * amount;
    });
    console.log("Total Price:", total); // 調試信息
    if (totalPriceElement) {
        totalPriceElement.textContent = total.toFixed(2);
    } else {
        console.error('Total Price element not found');
    }
}

  quantities.forEach(quantity => {
      quantity.addEventListener('input', updateTotalPrice);
  });

  updateTotalPrice();
});

// submitOrder
function submitOrder() {
  var orderItems = [];
  document.querySelectorAll('.foodTypeSection').forEach(section => {
      section.querySelectorAll('.foodTable tr').forEach((row, index) => {
          if (index > 0) { // 跳過標題行
              var quantityInput = row.querySelector('.quantity');
              var quantity = parseInt(quantityInput.value, 10);
              if (quantity > 0) {
                  orderItems.push({
                      name: row.querySelector('input[name*="name"]').value,
                      description: row.querySelector('input[name*="description"]').value,
                      price: parseFloat(row.querySelector('input[name*="price"]').value),
                      quantity: quantity
                  });
              }
          }
      });
  });

  // 使用 AJAX 發送數據
  sendOrderData(orderItems);
}

// sendOrderData
function sendOrderData(orderItems) {
  fetch('/submit-order', {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify(orderItems)
  })
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
}
