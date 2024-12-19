document.addEventListener('DOMContentLoaded', function() {
    const breedSelect = document.getElementById('breedSelect');
    const breedShowcase = document.getElementById('breedShowcase');
    const carouselInner = document.querySelector('.carousel-inner');
    const indicators = document.querySelector('.carousel-indicators');
    
    // Store breed data globally
    let breedsData = {};

    // Load breeds
    fetch('/api/breeds')
        .then(response => response.json())
        .then(breeds => {
            breeds.forEach(breed => {
                breedsData[breed.id] = breed;
                const option = document.createElement('option');
                option.value = breed.id;
                option.textContent = breed.name;
                breedSelect.appendChild(option);
            });
        })
        .catch(error => console.error('Error loading breeds:', error));

    // Handle breed selection
    breedSelect.addEventListener('change', function() {
        const breedId = this.value;
        if (!breedId) {
            breedShowcase.style.display = 'none';
            return;
        }

        // Show loader or placeholder
        breedShowcase.style.display = 'block';
        carouselInner.innerHTML = '<div class="text-center p-5">Loading...</div>';

        // Fetch images for selected breed
        fetch(`/api/breeds/${breedId}`)  // Changed this line to match the router path
            .then(response => response.json())
            .then(images => {
                if (images.error) {
                    throw new Error(images.error);
                }

                // Clear existing carousel content
                carouselInner.innerHTML = '';
                indicators.innerHTML = '';

                // Add images to carousel
                images.forEach((image, index) => {
                    // Create carousel item
                    const item = document.createElement('div');
                    item.className = `carousel-item ${index === 0 ? 'active' : ''}`;
                    
                    // Make sure we're using the correct URL from the API response
                    const imageUrl = image.url;
                    if (!imageUrl) {
                        console.error('No URL found in image object:', image);
                        return;
                    }

                    item.innerHTML = `
                        <img src="${imageUrl}" class="d-block w-100" alt="Cat" 
                             style="height: 400px; object-fit: cover;">
                    `;
                    carouselInner.appendChild(item);

                    // Create indicator
                    const indicator = document.createElement('button');
                    indicator.setAttribute('type', 'button');
                    indicator.setAttribute('data-bs-target', '#breedCarousel');
                    indicator.setAttribute('data-bs-slide-to', index.toString());
                    if (index === 0) indicator.classList.add('active');
                    indicator.setAttribute('aria-label', `Slide ${index + 1}`);
                    indicators.appendChild(indicator);
                });

                // Update breed information
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

                // Initialize carousel
                new bootstrap.Carousel(document.getElementById('breedCarousel'), {
                    interval: 5000,
                    wrap: true
                });
            })
            .catch(error => {
                console.error('Error loading breed images:', error);
                carouselInner.innerHTML = `
                    <div class="alert alert-danger m-3">
                        Error loading images. Please try again later.
                    </div>
                `;
            });
    });
});