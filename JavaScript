const username = "admin";
const password = "secret123";

function login() {
    const user = prompt("Enter username: ");
    const passwd = prompt("Enter password: ");
    if (user === username && passwd === password) {
        console.log("Access granted");
    } else {
        console.log("Access denied");
    }
}
