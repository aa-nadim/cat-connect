let favoritesData = [];

function initializeFavorites(initialFavorites) {
    if (initialFavorites) {
        favoritesData = initialFavorites;
        displayFavorites('grid');
    }
}

function refreshFavorites() {
    loadFavorites();
}

document.addEventListener('DOMContentLoaded', () => {
    const favsGallery = document.getElementById('favsGallery');
    const gridViewBtn = document.getElementById('gridView');
    const listViewBtn = document.getElementById('listView');
    const subID = 'user-123'; // Should match the subID used in voting.js

    // Set initial active state
    gridViewBtn.classList.add('active');

    async function loadFavorites() {
        try {
            const response = await fetch(`/api/favorites?sub_id=${subID}`);
            if (!response.ok) {
                throw new Error('Failed to load favorites');
            }
            favoritesData = await response.json();
            displayFavorites(favsGallery.classList.contains('grid-view') ? 'grid' : 'list');
        } catch (error) {
            console.error('Error loading favorites:', error);
            favoritesData = [];
            displayFavorites('grid');
        }
    }

    async function deleteFavorite(favoriteId) {
        try {
            const response = await fetch(`/api/favorites/${favoriteId}`, {
                method: 'DELETE'
            });
            if (!response.ok) {
                throw new Error('Failed to delete favorite');
            }
            await loadFavorites();
            return true;
        } catch (error) {
            console.error('Error deleting favorite:', error);
            return false;
        }
    }

    function displayFavorites(viewType) {
        favsGallery.innerHTML = '';
        favsGallery.className = viewType === 'grid' ? 'row grid-view g-4' : 'row list-view g-4';

        if (favoritesData.length === 0) {
            favsGallery.innerHTML = `
                <div class="col-12 text-center">
                    <p class="text-muted">No favorites yet. Add some from the voting section!</p>
                </div>
            `;
            return;
        }

        favoritesData.forEach(fav => {
            const col = document.createElement('div');
            col.className = viewType === 'grid' ? 'col-md-4 col-sm-6' : 'col-12';
            
            col.innerHTML = `
                <div class="card h-100">
                    <img src="${fav.image.url}" class="card-img-top" alt="Favorite Cat" 
                         style="height: ${viewType === 'grid' ? '200px' : '300px'}; object-fit: cover;">
                    <div class="card-body">
                        <div class="d-flex justify-content-between align-items-center">
                            <span class="text-muted">Added to favorites</span>
                            <button class="btn btn-outline-danger btn-sm delete-favorite" data-id="${fav.id}">
                                <i class="fas fa-trash"></i>
                            </button>
                        </div>
                    </div>
                </div>
            `;
            
            favsGallery.appendChild(col);
        });

        // Add event listeners to delete buttons
        document.querySelectorAll('.delete-favorite').forEach(btn => {
            btn.addEventListener('click', async (e) => {
                const favoriteId = e.currentTarget.dataset.id;
                await deleteFavorite(favoriteId);
            });
        });
    }

    // Add click handlers for view toggles
    gridViewBtn.addEventListener('click', () => {
        gridViewBtn.classList.add('active');
        listViewBtn.classList.remove('active');
        displayFavorites('grid');
    });

    listViewBtn.addEventListener('click', () => {
        listViewBtn.classList.add('active');
        gridViewBtn.classList.remove('active');
        displayFavorites('list');
    });

    // Initialize if no data loaded yet
    if (favoritesData.length === 0) {
        loadFavorites();
    }
});

// Export functions for tab switching
window.refreshFavorites = refreshFavorites;
window.initializeFavorites = initializeFavorites;