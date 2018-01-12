// Service Worker installing...
require('offline-plugin/runtime').install();
import "../less/landing/main.less";

import "./template-utils/menu.js";

document.addEventListener("DOMContentLoaded", _ => {
    document.getElementById("contact-form").addEventListener("submit", e => {
        e.preventDefault();

        const name = document.getElementById("name");
        const email = document.getElementById("email");
        const message = document.getElementById("message");

        const captcha = grecaptcha.getResponse();

        console.log(captcha)

        if (!captcha) return alert("I'm just wondering? Are you Robot ?")
        if (!name || !email || !message) return alert("Missing fields.");

        const ajax = new XMLHttpRequest();
        ajax.open("POST", "/api/contact", true);
        ajax.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
        ajax.onload = function () {
            console.log(ajax.response)
        }
        ajax.send(JSON.stringify({
            Email: email.value,
            Name: name.value,
            Message: message.value,
            Captcha: captcha
        }))
    })
})