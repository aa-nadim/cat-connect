class BreedsTab {
    constructor() {
        this.breeds = [];
        this.selectedBreed = null;
        this.catImages = [];
        this.currentImageIndex = 0;
        this.isDropdownOpen = false;
        this.loading = true;
        this.error = null;
        this.slideInterval = null;
    }

    async init() {
        if (this.initialized) return;

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
            this.catImages = response.data;
            this.currentImageIndex = 0;
            // Start slideshow after loading images
            this.startSlideshow();
        } catch (error) {
            this.error = 'Failed to load cat image. Please try again later.';
            console.error(error);
        } finally {
            this.loading = false;
            this.updateUI();
        }
    }


    startSlideshow() {
        this.stopSlideshow(); // Clear any previous interval
        if (this.catImages.length > 1) {
            this.slideInterval = setInterval(() => {
                // Only update if we're still on the breeds tab
                if (window.location.pathname === '/breeds') {
                    this.currentImageIndex = (this.currentImageIndex + 1) % this.catImages.length;
                    this.updateUI();
                } else {
                    this.stopSlideshow(); // Stop if we're not on breeds tab
                }
            }, 3000); // Changed to 3 seconds for smoother transitions
        }
    }

    stopSlideshow() {
        if (this.slideInterval) {
            clearInterval(this.slideInterval);
            this.slideInterval = null;
        }
    }

    updateUI() {
        const content = document.getElementById('tab-content');
        // Only update if we're on the breeds tab
        if (content && window.location.pathname === '/breeds') {
            content.innerHTML = this.render();
        }
    }

    async handleBreedSelect(breed) {
        this.stopSlideshow(); // Stop current slideshow
        this.selectedBreed = breed;
        this.isDropdownOpen = false;
        await this.fetchCatImageByBreed(breed.id);
    }

    toggleDropdown() {
        this.isDropdownOpen = !this.isDropdownOpen;
        this.updateUI();
    }

    async handleBreedSelect(breed) {
        this.stopSlideshow(); // Stop current slideshow
        this.selectedBreed = breed;
        this.isDropdownOpen = false;
        await this.fetchCatImageByBreed(breed.id);
        // fetchCatImageByBreed will start the new slideshow
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
            <div class="cat-image-container mb-3 position-relative">
                ${this.loading ? `
                    <div class="loader">
                        <div class="spinner-border text-primary" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    </div>
                ` : this.catImages.length > 0 ? `
                    <img src="${this.catImages[this.currentImageIndex].url}" 
                         alt="${this.selectedBreed.name} cat" 
                         class="cat-image w-100" />

                    <div class="dots-container position-absolute bottom-0 w-100 d-flex justify-content-center gap-2 p-2">
                        ${this.catImages.map((_, index) => `
                            <span class="rounded-circle ${index === this.currentImageIndex ? 'bg-primary' : 'bg-secondary'}"
                                style="width: 10px; height: 10px; display: inline-block;"></span>
                        `).join('')}
                    </div>
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