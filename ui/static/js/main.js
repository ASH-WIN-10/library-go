async function removeBook(id) {
    try {
        const response = await fetch(`/remove/${id}`, {
            method: "DELETE",
        });
        if (response.ok) {
            window.location.reload();
        }
    } catch (err) {
        console.error(err);
    }
}

async function updateBook(id) {
    try {
        const response = await fetch(`/update/${id}`, {
            method: "PUT",
        });
        if (response.ok) {
            window.location.reload();
        }
    } catch (err) {
        console.error(err);
    }
}

// Open form
document.querySelector("header > button").addEventListener("click", () => {
    document.querySelector(".addBook").showModal();
});

// Close form
document.getElementById("closeBtn").addEventListener("click", () => {
    document.querySelector(".addBook").close();
});

// Delete Book
try {
    const removeBtns = document.querySelectorAll("#removeBtn");
    for (const removeBtn of removeBtns) {
        removeBtn.onclick = (e) => {
            const id = e.target.attributes["data-id"].value;
            removeBook(id);
        };
    }

    const readBtns = document.querySelectorAll("#readBtn");
    for (const readBtn of readBtns) {
        readBtn.onclick = (e) => {
            const id = e.target.attributes["data-id"].value;
            updateBook(id);
        };
    }
} catch (error) {}
