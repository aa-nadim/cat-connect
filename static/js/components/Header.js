// js/components/Header.js
class Header {
    constructor() {
        this.isMenuOpen = false;
    }

    toggleMenu() {
        this.isMenuOpen = !this.isMenuOpen;
        const navMenu = document.querySelector('.navbar-collapse');
        if (this.isMenuOpen) {
            navMenu.classList.add('show');
        } else {
            navMenu.classList.remove('show');
        }
    }

    render() {
        return `
            <header class="navbar navbar-expand-lg navbar-light border-bottom py-3">
                <div class="container">
                    <a class="navbar-brand fw-bold" href="#">TheCatAPI</a>
                    <button class="navbar-toggler" type="button" onclick="header.toggleMenu()">
                        <i class="fas ${this.isMenuOpen ? 'fa-times' : 'fa-bars'}"></i>
                    </button>
                    <div class="collapse navbar-collapse">
                        <ul class="navbar-nav ms-auto">
                            <li class="nav-item"><a class="nav-link" href="#">PRICING</a></li>
                            <li class="nav-item"><a class="nav-link" href="#">DOCUMENTATION</a></li>
                            <li class="nav-item"><a class="nav-link" href="#">MORE APIS</a></li>
                            <li class="nav-item"><a class="nav-link" href="#">SHOWCASE</a></li>
                        </ul>
                    </div>
                </div>
            </header>
        `;
    }
}
