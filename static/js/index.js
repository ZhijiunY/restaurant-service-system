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
    e.preventDefault();

    const id = e.target.getAttribute("href").slice(1);
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
  });
});

// document.querySelector('.loginBox input[type="submit"]').addEventListener('click', function() {
//   window.location.href = '/login';
// });

AOS.init();

// document.addEventListener('DOMContentLoaded', function() {
//   function calculateTotal() {
//     var total = 0;
//     document.querySelectorAll('.quantity').forEach(function(input) {
//       var price = parseFloat(input.getAttribute('data-price'));
//       var quantity = parseInt(input.value);
//       total += price * quantity;
//     });
//     document.getElementById('totalPrice').textContent = total.toFixed(2) + '元';
//   }

//   document.querySelectorAll('.quantity').forEach(function(input) {
//     input.addEventListener('change', calculateTotal);
//   });

//   calculateTotal();
// });

// document.addEventListener('DOMContentLoaded', function() {
//   function calculateTotal() {
//       var total = 0;
//       document.querySelectorAll('.quantity').forEach(function(input) {
//           var price = parseFloat(input.getAttribute('data-price'));
//           var quantity = parseInt(input.value);
//           total += price * quantity;
//       });
//       document.getElementById('totalPrice').textContent = total.toFixed(2) + '元';
//   }

//   document.querySelectorAll('.quantity').forEach(function(input) {
//       input.addEventListener('change', calculateTotal);
//   });

//   calculateTotal();
// });

// 
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
      totalPriceElement.textContent = total.toFixed(2);
  }

  quantities.forEach(quantity => {
      quantity.addEventListener('input', updateTotalPrice);
  });

  updateTotalPrice();
});
