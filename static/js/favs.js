document.addEventListener('DOMContentLoaded', function() {
    const favsGallery = document.getElementById('favsGallery');
    const gridViewBtn = document.getElementById('gridView');
    const listViewBtn = document.getElementById('listView');

    function loadFavorites() {
        const favorites = JSON.parse(localStorage.getItem('catFavorites') || '[]');
        return favorites;
    }

    function displayGridView(favorites) {
        favsGallery.innerHTML = '';
        favsGallery.className = 'row g-4';
        
        favorites.forEach(image => {
            const col = document.createElement('div');
            col.className = 'col-md-4';
            col.innerHTML = `
                <div class="card">
                    <img src="${image.url}" class="card-img-top" alt="Cat">
                    <div class="card-body">
                        <button class="btn btn-sm btn-danger remove-fav" data-id="${image.id}">
                            <i class="fas fa-trash"></i> Remove
                        </button>
                    </div>
                </div>
            `;
            favsGallery.appendChild(col);
        });
    }

    function displayListView(favorites) {
        favsGallery.innerHTML = '';
        favsGallery.className = 'list-group';
        
        favorites.forEach(image => {
            const item = document.createElement('div');
            item.className = 'list-group-item d-flex align-items-center';
            item.innerHTML = `
                <img src="${image.url}" alt="Cat" style="width: 100px; height: 100px; object-fit: cover; margin-right: 15px;">
                <button class="btn btn-sm btn-danger ms-auto remove-fav" data-id="${image.id}">
                    <i class="fas fa-trash"></i> Remove
                </button>
            `;
            favsGallery.appendChild(item);
        });
    }

    function removeFavorite(imageId) {
        const favorites = loadFavorites();
        const updatedFavorites = favorites.filter(img => img.id !== imageId);
        localStorage.setItem('catFavorites', JSON.stringify(updatedFavorites));
        return updatedFavorites;
    }

    // Event listeners for view toggles
    gridViewBtn.addEventListener('click', () => {
        listViewBtn.classList.remove('active');
        gridViewBtn.classList.add('active');
        displayGridView(loadFavorites());
    });

    listViewBtn.addEventListener('click', () => {
        gridViewBtn.classList.remove('active');
        listViewBtn.classList.add('active');
        displayListView(loadFavorites());
    });

    // Event delegation for remove buttons
    favsGallery.addEventListener('click', (e) => {
        if (e.target.closest('.remove-fav')) {
            const imageId = e.target.closest('.remove-fav').dataset.id;
            const updatedFavorites = removeFavorite(imageId);
            if (gridViewBtn.classList.contains('active')) {
                displayGridView(updatedFavorites);
            } else {
                displayListView(updatedFavorites);
            }
        }
    });

    // Initial display
    displayGridView(loadFavorites());
});

