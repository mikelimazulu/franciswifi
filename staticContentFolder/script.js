document.getElementById('loginForm').addEventListener('submit', function (event) {
    event.preventDefault();
    document.getElementById('errorModal').style.display = 'flex';
});

document.querySelector('.close').addEventListener('click', function () {
    document.getElementById('errorModal').style.display = 'none';
});

document.getElementById('emailForm').addEventListener('submit', function (event) {
    event.preventDefault();
    alert('Merci! Vous recevrez votre certificat cadeau par e-mail.');
    window.close();
});
