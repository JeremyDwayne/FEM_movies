import { API } from "./services/API.js";
import "./components/HomePage.js";

window.app = {
    search: (event) => {
        event.preventDefault();
        const searchInput = document.querySelector('input[type="search"]').value.trim();
        if (searchInput) {
            window.location.href = `/search?q=${searchInput}`;
        }
    },
    api: API,
}