document.addEventListener('DOMContentLoaded', function() {
    let currentImages = [];
    let currentIndex = 0;

    const votingImage = document.getElementById('votingImage');
    const prevBtn = document.getElementById('prevBtn');
    const nextBtn = document.getElementById('nextBtn');
    const loveBtn = document.getElementById('loveBtn');

    function loadImages() {
        fetch('/api/voting')
            .then(response => response.json())
            .then(images => {
                currentImages = images;
                showImage(0);
            });
    }

    function showImage(index) {
        if (index >= 0 && index < currentImages.length) {
            currentIndex = index;
            votingImage.src = currentImages[index].url;
        }
    }

    function saveFavorite(image) {
        const favorites = JSON.parse(localStorage.getItem('catFavorites') || '[]');
        favorites.push(image);
        localStorage.setItem('catFavorites', JSON.stringify(favorites));
    }

    prevBtn.addEventListener('click', () => {
        if (currentIndex > 0) {
            showImage(currentIndex - 1);
        }
    });

    nextBtn.addEventListener('click', () => {
        if (currentIndex < currentImages.length - 1) {
            showImage(currentIndex + 1);
        } else {
            loadImages(); // Load new images when we reach the end
        }
    });

    loveBtn.addEventListener('click', () => {
        const currentImage = currentImages[currentIndex];
        saveFavorite(currentImage);
        if (currentIndex < currentImages.length - 1) {
            showImage(currentIndex + 1);
        } else {
            loadImages(); // Load new images when we reach the end
        }
    });

    // Initial load
    loadImages();
});