// js/components/BreedsTab.js
class BreedsTab {
    constructor() {
        this.breeds = [];
        this.selectedBreed = null;
        this.catImage = null;
        this.isDropdownOpen = false;
        this.loading = true;
        this.error = null;
    }

    // Add to each component's init method
    async init() {
        if (this.initialized) return; // Prevent multiple initializations

        console.log('Initializing BreedsTab component...');
        try {
            this.loading = true;
            this.updateUI();
            await this.fetchBreeds();
            this.initialized = true;
            console.log('BreedsTab Component initialized successfully');
        } catch (error) {
            console.error('BreedsTab Component initialization failed:', error);
        }
    }

    async fetchBreeds() {
        try {
            const response = await axios.get(`${config.apiBaseURL}/api/breeds`);
            this.breeds = response.data;
            if (this.breeds.length > 0) {
                this.selectedBreed = this.breeds[0];
                await this.fetchCatImageByBreed(this.breeds[0].id);
                }
            } catch (error) {
                this.error = 'Failed to load breeds. Please try again later.';
                console.error(error);
            } finally {
                this.loading = false;
                this.updateUI();
            }
    }
    
    async fetchCatImageByBreed(breedId) {
        try {
            this.loading = true;
            this.updateUI();
            const response = await axios.get(`${config.apiBaseURL}/api/cat-images/by-breed?breed_id=${breedId}`);
            if (response.data.length > 0) {
                this.catImage = response.data[0];
            }
        } catch (error) {
            this.error = 'Failed to load cat image. Please try again later.';
            console.error(error);
        } finally {
            this.loading = false;
            this.updateUI();
        }
    }
    
    toggleDropdown() {
        this.isDropdownOpen = !this.isDropdownOpen;
        this.updateUI();
    }

    async handleBreedSelect(breed) {
        this.selectedBreed = breed;
        this.isDropdownOpen = false;
        await this.fetchCatImageByBreed(breed.id);
    }

    updateUI() {
        const content = document.getElementById('tab-content');
        if (content) {
            content.innerHTML = this.render();
        }
    }

    render() {
        if (this.loading && this.breeds.length === 0) {
            return `
                <div class="loader">
                    <div class="spinner-border text-primary" role="status">
                        <span class="visually-hidden">Loading...</span>
                    </div>
                </div>
            `;
        }

        if (this.error) {
            return `<div class="alert alert-danger">${this.error}</div>`;
        }

        return `
            <div class="breeds-dropdown mb-4">
                <button class="btn btn-outline-secondary w-100 d-flex justify-content-between align-items-center"
                        onclick="breedsTab.toggleDropdown()">
                    <span>${this.selectedBreed ? this.selectedBreed.name : 'Select a breed'}</span>
                    <i class="fas fa-chevron-${this.isDropdownOpen ? 'up' : 'down'}"></i>
                </button>
                <div class="breeds-list ${this.isDropdownOpen ? 'show' : ''}">
                    ${this.breeds.map(breed => `
                        <button class="dropdown-item w-100 text-start py-2 px-3"
                                onclick="breedsTab.handleBreedSelect(${JSON.stringify(breed).replace(/"/g, '&quot;')})">
                            ${breed.name}
                        </button>
                    `).join('')}
                </div>
            </div>
            <div class="cat-image-container mb-3">
                ${this.loading ? `
                    <div class="loader">
                        <div class="spinner-border text-primary" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    </div>
                ` : this.catImage ? `
                    <img src="${this.catImage.url}" alt="${this.selectedBreed.name} cat" class="cat-image">
                ` : `
                    <div class="d-flex justify-content-center align-items-center h-100">
                        No image available for this breed
                    </div>
                `}
            </div>
            ${this.selectedBreed ? `
                <div class="mt-4">
                    <h3 class="fw-bold">
                        ${this.selectedBreed.name} 
                        <span class="text-muted">(${this.selectedBreed.origin})</span>
                        <span class="text-sm italic fw-normal text-muted">${this.selectedBreed.id}</span>
                    </h3>
                    <p class="mt-2 text-muted">${this.selectedBreed.description}</p>
                    <a href="https://en.wikipedia.org/wiki/${this.selectedBreed.name}_cat" 
                        target="_blank" 
                        rel="noopener noreferrer" 
                        class="text-orange-500 mt-2 d-block text-uppercase fw-semibold text-decoration-none">
                        Wikipedia
                    </a>
                </div>
            ` : ''}
        `;
    }
}