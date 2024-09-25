var subnav = document.querySelectorAll('li a')
const current = window.location.href;
window.onload = select

function select(select) {
  subnav.forEach(function(anchor) {
    if (current.indexOf(anchor.href) != -1){
      anchor.classList.add("selected");
    } else {
      anchor.classList.remove("selected");
   }
  });
}

/* selecting mainbtn and subnavbtn based on current url */
const btnnav = document.querySelectorAll('.btnnav');
const subnavbtn = document.querySelectorAll('li a')
const data = document.querySelector(".pagenav ul").getAttribute("data-mainnav");

window.addEventListener("DOMContentLoaded", (event) => {
  /* main nav */
  btnnav.forEach(el => {
    if (data.includes(el.pathname)) {
      el.classList.add("selected");
    } else {
      el.classList.remove("selected")
    }
  });
});