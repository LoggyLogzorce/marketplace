document.getElementById("authForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Предотвращение стандартного поведения формы (отправка)

    const login = document.getElementById("login").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:8080/api/auth/admin", {
        mode: "no-cors",
        method: "POST",
        headers: {
            "Access-Control-Allow-Headers" : "Content-Type",
            "Access-Control-Allow-Origin": "*",
            'Content-Type': 'application/json',
            "Access-Control-Allow-Methods": "OPTIONS,POST,GET,PATCH",
        },
        body: JSON.stringify({login, password})
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Authentication failed");
            }
            window.location.href = "/";
        })
        .catch(error => {
            console.error("Error:", error);
            // Обработка ошибки аутентификации
            alert("Authentication failed. Please check your credentials and try again.");
        });
});