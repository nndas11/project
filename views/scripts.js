const pushForm = document.getElementById('push-form');
const popButton = document.getElementById('pop-button');
const topButton = document.getElementById('top-button');
const resultDiv = document.getElementById('result');
const displayStackButton = document.getElementById('display-stack-button');
const sizeForm = document.getElementById("size-form")



sizeForm.addEventListener('submit', (event)=>{
    event.preventDefault();
    const item = document.getElementById('size').value;
    console.log(item);

    //validation using javascript
    const pattern = /^\d+$/;
    if (!pattern.test(item)) {
        alert("Invalid input. Only digits are allowed.");
        return;
    }

    fetch('/declare', {
        method: 'POST',
        body: JSON.stringify({number:  item}),
        headers: { 
            'Content-Type': 'application/json' 
        }
    })
    .then(res => res.json())
    .then(data => {
        resultDiv.innerHTML = data.message + data.content;
    }).catch(error => console.log(error));;
})


pushForm.addEventListener('submit', (event) => {
    event.preventDefault();
    const item = document.getElementById('item').value;

    //validation using javascript
    const pattern = /^\d+$/;
    if (!pattern.test(item)) {
        alert("Invalid input. Only digits are allowed.");
        return;
    }

    fetch('/push', {
        method: 'POST',
        body: JSON.stringify({number:  item}),
        headers: { 
            'Content-Type': 'application/json' 
        }
    })
    .then(res => res.json())
    .then(data => {
        resultDiv.innerHTML = data.message + data.content;
    }).catch(error => console.log(error));;
});

popButton.addEventListener('click', () => {
    fetch('/pop', {method: 'DELETE'})
    .then(res => res.json())
    .then(data => {
        resultDiv.innerHTML = data.message + data.content;
    }).catch(error => console.log(error));
});

topButton.addEventListener('click', () => {
    fetch('/top')
    .then(res => res.json())
    .then(data => {
        resultDiv.innerHTML = data.message  + data.content;
    }).catch(error => console.log(error));
});



displayStackButton.addEventListener('click', () => {
    fetch('/display')
    .then(res => res.json())
    .then(data => {
        resultDiv.innerHTML = data.message + ':';
        for (let i = data.content.length-1; i>=0; i--) {
            resultDiv.innerHTML += "<p style='display:inline'> &nbsp" + data.content[i] + "</p>";
        }
    }).catch(error => console.log(error));
});


// // display content on page loading and reloading 
// async function displayContent(){
//     const res = await fetch('/display');
//     var data = await res.json();
//     console.log(data);
//     show(data);
// }

// displayContent()

// function show(data){
//     resultDiv.innerHTML = data.message + ':';
//     for (let i = data.content.length-1; i>=0; i--) {
//         resultDiv.innerHTML += "<p style='display:inline'> &nbsp" + data.content[i] + "</p>";
//     }
// }


