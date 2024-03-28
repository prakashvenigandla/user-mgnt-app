class User {
    #_name;
    #_gender;
    #_email;
    #_country;

    constructor(name, gender, email, country) {
        this.#_name = name;
        this.#_gender = gender;
        this.#_email = email;
        this.#_country = country;
    }

    toString() {
        return `Name: ${this.#_name} Gender: ${this.#_gender} Email: ${this.#_email} Country: ${this.#_country}`;
    }

    equals(otherUser) {
        return this.#_name.toLowerCase() === otherUser.#_name.toLowerCase() &&
               this.#_email.toLowerCase() === otherUser.#_email.toLowerCase();
    }
}

function getUserDetails() {
    return new Promise(resolve => {
        let userDetails = [];
        let prompts = ['Name', 'Gender', 'Email', 'Country'];
        let currentPrompt = 0;

        function promptUser() {
            process.stdout.write(`${prompts[currentPrompt]}: `);
        }

        process.stdin.on('data', data => {
            const input = data.toString().trim();
            userDetails.push(input);

            if (currentPrompt < prompts.length - 1) {
                currentPrompt++;
                promptUser();
            } else {
                process.stdin.removeAllListeners('data');
                resolve(new User(...userDetails));
            }
        });

        promptUser();
    });
}

(async () => {
    const user1 = await getUserDetails();
    const user2 = await getUserDetails();

    console.log("\nUser 1");
    console.log(user1.toString());

    console.log("\nUser 2");
    console.log(user2.toString());

    if (user1.equals(user2)) {
        console.log("User 1 is same as User 2");
    } else {
        console.log("User 1 and User 2 are different");
    }

    process.stdin.destroy();
})();
