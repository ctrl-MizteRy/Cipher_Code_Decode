const codeType = document.getElementById("select_cipher");
const cipherType = document.getElementById("cipher-type");
const messageType = document.getElementById("code-or-decode");
const submitMsg = document.getElementById("submit-msg");
const msg = document.getElementById("text-msg");
const encriptKey = document.getElementById("key-container");
const cipherKey = document.getElementById("cipher-key");
const answerBox = document.getElementById("answer");
const ciphers = ["caesar", "substitution", "fence"];

function getCipherKey() {
    if (ciphers.includes(cipherType.value)) {
        if (messageType.value == "code") {
            encriptKey.removeAttribute("hidden");
            encriptKey.innerHTML = `
        Encript Key:
        <input type="text" id="message-key" value="" required placeholder="Please insert your encryption key">
`;
        } else if (messageType.value == "decode") {
            encriptKey.removeAttribute("hidden");
            if (cipherType.value == "substitution") {
                encriptKey.innerHTML = `
        Encript Key:
        <input type="text" id="message-key" value="" required placeholder="Please insert your decode encrypted key">
`;
            } else if (cipherType.value == "caesar" || cipherType.value == "fence") {

                encriptKey.innerHTML = `
        Encript Key:
        <input type="text" id="message-key" value="" placeholder="Please insert your decode encryption key if known">
`;
            }
        }
    } else {
        encriptKey.setAttribute("hidden", true)
    }
}

codeType.addEventListener('click', getCipherKey);

submitMsg.addEventListener("click", function(e) {
    e.preventDefault();
    submitButtonRespone();
});

/**
 * @param {string} msg 
 **/
function printoutDecode(msg) {
    let message = msg.result.split(":")[1];
    answerBox.innerHTML = `
<h3>
    The message:
    </h3>
<textarea id="ans-text" rows="8" cols="75">${message}</textarea>
    `;
}


/**
 * @param {string} key
 * @param {string} cipher_type 
 * @param {string} cod 
 * @returns {boolean}
 **/
function checkCiphperKey(cipher_type, key, cod) {
    switch (cipher_type) {
        case "caesar":
        case "fence":
            if (cod == "code") {
                let val = parseInt(key);
                if (!isNaN(val) && cipher_type == "fence" && val <= 0) {
                    return false;
                } else {
                    return !isNaN(val);
                }
            } else {
                return true;
            }
        case "substitution":
            return key.length == 26;
        default:
            return true
    }
}

function submitButtonRespone() {
    let cipher_type = cipherType.value;
    let key = document.getElementById('message-key');
    let keyVal = (key != null) ? key.value : "";
    let cod = messageType.value;

    if (cipher_type == "" || messageType.value == "") {
        alert(`Please select the type of cipher you want to use and code or decode option
`);
    } else if (ciphers.includes(cipher_type) &&
        keyVal == "" && cod == "code") {
        alert("Please input the cipher key");
        getCipherKey();
    } else {

        if (checkCiphperKey(cipher_type, keyVal, cod)) {
            connectBackend(keyVal);
        } else {
            alert(`Please input the correct format keyvalue for ${cipher_type}`);
            getCipherKey();
        }
    }
}

function connectBackend(keyVal) {
    fetch("http://localhost:8080/process", {
        method: "POST", headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            text: msg.value,
            key: keyVal,
            cod: messageType.value,
            cipher: cipherType.value,
        })
    })
        .then(response => response.json())
        .then(data => printoutDecode(data))
        .catch(error => console.error("Error:", error));
}

