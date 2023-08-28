const create_acc = document.querySelector("#create_acc");
const login_acc = document.querySelector("#login_btn");
const public_post = document.querySelector("#scd_div");

create_acc.addEventListener("click", () => {
    document.location = "/register";
});

login_acc.addEventListener("click", () => {
    document.location = "/login";
});

public_post.addEventListener("click", () => {
    document.location = "/public";
});