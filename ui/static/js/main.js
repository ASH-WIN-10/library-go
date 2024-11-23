async function removeBook(id) {
    try {
        const response = await fetch(`/remove/${id}`, {
            method: "DELETE",
        });
        if (response.ok) {
            window.location.reload();
        }
    } catch (error) {
        console.error(error);
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
} catch (error) {}
