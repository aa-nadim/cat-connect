document.addEventListener('DOMContentLoaded', function() {
    const breedSelect = document.getElementById('breedSelect');
    const breedShowcase = document.getElementById('breedShowcase');
    const carouselInner = document.querySelector('.carousel-inner');
    const indicators = document.querySelector('.carousel-indicators');

    // Static placeholder image
    const placeholderImage = 'static/images/placeholder.jpg';

    // Store breed data globally
    let breedsData = {};

    // Load breeds
    fetch('/api/breeds')
        .then(response => response.json())
        .then(breeds => {
            breeds.forEach((breed, index) => {
                breedsData[breed.id] = breed;
                const option = document.createElement('option');
                option.value = breed.id;
                option.textContent = breed.name;
                breedSelect.appendChild(option);
            });

            // Automatically select the second option
            if (breedSelect.options.length > 1) {
                breedSelect.selectedIndex = 1; // Select second option
                breedSelect.dispatchEvent(new Event('change')); // Trigger change event
            }
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
        fetch(`/api/breeds/${breedId}`)
            .then(response => response.json())
            .then(images => {
                if (images.error || images.length === 0) {
                    throw new Error('No images found.');
                }

                // Clear existing carousel content
                carouselInner.innerHTML = '';
                indicators.innerHTML = '';

                // Add images to carousel
                images.forEach((image, index) => {
                    const item = document.createElement('div');
                    item.className = `carousel-item ${index === 0 ? 'active' : ''}`;
                    const imageUrl = image.url || placeholderImage; // Use placeholder if no URL
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
                    <div class="carousel-item active">
                        <img src="${placeholderImage}" class="d-block w-100" alt="Placeholder" 
                             style="height: 400px; object-fit: cover;">
                    </div>
                `;
                indicators.innerHTML = ''; // No indicators for placeholder
            });
    });
});
