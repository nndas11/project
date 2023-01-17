// display content on page loading and reloading 
async function displayContent(){
    const res = await fetch('https:localhost:3000/display');
    var data = await res.json();
    console.log(data);
    show(data);
}

displayContent()

function show(data){
    resultDiv.innerHTML = data.message + ':';
    for (let i = data.content.length-1; i>=0; i--) {
        resultDiv.innerHTML += "<p style='display:inline'> &nbsp" + data.content[i] + "</p>";
    }
}

module.exports = displayContent