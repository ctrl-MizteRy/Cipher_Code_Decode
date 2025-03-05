const codeType = document.getElementById("select_cipher");
const cipherType = document.getElementById("cipher-type");
const messageType = document.getElementById("code-or-decode");
const submitMsg = document.getElementById("submit-msg");
const msg = document.getElementById("text-msg");
const encodeKey = document.getElementById("encode-key-container");
const decodeKey = document.getElementById("decode-key-container");
const cipherKey = document.getElementById("cipher-key");
const answerBox = document.getElementById("answer");

function getCipherKey() {
    if (messageType.value == "code") {
        encodeKey.removeAttribute("hidden");
        if (!decodeKey.hasAttribute("hidden")) {
            decodeKey.setAttribute("hidden", true);
        }
    } else {
        decodeKey.removeAttribute("hidden");
        if (!encodeKey.hasAttribute("hidden")) {
            encodeKey.setAttribute("hidden", true);
        }
    }
}

codeType.addEventListener('click', getCipherKey);

submitMsg.addEventListener("click", function(e) {
    e.preventDefault();


    fetch("http://localhost:8080/process", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            text: msg.value
        })
    })
        .then(response => response.json())
        .then(data => printoutDecode(data))
        .catch(error => console.error("Error:", error));
});

function printoutDecode(msg) {
    let message = msg.result.split(":")[1];
    answerBox.innerHTML = `
<h3>
    The message:
    </h3>
<textarea rows="8" cols="75">${message}</textarea>
    `;
}
