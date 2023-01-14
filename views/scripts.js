const pushForm = document.getElementById('push-form');
const popButton = document.getElementById('pop-button');
const topButton = document.getElementById('top-button');
const resultDiv = document.getElementById('result');

pushForm.addEventListener('submit', (event) => {
    event.preventDefault();
    const item = document.getElementById('item').value;
    alert(item);
    console.log(item);
    fetch('/push', {
        method: 'POST',
        body: JSON.stringify({number:  item}),
        headers: { 'Content-Type': 'application/json' }
    })
    .then(res => res.text())
    .then(data => {
        resultDiv.innerHTML = data;
    });
});

popButton.addEventListener('click', () => {
    fetch('/pop', {method: 'POST'})
    .then(res => res.text())
    .then(data => {
        resultDiv.innerHTML = data;
    });
});

topButton.addEventListener('click', () => {
    fetch('/top')
    .then(res => res.text())
    .then(data => {
        resultDiv.innerHTML = data;
    });
});