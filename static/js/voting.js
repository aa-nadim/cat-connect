document.addEventListener('DOMContentLoaded', () => {
    const currentImage = document.getElementById('currentImage');
    const loveBtn = document.getElementById('loveBtn');
    const likeBtn = document.getElementById('likeBtn');
    const dislikeBtn = document.getElementById('dislikeBtn');
    const loadingSpinner = document.querySelector('.loading-spinner');
    
    let images = [];
    let currentIndex = 0;
    const subID = 'user-123'; // You should generate this dynamically per user

    // Show loading state
    function showLoading() {
        loadingSpinner.classList.remove('d-none');
        currentImage.style.opacity = '0.3';
    }

    // Hide loading state
    function hideLoading() {
        loadingSpinner.classList.add('d-none');
        currentImage.style.opacity = '1';
    }

    // Fetch images from the API
    async function fetchImages() {
        showLoading();
        try {
            const response = await fetch('https://api.thecatapi.com/v1/images/search?limit=10', {
                headers: {
                    'x-api-key': 'your-api-key-here' // Replace with your actual API key
                }
            });
            
            if (!response.ok) {
                throw new Error('Failed to fetch images');
            }

            images = await response.json();
            currentIndex = 0;
            await showCurrentImage();
        } catch (error) {
            console.error('Error fetching images:', error);
            // Show error message to user
            alert('Failed to load images. Please try again.');
        } finally {
            hideLoading();
        }
    }

    // Show current image with preloading
    async function showCurrentImage() {
        if (images.length === 0) {
            return;
        }

        showLoading();

        // Create a promise to handle image loading
        const loadImage = () => {
            return new Promise((resolve, reject) => {
                currentImage.onload = resolve;
                currentImage.onerror = reject;
                currentImage.src = images[currentIndex].url;
            });
        };

        try {
            await loadImage();
        } catch (error) {
            console.error('Error loading image:', error);
            currentImage.src = 'path/to/fallback-image.jpg'; // Add a fallback image
        } finally {
            hideLoading();
        }
    }

    // Add to favorites
    async function addToFavorites(imageId) {
        try {
            const response = await fetch('/api/favorites', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'x-api-key': 'your-api-key-here' // Replace with your actual API key
                },
                body: JSON.stringify({
                    image_id: imageId,
                    sub_id: subID
                })
            });

            if (!response.ok) {
                throw new Error('Failed to add to favorites');
            }

            return await response.json();
        } catch (error) {
            console.error('Error adding to favorites:', error);
            alert('Failed to add to favorites. Please try again.');
        }
    }

    // Event Listeners
    loveBtn.addEventListener('click', async () => {
        if (images.length > 0) {
            const currentImage = images[currentIndex];
            loveBtn.disabled = true;
            await addToFavorites(currentImage.id);
            loveBtn.disabled = false;
            await fetchImages(); // Get new images after saving
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

    // Disable buttons while loading
    function setButtonsState(disabled) {
        loveBtn.disabled = disabled;
        likeBtn.disabled = disabled;
        dislikeBtn.disabled = disabled;
    }

    // Preload next image
    function preloadNextImage() {
        if (images.length > currentIndex + 1) {
            const img = new Image();
            img.src = images[currentIndex + 1].url;
        }
    }

    // Initialize
    fetchImages();
});