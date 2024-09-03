function auth() {
    const login = document.getElementById("login").value;
    const password = document.getElementById("password").value;

    fetch("/auth", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({login, password})
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Authentication failed")
            }

            window.location.href = "/homepage";
        })
        .catch(error => {
            console.error("Error:", error);

            alert("Authentication failed. Please check your credentials and try again.")
        })
}