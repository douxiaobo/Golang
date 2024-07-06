const navmenu = document.querySelector(".navmenu");
const btn = document.querySelector(".hamburgerMenu");

function crossSymbol(){
    btn.classList.toggle("iconclass");
}
btn.addEventListener("click", ()=>{
    const datavisible = navmenu.getAttribute("data-hidden");
    
    if(datavisible ==="true"){
        navmenu.setAttribute("data-hidden",false);
    }
    else{
        navmenu.setAttribute("data-hidden",true);
    }
});