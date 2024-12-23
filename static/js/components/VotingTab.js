// js/components/VotingTab.js
class VotingTab {
    constructor() {
        this.images = [];
        this.currentImageIndex = 0;
        this.votes = {};
        this.loading = true;
        this.error = null;
    }

    async init() {
        if (this.initialized) return; // Prevent multiple initializations
        
        console.log('Initializing VotingTab component...');
        try {
            this.loading = true;
            this.updateUI();
            await this.fetchImages();
            this.initialized = true;
            console.log('VotingTab Component initialized successfully');
        } catch (error) {
            console.error('VotingTab Component initialization failed:', error);
        }
    }

    async fetchImages() {
        try {
            const response = await axios.get(`${config.apiBaseURL}/api/cat-images`);
            this.images = response.data;
            this.error = null;
        } catch (error) {
            this.error = 'Error fetching cat images';
            console.error('Error fetching cat images:', error);
        } finally {
            this.loading = false;
            this.updateUI();
        }
    }

    async handleFavorite() {
        if (this.images.length > 0 && this.images[this.currentImageIndex]) {
            this.loading = true;
            this.updateUI();

            const currentImage = this.images[this.currentImageIndex];
            const urlParts = currentImage.url.split('/');
            const fileName = urlParts[urlParts.length - 1];
            const imageId = fileName.split('.')[0];

            try {
                const response = await axios.post(`${config.apiBaseURL}/api/favorites`, {
                    image_id: imageId,
                    sub_id: 'user-123'
                });
                if (response.status === 200) {
                    this.handleNextImage();
                }
            } catch (error) {
                console.error('Error adding to favorites:', error);
            } finally {
                this.loading = false;
                this.updateUI();
            }
        }
    }

    async handleVote(value) {
        if (this.images.length > 0 && this.images[this.currentImageIndex]) {
            this.loading = true;
            this.updateUI();

            const currentImage = this.images[this.currentImageIndex];
            const urlParts = currentImage.url.split('/');
            const fileName = urlParts[urlParts.length - 1];
            const imageId = fileName.split('.')[0];

            try {
                const response = await axios.post(`${config.apiBaseURL}/api/votes`, {
                    image_id: imageId,
                    sub_id: 'user-123',
                    value: value
                });
                console.log('Vote response:', response);
                if (response.status === 201) {
                    this.votes[imageId] = value;
                }
                this.handleNextImage();
            } catch (error) {
                console.error('Error voting:', error);
            } finally {
                this.loading = false;
                this.updateUI();
            }
        }
    }

    handleNextImage() {
        this.currentImageIndex = (this.currentImageIndex + 1) % this.images.length;
        this.updateUI();
    }

    updateUI() {
        const content = document.getElementById('tab-content');
        if (content) {
            content.innerHTML = this.render();
        }
    }

    render() {
        if (this.error) {
            return `<div class="alert alert-danger">${this.error}</div>`;
        }

        if (this.loading) {
            return `
                <div class="loader">
                    <div class="spinner-border text-primary" role="status">
                        <span class="visually-hidden">Loading...</span>
                    </div>
                </div>
            `;
        }

        if (this.images.length === 0) {
            return '<div class="text-center p-4">No images available</div>';
        }

        const currentImage = this.images[this.currentImageIndex];
        return `
            <div class="cat-image-container mb-3">
                <img src="${currentImage.url}" alt="Random cat" class="cat-image">
                <span class="image-id">${currentImage.id}</span>
            </div>
            <div class="d-flex justify-content-between align-items-center px-2">
                <i class="fas fa-heart action-icon" onclick="votingTab.handleFavorite()"></i>
                <div class="d-flex gap-4">
                    <i class="fas fa-thumbs-up action-icon" onclick="votingTab.handleVote(1)"></i>
                    <i class="fas fa-thumbs-down action-icon" onclick="votingTab.handleVote(-1)"></i>
                </div>
            </div>
        `;
    }
}