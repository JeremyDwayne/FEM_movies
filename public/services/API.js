export const API = {
    baseURL: "/api/",
    getTopMovies: async () => {
        return await API.fetch("movies/top");
    },
    getRandomMovies: async () => {
        return await API.fetch("movies/random");
    },
    searchMovies: async (query, order, genre) => {
        return await API.fetch("movies/search", { query, order, genre });
    },
    getMovie: async (id) => {
        return await API.fetch(`movies/${id}`);
    },
    getGenres: async () => {
        return await API.fetch("genres");
    },
    fetch: async (url, options = {}) => {
        try {
            const queryString = new URLSearchParams(options).toString();
            const response = await fetch(API.baseURL + url + (queryString ? `?${queryString}` : ""));
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return await response.json();
        } catch (error) {
            console.error("API error:", error);
            return null;
        }
    },
}