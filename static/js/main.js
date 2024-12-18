// static/js/main.js
document.addEventListener('DOMContentLoaded', function() {
    const breedSelect = document.getElementById('breedSelect');
    const carousel = document.getElementById('catCarousel');
    const breedName = document.getElementById('breedName');
    const breedDescription = document.getElementById('breedDescription');
    const breedOrigin = document.getElementById('breedOrigin');
    let carouselInstance = null;

    // Initialize the carousel
    function initCarousel() {
        carouselInstance = new bootstrap.Carousel(carousel, {
            interval: 3000,
            wrap: true
        });
    }

    // Load breeds
    fetch('/api/breeds')
        .then(response => response.json())
        .then(breeds => {
            if (!breeds || breeds.length === 0) {
                console.error('No breeds received');
                return;
            }
            breeds.forEach(breed => {
                const option = document.createElement('option');
                option.value = breed.id;
                option.textContent = breed.name;
                breedSelect.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Error loading breeds:', error);
        });
    
    // Handle breed selection
    breedSelect.addEventListener('change', function() {
        if (this.value) {
            fetch(`/api/breed-images?breed_id=${this.value}`)
                .then(response => response.json())
                .then(images => {
                    if (!images || images.length === 0) {
                        console.error('No images received');
                        return;
                    }
                    updateCarousel(images);
                    if (images[0].breeds && images[0].breeds[0]) {
                        const breed = images[0].breeds[0];
                        breedName.textContent = breed.name || '';
                        breedDescription.textContent = breed.description || '';
                        breedOrigin.textContent = breed.origin ? `Origin: ${breed.origin}` : '';
                    }
                })
                .catch(error => {
                    console.error('Error loading images:', error);
                });
        } else {
            // Reset display when no breed is selected
            breedName.textContent = '';
            breedDescription.textContent = '';
            breedOrigin.textContent = '';
            updateCarousel([]);
        }
    });
    
    function updateCarousel(images) {
        const carouselInner = document.querySelector('.carousel-inner');
        carouselInner.innerHTML = '';
        
        if (images.length === 0) {
            // Show placeholder if no images
            const div = document.createElement('div');
            div.className = 'carousel-item active';
            const img = document.createElement('img');
            img.src = '/static/img/placeholder.jpg';
            img.className = 'd-block w-100';
            img.alt = 'Select a breed';
            div.appendChild(img);
            carouselInner.appendChild(div);
        } else {
            images.forEach((image, index) => {
                const div = document.createElement('div');
                div.className = `carousel-item ${index === 0 ? 'active' : ''}`;
                
                const img = document.createElement('img');
                img.src = image.url;
                img.className = 'd-block w-100';
                img.alt = 'Cat Image';
                
                div.appendChild(img);
                carouselInner.appendChild(div);
            });
        }

        // Reinitialize carousel
        if (carouselInstance) {
            carouselInstance.dispose();
        }
        initCarousel();
    }
    
    // Handle manual navigation
    document.getElementById('prevBtn').addEventListener('click', () => {
        if (carouselInstance) {
            carouselInstance.prev();
        }
    });
    
    document.getElementById('nextBtn').addEventListener('click', () => {
        if (carouselInstance) {
            carouselInstance.next();
        }
    });

    // Initialize carousel on page load
    initCarousel();
});