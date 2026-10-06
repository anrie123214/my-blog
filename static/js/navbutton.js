// Find the hamburger button and the navigation menu it controls.
const menuToggle = document.getElementById("menu-toggle");
const mainNav = document.getElementById("main-nav");

// Toggle the mobile navigation when the hamburger button is clicked.
menuToggle.addEventListener("click", () => {
    const isOpen = mainNav.classList.toggle("open");

    // Keep the accessibility state and label synchronized with the menu.
    menuToggle.setAttribute("aria-expanded", String(isOpen));
    menuToggle.setAttribute(
        "aria-label",
        isOpen ? "Close navigation" : "Open navigation"
    );
});
