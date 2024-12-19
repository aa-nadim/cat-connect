// Initialize function to be called with initial data
let breedsData = {};

function initializeBreeds(initialBreeds) {
    if (initialBreeds) {
        initialBreeds.forEach((breed) => {
            breedsData[breed.id] = breed;
        });
        populateBreedSelect();
    }
}

function refreshBreeds() {
    if (Object.keys(breedsData).length === 0) {
        loadBreeds();
    }
}

document.addEventListener('DOMContentLoaded', function() {
    const breedSelect = document.getElementById('breedSelect');
    const breedShowcase = document.getElementById('breedShowcase');
    const carouselInner = document.querySelector('.carousel-inner');
    const indicators = document.querySelector('.carousel-indicators');

    // Static placeholder image
    const placeholderImage = 'static/images/placeholder.jpg';

    async function loadBreeds() {
        try {
            const response = await fetch('/api/breeds');
            const breeds = await response.json();
            breeds.forEach((breed) => {
                breedsData[breed.id] = breed;
            });
            populateBreedSelect();
        } catch (error) {
            console.error('Error loading breeds:', error);
        }
    }

    function populateBreedSelect() {
        breedSelect.innerHTML = '<option value="">Select a breed</option>';
        Object.values(breedsData).forEach((breed) => {
            const option = document.createElement('option');
            option.value = breed.id;
            option.textContent = breed.name;
            breedSelect.appendChild(option);
        });

        if (breedSelect.options.length > 1) {
            breedSelect.selectedIndex = 1;
            breedSelect.dispatchEvent(new Event('change'));
        }
    }

    // Handle breed selection
    breedSelect.addEventListener('change', async function() {
        const breedId = this.value;
        if (!breedId) {
            breedShowcase.style.display = 'none';
            return;
        }

        breedShowcase.style.display = 'block';
        carouselInner.innerHTML = '<div class="text-center p-5">Loading...</div>';

        try {
            const response = await fetch(`/api/breeds/${breedId}`);
            const images = await response.json();

            if (images.error || images.length === 0) {
                throw new Error('No images found.');
            }

            updateCarousel(images);
            updateBreedInfo(breedId);
        } catch (error) {
            console.error('Error loading breed images:', error);
            showPlaceholder();
        }
    });

    function updateCarousel(images) {
        carouselInner.innerHTML = '';
        indicators.innerHTML = '';

        images.forEach((image, index) => {
            const item = document.createElement('div');
            item.className = `carousel-item ${index === 0 ? 'active' : ''}`;
            item.innerHTML = `
                <img src="${image.url || placeholderImage}" class="d-block w-100" alt="Cat" 
                     style="height: 400px; object-fit: cover;">
            `;
            carouselInner.appendChild(item);

            const indicator = document.createElement('button');
            indicator.setAttribute('type', 'button');
            indicator.setAttribute('data-bs-target', '#breedCarousel');
            indicator.setAttribute('data-bs-slide-to', index.toString());
            if (index === 0) indicator.classList.add('active');
            indicator.setAttribute('aria-label', `Slide ${index + 1}`);
            indicators.appendChild(indicator);
        });

        new bootstrap.Carousel(document.getElementById('breedCarousel'), {
            interval: 5000,
            wrap: true
        });
    }

    function updateBreedInfo(breedId) {
        const breed = breedsData[breedId];
        document.getElementById('breedName').textContent = breed.name;
        document.getElementById('breedOrigin').textContent = `Origin: ${breed.origin}`;
        document.getElementById('breedDescription').textContent = breed.description;
        
        const wikiLink = document.getElementById('wikiLink');
        if (breed.wikipedia_url) {
            wikiLink.href = breed.wikipedia_url;
            wikiLink.style.display = 'inline-block';
        } else {
            wikiLink.style.display = 'none';
        }
    }

    function showPlaceholder() {
        carouselInner.innerHTML = `
            <div class="carousel-item active">
                <img src="${placeholderImage}" class="d-block w-100" alt="Placeholder" 
                     style="height: 400px; object-fit: cover;">
            </div>
        `;
        indicators.innerHTML = '';
    }

    // Initialize if not already loaded
    if (Object.keys(breedsData).length === 0) {
        loadBreeds();
    }
});

// Export functions for tab switching
window.refreshBreeds = refreshBreeds;
window.initializeBreeds = initializeBreeds;