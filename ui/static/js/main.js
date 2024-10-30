// Open form
document.querySelector("header > button").addEventListener("click", () => {
	document.querySelector(".addBook").showModal();
});

// Close form
document.getElementById("closeBtn").addEventListener("click", () => {
	document.querySelector(".addBook").close();
});
