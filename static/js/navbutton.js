// Find the hamburger button and the navigation menu it controls.
const menuToggle = document.getElementById("menu-toggle");
const mainNav = document.getElementById("main-nav");
const menuClose = document.getElementById("menu-close");

// Keep the menu state, accessibility attributes, and page scrolling synchronized.
function setMenuOpen(isOpen) {
    mainNav.classList.toggle("open", isOpen);
    document.body.classList.toggle("menu-open", isOpen);

    menuToggle.setAttribute("aria-expanded", String(isOpen));
    menuToggle.setAttribute(
        "aria-label",
        isOpen ? "Close navigation" : "Open navigation"
    );
}

// Toggle the mobile navigation when the hamburger button is clicked.
menuToggle.addEventListener("click", () => {
    setMenuOpen(!mainNav.classList.contains("open"));
});

// Close the drawer when its white close button is clicked.
menuClose.addEventListener("click", () => setMenuOpen(false));

// Close the drawer after choosing a page from the mobile navigation.
mainNav.querySelectorAll("a").forEach((link) => {
    link.addEventListener("click", () => setMenuOpen(false));
});

// Close the mobile menu when clicking outside the drawer or menu button.
document.addEventListener("click", (event) => {
    const clickedInsideMenu = mainNav.contains(event.target);
    const clickedMenuButton = menuToggle.contains(event.target);
    const menuIsOpen = mainNav.classList.contains("open");

    if (menuIsOpen && !clickedInsideMenu && !clickedMenuButton) {
        setMenuOpen(false);
    }
});
//toggle dark mode
const themeToggle = document.getElementById("theme-toggle");
const themeIcon = document.getElementById("theme-icon");
const themeLabel = document.getElementById("theme-label");
const savedTheme = localStorage.getItem("theme");

const moonIcon = `
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
    viewBox="0 0 24 24" fill="none" stroke="currentColor"
    stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"
        fill="currentColor" fill-opacity="0.2"></path>
</svg>`;

const sunIcon = `
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
    viewBox="0 0 24 24" fill="none" stroke="currentColor"
    stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="4"
        fill="currentColor" fill-opacity="0.2"></circle>
    <path d="M12 2v2M12 20v2M4.93 4.93l1.41 1.41
        M17.66 17.66l1.41 1.41M2 12h2M20 12h2
        M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41"></path>
</svg>`;

if (savedTheme === "dark") {
    document.body.classList.add("dark-mode");
}

function updateThemeButton() {
    const darkMode = document.body.classList.contains("dark-mode");

    themeIcon.innerHTML = darkMode ? sunIcon : moonIcon;
    themeLabel.textContent = darkMode ? "Light mode" : "Dark mode";
    themeToggle.setAttribute(
        "aria-label",
        darkMode ? "Switch to light mode" : "Switch to dark mode"
    );
}

updateThemeButton();

themeToggle.addEventListener("click", () => {
    const darkMode = document.body.classList.toggle("dark-mode");

    localStorage.setItem("theme", darkMode ? "dark" : "light");
    updateThemeButton();
});

// Track the previous scroll position to detect the user's scroll direction.
let lastScrollTop = 0;
const header = document.querySelector(".site-header");

window.addEventListener("scroll", () => {
    const currentScroll = window.scrollY;

    if (currentScroll > lastScrollTop) {
        // Scrolling down: hide the header.
        header.classList.add("header-hidden");
    } else if (currentScroll < lastScrollTop) {
        // Scrolling up: show the header.
        header.classList.remove("header-hidden");
    }

    lastScrollTop = currentScroll;
}, { passive: true });
