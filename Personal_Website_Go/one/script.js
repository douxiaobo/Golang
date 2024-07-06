// script.js
document.getElementById('menuToggle').addEventListener('click', function() {
    document.querySelector('.nav-links').classList.toggle('show');
});

// 添加CSS以支持.show类
.nav-links.show {
    display: block; /* 或使用flex，grid等，根据需要调整 */
}