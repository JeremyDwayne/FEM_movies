export class MovieItemComponent extends HTMLElement {
    constructor(movie) {
        super();
        this.movie = movie;
    }

    connectedCallback() {
        this.innerHTML = `
            <a href="/movie/${this.movie.id}">
                <article>
                    <img src="${this.movie.poster_url}" alt="${this.movie.title} Poster">
                    <p>${this.movie.title} (${this.movie.release_year})</p>
                </article>
            </a>
        `;
    }
}

customElements.define("movie-item", MovieItemComponent);