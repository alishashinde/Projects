
document.addEventListener("DOMContentLoaded", () => {
    const registerForm = document.getElementById("registerForm");
    const loginForm = document.getElementById("loginForm");

    if (registerForm) {
        registerForm.addEventListener("submit", async (e) => {
            e.preventDefault();
            const data = {
                username: registerForm.username.value,
                email: registerForm.email.value,
                password: registerForm.password.value,
                location: registerForm.location.value
            };
            const res = await fetch("http://localhost:8080/api/users/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(data)
            });
            const json = await res.json();
            alert("Registered successfully!");
            console.log(json);
        });
    }

    if (loginForm) {
        loginForm.addEventListener("submit", async (e) => {
            e.preventDefault();
            const data = {
                email: loginForm.email.value,
                password: loginForm.password.value
            };
            const res = await fetch("http://localhost:8080/api/users/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(data)
            });
            const json = await res.json();
            if (json.token) {
                localStorage.setItem("token", json.token);
                alert("Login successful!");
            } else {
                alert("Login failed!");
            }
        });
    }
});
