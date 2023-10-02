const wrapper = document.querySelector(".wrapper");
const login_link = document.querySelector(".login-link");
const register_link = document.querySelector(".register-link");
const btn_login = document.querySelector(".btnLogin");
const icon_close = document.querySelector(".icon-close");
const inputs = document.querySelectorAll("input");

register_link.addEventListener("click", () => {
    wrapper.classList.add("active");
});

login_link.addEventListener("click", () => {
    wrapper.classList.remove("active");
});

btn_login.addEventListener("click", () => {
    inputs.forEach(i => i.value = "");
    wrapper.classList.add("active-popup");
});

icon_close.addEventListener("click", () => {
    wrapper.classList.remove("active-popup");
});