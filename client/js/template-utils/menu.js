document.querySelector(".menu-button").addEventListener("click", e=>{
    e.currentTarget.classList.toggle("active");
    document.querySelector(".menu").classList.toggle("active");
})