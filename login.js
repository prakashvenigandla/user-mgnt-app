function login() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const errormsg = document.getElementById('result');

    fetch("http://localhost:8080/employee?username="+username+"&password="+password)

    .then(response => {
        if (response.ok) {
            throw new Error(response.status);
        }
        return response.json();
    })

    .then((data) => {
       if (data.message === "Login success") {
        errormsg.innerHTML = data.message;
       }else { 
        errormsg.innerHTML = "Invalid username and password";
       }

    } )
    return false;
}