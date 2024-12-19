document.addEventListener('DOMContentLoaded', () => {
    const currentImage = document.getElementById('currentImage');
    const loveBtn = document.getElementById('loveBtn');
    const likeBtn = document.getElementById('likeBtn');
    const dislikeBtn = document.getElementById('dislikeBtn');
    const loadingSpinner = document.querySelector('.loading-spinner');
    
    let images = [];
    let currentIndex = 0;
    const subID = 'user-123'; // You should generate this dynamically per user

    function showLoading() {
        loadingSpinner.classList.remove('d-none');
        currentImage.style.opacity = '0.3';
        setButtonsState(true);
    }

    function hideLoading() {
        loadingSpinner.classList.add('d-none');
        currentImage.style.opacity = '1';
        setButtonsState(false);
    }

    async function fetchImages() {
        showLoading();
        try {
            const response = await fetch('https://api.thecatapi.com/v1/images/search?limit=10');
            if (!response.ok) {
                throw new Error('Failed to fetch images');
            }

            images = await response.json();
            currentIndex = 0;
            await showCurrentImage();
        } catch (error) {
            console.error('Error fetching images:', error);
            alert('Failed to load images. Please try again.');
        } finally {
            hideLoading();
        }
    }

    async function showCurrentImage() {
        if (images.length === 0) return;

        showLoading();
        try {
            await new Promise((resolve, reject) => {
                currentImage.onload = resolve;
                currentImage.onerror = reject;
                currentImage.src = images[currentIndex].url;
            });
            preloadNextImage();
        } catch (error) {
            console.error('Error loading image:', error);
            currentImage.src = 'static/images/placeholder.jpg';
        } finally {
            hideLoading();
        }
    }

    async function addToFavorites(imageId) {
        try {
            const response = await fetch('/api/favorites', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    image_id: imageId,
                    sub_id: subID
                })
            });

            if (!response.ok) {
                throw new Error('Failed to add to favorites');
            }

            // Refresh favorites tab content
            if (typeof refreshFavorites === 'function') {
                refreshFavorites();
            }

            return await response.json();
        } catch (error) {
            console.error('Error adding to favorites:', error);
            alert('Failed to add to favorites. Please try again.');
        }
    }

    function setButtonsState(disabled) {
        loveBtn.disabled = disabled;
        likeBtn.disabled = disabled;
        dislikeBtn.disabled = disabled;
    }

    function preloadNextImage() {
        if (images.length > currentIndex + 1) {
            const img = new Image();
            img.src = images[currentIndex + 1].url;
        }
    }

    // Event Listeners
    loveBtn.addEventListener('click', async () => {
        if (images.length > 0) {
            const currentImage = images[currentIndex];
            loveBtn.disabled = true;
            await addToFavorites(currentImage.id);
            loveBtn.disabled = false;
            await fetchImages();
        }
    });

    likeBtn.addEventListener('click', async () => {
        if (images.length === 0) return;
        currentIndex = (currentIndex + 1) % images.length;
        await showCurrentImage();
    });

    dislikeBtn.addEventListener('click', async () => {
        if (images.length === 0) return;
        currentIndex = (currentIndex - 1 + images.length) % images.length;
        await showCurrentImage();
    });

    // Initialize
    fetchImages();
});

// Export refresh function for tab switching
window.refreshVoting = function() {
    fetchImages();
};