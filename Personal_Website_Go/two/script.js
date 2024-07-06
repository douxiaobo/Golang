function toggleNav() {
    var sidebar = document.getElementById("sidebar");
    if (sidebar.style.width === "0px") {
        sidebar.style.width = "250px";
    } else {
        sidebar.style.width = "0";
    }
}