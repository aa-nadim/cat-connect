// js/components/FavsTab.js
class FavsTab {
    constructor() {
        this.favoriteCats = [];
        this.loading = true;
        this.viewMode = 'grid';
    }

    // Add to each component's init method
    async init() {
        if (this.initialized) return; // Prevent multiple initializations

        console.log('Initializing FavsTab component...');
        try {
            // this.loading = true;
            this.updateUI();
            await this.fetchFavorites();
            this.initialized = true;
            console.log('FavsTab Component initialized successfully');
        } catch (error) {
            console.error('FavsTab Component initialization failed:', error);
        }
    }

    async fetchFavorites() {
        this.loading = true;
        this.updateUI();

        try {
            const response = await axios.get(`${config.apiBaseURL}/api/favorites`, {
                params: { sub_id: 'user-123' }
            });
            this.favoriteCats = response.data;
        } catch (error) {
            console.error('Error fetching favorites:', error);
        } finally {
            this.loading = false;
            this.updateUI();
        }
    }

    async removeFavorite(favoriteId) {
        this.loading = true;
        this.updateUI();

        try {
            await axios.delete(`${config.apiBaseURL}/api/favorites/${favoriteId}`);
            await this.fetchFavorites();
        } catch (error) {
            console.error('Error removing favorite:', error);
        }
    }

    setViewMode(mode) {
        this.viewMode = mode;
        this.updateUI();
    }

    updateUI() {
        const content = document.getElementById('tab-content');
        if (content) {
            content.innerHTML = this.render();
        }
    }

    render() {
        return `
            <div class="d-flex flex-column h-100">
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <div class="btn-group">
                        <button class="btn ${this.viewMode === 'grid' ? 'btn-primary' : 'btn-light'}"
                                onclick="favsTab.setViewMode('grid')">
                            <i class="fas fa-th"></i>
                        </button>
                        <button class="btn ${this.viewMode === 'list' ? 'btn-primary' : 'btn-light'}"
                                onclick="favsTab.setViewMode('list')">
                            <i class="fas fa-list"></i>
                        </button>
                    </div>
                </div>
                ${this.loading ? `
                    <div class="loader">
                        <div class="spinner-border text-primary" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    </div>
                ` : `
                    <div class="flex-grow-1 overflow-auto" style="max-height: 500px;">
                        <div class="row g-4 ${this.viewMode === 'list' ? 'flex-column' : ''}">
                            ${this.favoriteCats.map(cat => `
                                <div class="${this.viewMode === 'grid' ? 'col-6' : 'col-12 mb-4'}">
                                    <div class="position-relative">
                                        <img src="${cat.image.url}" 
                                            alt="Favorite cat ${cat.id}"
                                            class="img-fluid rounded"
                                            style="height: ${this.viewMode === 'grid' ? '200px' : '300px'}; 
                                                    width: 100%; 
                                                    object-fit: cover;">
                                        <button class="btn btn-danger btn-sm position-absolute top-0 end-0 m-2"
                                                onclick="favsTab.removeFavorite(${cat.id})">
                                            <i class="fas fa-trash"></i>
                                        </button>
                                    </div>
                                </div>
                            `).join('')}
                        </div>
                    </div>
                `}
            </div>
        `;
    }
}